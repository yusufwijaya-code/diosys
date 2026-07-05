package project_service

import (
	"database/sql"
	"errors"
	"mime/multipart"

	"portfolio-api/base/helpers/error_helper"
	"portfolio-api/base/helpers/gdrive_helper"
	"portfolio-api/modules/project/project_dto"
	"portfolio-api/modules/project/project_model"
	"portfolio-api/modules/project/project_repository"

	"gopkg.in/guregu/null.v4"
)

// ProjectService exposes project business operations.
type ProjectService interface {
	GetAllPublic() ([]project_dto.ProjectResponse, error)
	GetByUser(userID int) ([]project_dto.ProjectResponse, error)
	GetByID(projectID int) (project_dto.ProjectResponse, error)
	Create(userID int, request project_dto.ProjectRequest) (project_dto.ProjectResponse, error)
	Update(projectID int, request project_dto.ProjectRequest) (project_dto.ProjectResponse, error)
	Delete(projectID int) error
	UploadThumbnail(projectID int, fileHeader *multipart.FileHeader) (project_dto.ProjectResponse, error)
	AddImage(projectID int, fileHeader *multipart.FileHeader, caption string) (project_dto.ProjectResponse, error)
	DeleteImage(projectImageID int) error
	AddFeatureImage(projectFeatureID int, fileHeader *multipart.FileHeader, caption string) (project_dto.ProjectResponse, error)
	DeleteFeatureImage(projectFeatureImageID int) error
}

type projectServiceImpl struct {
	repository project_repository.ProjectRepository
	gdrive     *gdrive_helper.Client
}

// NewProjectService builds a ProjectService.
func NewProjectService(repository project_repository.ProjectRepository, gdrive *gdrive_helper.Client) ProjectService {
	return &projectServiceImpl{repository: repository, gdrive: gdrive}
}

func (s *projectServiceImpl) buildResponse(project project_model.Project, ownerUsername, ownerFullName string) (project_dto.ProjectResponse, error) {
	features, err := s.repository.GetFeatures(project.ProjectID)
	if err != nil {
		return project_dto.ProjectResponse{}, error_helper.Internal(err)
	}
	technologies, err := s.repository.GetTechnologies(project.ProjectID)
	if err != nil {
		return project_dto.ProjectResponse{}, error_helper.Internal(err)
	}
	images, err := s.repository.GetImages(project.ProjectID)
	if err != nil {
		return project_dto.ProjectResponse{}, error_helper.Internal(err)
	}

	featureResponses := make([]project_dto.ProjectFeatureResponse, 0, len(features))
	for _, feature := range features {
		featureImages, err := s.repository.GetFeatureImages(feature.ProjectFeatureID)
		if err != nil {
			return project_dto.ProjectResponse{}, error_helper.Internal(err)
		}
		imgResponses := make([]project_dto.ProjectFeatureImageResponse, 0, len(featureImages))
		for _, img := range featureImages {
			imgResponses = append(imgResponses, project_dto.ProjectFeatureImageResponse{
				ProjectFeatureImageID: img.ProjectFeatureImageID,
				Url:                   gdrive_helper.PublicURL(img.GdriveID),
				Caption:               img.Caption.String,
				OrderNo:               img.OrderNo,
			})
		}
		featureResponses = append(featureResponses, project_dto.ProjectFeatureResponse{
			ProjectFeatureID: feature.ProjectFeatureID,
			Text:             feature.Text,
			Description:      feature.Description.String,
			Images:           imgResponses,
			OrderNo:          feature.OrderNo,
		})
	}
	technologyNames := make([]string, 0, len(technologies))
	for _, technology := range technologies {
		technologyNames = append(technologyNames, technology.Name)
	}
	imageResponses := make([]project_dto.ProjectImageResponse, 0, len(images))
	for _, image := range images {
		imageResponses = append(imageResponses, project_dto.ProjectImageResponse{
			ProjectImageID: image.ProjectImageID,
			FileName:       image.FileName.String,
			GdriveID:       image.GdriveID.String,
			Url:            gdrive_helper.PublicURL(image.GdriveID.String),
			Caption:        image.Caption.String,
			DisplayOrder:   image.DisplayOrder,
		})
	}

	var statusID *int
	if project.ProjectStatusID.Valid {
		value := int(project.ProjectStatusID.Int64)
		statusID = &value
	}

	return project_dto.ProjectResponse{
		ProjectID:         project.ProjectID,
		UserID:            project.UserID,
		OwnerUsername:     ownerUsername,
		OwnerFullName:     ownerFullName,
		Title:             project.Title,
		Summary:           project.Summary.String,
		Body:              project.Body.String,
		Client:            project.Client.String,
		ProjectLink:       project.ProjectLink.String,
		RepoLink:          project.RepoLink.String,
		ProjectStatusID:   statusID,
		IsFeatured:        project.IsFeatured == 1,
		ThumbnailFileName: project.ThumbnailFileName.String,
		ThumbnailGdriveID: project.ThumbnailGdriveID.String,
		ThumbnailUrl:      gdrive_helper.PublicURL(project.ThumbnailGdriveID.String),
		OrderNo:           project.OrderNo,
		Features:          featureResponses,
		Technologies:      technologyNames,
		Images:            imageResponses,
	}, nil
}

func (s *projectServiceImpl) GetAllPublic() ([]project_dto.ProjectResponse, error) {
	projects, err := s.repository.FindAllPublic()
	if err != nil {
		return nil, error_helper.Internal(err)
	}

	responses := make([]project_dto.ProjectResponse, 0, len(projects))
	for _, project := range projects {
		response, err := s.buildResponse(project.Project, project.OwnerUsername, project.OwnerFullName)
		if err != nil {
			return nil, err
		}
		responses = append(responses, response)
	}
	return responses, nil
}

func (s *projectServiceImpl) GetByUser(userID int) ([]project_dto.ProjectResponse, error) {
	projects, err := s.repository.FindByUser(userID)
	if err != nil {
		return nil, error_helper.Internal(err)
	}

	responses := make([]project_dto.ProjectResponse, 0, len(projects))
	for _, project := range projects {
		response, err := s.buildResponse(project, "", "")
		if err != nil {
			return nil, err
		}
		responses = append(responses, response)
	}
	return responses, nil
}

func (s *projectServiceImpl) GetByID(projectID int) (project_dto.ProjectResponse, error) {
	project, err := s.repository.FindByIDWithOwner(projectID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return project_dto.ProjectResponse{}, error_helper.NotFound("project not found")
		}
		return project_dto.ProjectResponse{}, error_helper.Internal(err)
	}
	return s.buildResponse(project.Project, project.OwnerUsername, project.OwnerFullName)
}

func (s *projectServiceImpl) Create(userID int, request project_dto.ProjectRequest) (project_dto.ProjectResponse, error) {
	id, err := s.repository.Create(toModel(0, userID, request), request.Features, request.Technologies)
	if err != nil {
		return project_dto.ProjectResponse{}, error_helper.Internal(err)
	}
	return s.GetByID(id)
}

func (s *projectServiceImpl) Update(projectID int, request project_dto.ProjectRequest) (project_dto.ProjectResponse, error) {
	existing, err := s.repository.FindByID(projectID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return project_dto.ProjectResponse{}, error_helper.NotFound("project not found")
		}
		return project_dto.ProjectResponse{}, error_helper.Internal(err)
	}

	if err := s.repository.Update(toModel(projectID, existing.UserID, request), request.Features, request.Technologies); err != nil {
		return project_dto.ProjectResponse{}, error_helper.Internal(err)
	}
	return s.GetByID(projectID)
}

func (s *projectServiceImpl) Delete(projectID int) error {
	project, err := s.repository.FindByID(projectID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return error_helper.NotFound("project not found")
		}
		return error_helper.Internal(err)
	}

	images, _ := s.repository.GetImages(projectID)

	if err := s.repository.Delete(projectID); err != nil {
		return error_helper.Internal(err)
	}

	if s.gdrive != nil {
		if project.ThumbnailGdriveID.Valid && project.ThumbnailGdriveID.String != "" {
			_ = s.gdrive.DeleteFile(project.ThumbnailGdriveID.String)
		}
		for _, image := range images {
			if image.GdriveID.Valid && image.GdriveID.String != "" {
				_ = s.gdrive.DeleteFile(image.GdriveID.String)
			}
		}
	}
	return nil
}

func (s *projectServiceImpl) UploadThumbnail(projectID int, fileHeader *multipart.FileHeader) (project_dto.ProjectResponse, error) {
	if s.gdrive == nil {
		return project_dto.ProjectResponse{}, error_helper.Internal(errors.New("google drive is not configured"))
	}

	project, err := s.repository.FindByID(projectID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return project_dto.ProjectResponse{}, error_helper.NotFound("project not found")
		}
		return project_dto.ProjectResponse{}, error_helper.Internal(err)
	}

	uploaded, err := s.gdrive.UploadImage(fileHeader, "project")
	if err != nil {
		return project_dto.ProjectResponse{}, error_helper.Internal(err)
	}

	if project.ThumbnailGdriveID.Valid && project.ThumbnailGdriveID.String != "" {
		_ = s.gdrive.DeleteFile(project.ThumbnailGdriveID.String)
	}

	if err := s.repository.UpdateThumbnail(projectID, uploaded.FileName, uploaded.GdriveID); err != nil {
		return project_dto.ProjectResponse{}, error_helper.Internal(err)
	}
	return s.GetByID(projectID)
}

func (s *projectServiceImpl) AddImage(projectID int, fileHeader *multipart.FileHeader, caption string) (project_dto.ProjectResponse, error) {
	if s.gdrive == nil {
		return project_dto.ProjectResponse{}, error_helper.Internal(errors.New("google drive is not configured"))
	}

	if _, err := s.repository.FindByID(projectID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return project_dto.ProjectResponse{}, error_helper.NotFound("project not found")
		}
		return project_dto.ProjectResponse{}, error_helper.Internal(err)
	}

	existingImages, _ := s.repository.GetImages(projectID)

	uploaded, err := s.gdrive.UploadImage(fileHeader, "project-gallery")
	if err != nil {
		return project_dto.ProjectResponse{}, error_helper.Internal(err)
	}

	if _, err := s.repository.AddImage(project_model.ProjectImage{
		ProjectID:    projectID,
		FileName:     null.StringFrom(uploaded.FileName),
		GdriveID:     null.StringFrom(uploaded.GdriveID),
		Caption:      null.NewString(caption, caption != ""),
		DisplayOrder: len(existingImages),
	}); err != nil {
		return project_dto.ProjectResponse{}, error_helper.Internal(err)
	}
	return s.GetByID(projectID)
}

func (s *projectServiceImpl) DeleteImage(projectImageID int) error {
	image, err := s.repository.FindImageByID(projectImageID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return error_helper.NotFound("image not found")
		}
		return error_helper.Internal(err)
	}

	if err := s.repository.DeleteImage(projectImageID); err != nil {
		return error_helper.Internal(err)
	}

	if s.gdrive != nil && image.GdriveID.Valid && image.GdriveID.String != "" {
		_ = s.gdrive.DeleteFile(image.GdriveID.String)
	}
	return nil
}

func (s *projectServiceImpl) AddFeatureImage(projectFeatureID int, fileHeader *multipart.FileHeader, caption string) (project_dto.ProjectResponse, error) {
	if s.gdrive == nil {
		return project_dto.ProjectResponse{}, error_helper.Internal(errors.New("google drive is not configured"))
	}
	feature, err := s.repository.FindFeatureByID(projectFeatureID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return project_dto.ProjectResponse{}, error_helper.NotFound("feature not found")
		}
		return project_dto.ProjectResponse{}, error_helper.Internal(err)
	}
	uploaded, err := s.gdrive.UploadImage(fileHeader, "project-feature")
	if err != nil {
		return project_dto.ProjectResponse{}, error_helper.Internal(err)
	}
	existing, _ := s.repository.GetFeatureImages(projectFeatureID)
	if _, err := s.repository.AddFeatureImage(project_model.ProjectFeatureImage{
		ProjectFeatureID: projectFeatureID,
		GdriveID:         uploaded.GdriveID,
		FileName:         uploaded.FileName,
		Caption:          null.NewString(caption, caption != ""),
		OrderNo:          len(existing),
	}); err != nil {
		return project_dto.ProjectResponse{}, error_helper.Internal(err)
	}
	return s.GetByID(feature.ProjectID)
}

func (s *projectServiceImpl) DeleteFeatureImage(projectFeatureImageID int) error {
	img, err := s.repository.DeleteFeatureImage(projectFeatureImageID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return error_helper.NotFound("image not found")
		}
		return error_helper.Internal(err)
	}
	if s.gdrive != nil && img.GdriveID != "" {
		_ = s.gdrive.DeleteFile(img.GdriveID)
	}
	return nil
}

func toModel(projectID, userID int, request project_dto.ProjectRequest) project_model.Project {
	isFeatured := 0
	if request.IsFeatured {
		isFeatured = 1
	}
	var statusID null.Int
	if request.ProjectStatusID != nil {
		statusID = null.IntFrom(int64(*request.ProjectStatusID))
	}
	return project_model.Project{
		ProjectID:       projectID,
		UserID:          userID,
		Title:           request.Title,
		Summary:         null.NewString(request.Summary, request.Summary != ""),
		Body:            null.NewString(request.Body, request.Body != ""),
		Client:          null.NewString(request.Client, request.Client != ""),
		ProjectLink:     null.NewString(request.ProjectLink, request.ProjectLink != ""),
		RepoLink:        null.NewString(request.RepoLink, request.RepoLink != ""),
		ProjectStatusID: statusID,
		IsFeatured:      isFeatured,
		OrderNo:         request.OrderNo,
	}
}
