package professional_project_service

import (
	"database/sql"
	"errors"
	"mime/multipart"

	"portfolio-api/base/helpers/error_helper"
	"portfolio-api/base/helpers/gdrive_helper"
	"portfolio-api/modules/professional_project/professional_project_dto"
	"portfolio-api/modules/professional_project/professional_project_model"
	"portfolio-api/modules/professional_project/professional_project_repository"

	gopkg_null "gopkg.in/guregu/null.v4"
)

type ProfessionalProjectService interface {
	GetByUser(userID int) ([]professional_project_dto.ProfessionalProjectCardResponse, error)
	GetByID(projectID int) (professional_project_dto.ProfessionalProjectResponse, error)
	Create(userID int, req professional_project_dto.ProfessionalProjectRequest) (professional_project_dto.ProfessionalProjectResponse, error)
	Delete(projectID int) error
	UploadThumbnail(projectID int, fileHeader *multipart.FileHeader) (professional_project_dto.ProfessionalProjectResponse, error)
	AddFeature(projectID int, req professional_project_dto.ProjectFeatureRequest) (professional_project_dto.ProfessionalProjectResponse, error)
	DeleteFeature(featureID int) error
	AddFeatureImage(featureID int, fileHeader *multipart.FileHeader, caption string) (professional_project_dto.ProjectFeatureResponse, error)
	DeleteFeatureImage(featureImageID int) error
}

type profProjServiceImpl struct {
	repository professional_project_repository.ProfessionalProjectRepository
	gdrive     *gdrive_helper.Client
}

func NewProfessionalProjectService(
	repository professional_project_repository.ProfessionalProjectRepository,
	gdrive *gdrive_helper.Client,
) ProfessionalProjectService {
	return &profProjServiceImpl{repository: repository, gdrive: gdrive}
}

func (s *profProjServiceImpl) buildFeatureResponse(f professional_project_model.ProjectFeature) (professional_project_dto.ProjectFeatureResponse, error) {
	images, err := s.repository.GetFeatureImages(f.FeatureID)
	if err != nil {
		return professional_project_dto.ProjectFeatureResponse{}, error_helper.Internal(err)
	}
	imgResponses := make([]professional_project_dto.ProjectFeatureImageResponse, 0, len(images))
	for _, img := range images {
		imgResponses = append(imgResponses, professional_project_dto.ProjectFeatureImageResponse{
			FeatureImageID: img.FeatureImageID,
			Url:            gdrive_helper.PublicURL(img.GdriveID),
			Caption:        img.Caption.String,
			OrderNo:        img.OrderNo,
		})
	}
	return professional_project_dto.ProjectFeatureResponse{
		FeatureID:   f.FeatureID,
		Title:       f.Title,
		Description: f.Description.String,
		Images:      imgResponses,
		OrderNo:     f.OrderNo,
	}, nil
}

func (s *profProjServiceImpl) buildResponse(p professional_project_model.ProfessionalProject) (professional_project_dto.ProfessionalProjectResponse, error) {
	features, err := s.repository.GetFeatures(p.ProfessionalProjectID)
	if err != nil {
		return professional_project_dto.ProfessionalProjectResponse{}, error_helper.Internal(err)
	}
	featureResponses := make([]professional_project_dto.ProjectFeatureResponse, 0, len(features))
	for _, f := range features {
		fr, err := s.buildFeatureResponse(f)
		if err != nil {
			return professional_project_dto.ProfessionalProjectResponse{}, err
		}
		featureResponses = append(featureResponses, fr)
	}
	return professional_project_dto.ProfessionalProjectResponse{
		ProfessionalProjectID: p.ProfessionalProjectID,
		UserID:                p.UserID,
		Title:                 p.Title,
		Company:               p.Company,
		Summary:               p.Summary.String,
		ThumbnailUrl:          gdrive_helper.PublicURL(p.ThumbnailGdriveID.String),
		Features:              featureResponses,
		OrderNo:               p.OrderNo,
	}, nil
}

func (s *profProjServiceImpl) GetByUser(userID int) ([]professional_project_dto.ProfessionalProjectCardResponse, error) {
	projects, err := s.repository.FindByUser(userID)
	if err != nil {
		return nil, error_helper.Internal(err)
	}
	cards := make([]professional_project_dto.ProfessionalProjectCardResponse, 0, len(projects))
	for _, p := range projects {
		cards = append(cards, professional_project_dto.ProfessionalProjectCardResponse{
			ProfessionalProjectID: p.ProfessionalProjectID,
			Title:                 p.Title,
			Company:               p.Company,
			Summary:               p.Summary.String,
			ThumbnailUrl:          gdrive_helper.PublicURL(p.ThumbnailGdriveID.String),
			OrderNo:               p.OrderNo,
		})
	}
	return cards, nil
}

func (s *profProjServiceImpl) GetByID(projectID int) (professional_project_dto.ProfessionalProjectResponse, error) {
	p, err := s.repository.FindByIDWithOwner(projectID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return professional_project_dto.ProfessionalProjectResponse{}, error_helper.NotFound("professional project not found")
		}
		return professional_project_dto.ProfessionalProjectResponse{}, error_helper.Internal(err)
	}
	resp, err := s.buildResponse(p.ProfessionalProject)
	if err != nil {
		return resp, err
	}
	resp.OwnerPhone = p.OwnerPhone.String
	return resp, nil
}

func (s *profProjServiceImpl) Create(userID int, req professional_project_dto.ProfessionalProjectRequest) (professional_project_dto.ProfessionalProjectResponse, error) {
	id, err := s.repository.Create(professional_project_model.ProfessionalProject{
		UserID:  userID,
		Title:   req.Title,
		Company: req.Company,
		Summary: gopkg_null.NewString(req.Summary, req.Summary != ""),
		OrderNo: req.OrderNo,
	})
	if err != nil {
		return professional_project_dto.ProfessionalProjectResponse{}, error_helper.Internal(err)
	}
	return s.GetByID(id)
}

func (s *profProjServiceImpl) Delete(projectID int) error {
	p, err := s.repository.FindByID(projectID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return error_helper.NotFound("professional project not found")
		}
		return error_helper.Internal(err)
	}
	if err := s.repository.Delete(projectID); err != nil {
		return error_helper.Internal(err)
	}
	if s.gdrive != nil && p.ThumbnailGdriveID.Valid && p.ThumbnailGdriveID.String != "" {
		_ = s.gdrive.DeleteFile(p.ThumbnailGdriveID.String)
	}
	return nil
}

func (s *profProjServiceImpl) UploadThumbnail(projectID int, fileHeader *multipart.FileHeader) (professional_project_dto.ProfessionalProjectResponse, error) {
	if s.gdrive == nil {
		return professional_project_dto.ProfessionalProjectResponse{}, error_helper.Internal(errors.New("google drive is not configured"))
	}
	p, err := s.repository.FindByID(projectID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return professional_project_dto.ProfessionalProjectResponse{}, error_helper.NotFound("professional project not found")
		}
		return professional_project_dto.ProfessionalProjectResponse{}, error_helper.Internal(err)
	}
	uploaded, err := s.gdrive.UploadImage(fileHeader, "professional-project")
	if err != nil {
		return professional_project_dto.ProfessionalProjectResponse{}, error_helper.Internal(err)
	}
	if p.ThumbnailGdriveID.Valid && p.ThumbnailGdriveID.String != "" {
		_ = s.gdrive.DeleteFile(p.ThumbnailGdriveID.String)
	}
	if err := s.repository.UpdateThumbnail(projectID, uploaded.FileName, uploaded.GdriveID); err != nil {
		return professional_project_dto.ProfessionalProjectResponse{}, error_helper.Internal(err)
	}
	return s.GetByID(projectID)
}

func (s *profProjServiceImpl) AddFeature(projectID int, req professional_project_dto.ProjectFeatureRequest) (professional_project_dto.ProfessionalProjectResponse, error) {
	if _, err := s.GetByID(projectID); err != nil {
		return professional_project_dto.ProfessionalProjectResponse{}, err
	}
	if _, err := s.repository.AddFeature(professional_project_model.ProjectFeature{
		ProfessionalProjectID: projectID,
		Title:                 req.Title,
		Description:           gopkg_null.NewString(req.Description, req.Description != ""),
		OrderNo:               req.OrderNo,
	}); err != nil {
		return professional_project_dto.ProfessionalProjectResponse{}, error_helper.Internal(err)
	}
	return s.GetByID(projectID)
}

func (s *profProjServiceImpl) DeleteFeature(featureID int) error {
	if err := s.repository.DeleteFeature(featureID); err != nil {
		return error_helper.Internal(err)
	}
	return nil
}

func (s *profProjServiceImpl) AddFeatureImage(featureID int, fileHeader *multipart.FileHeader, caption string) (professional_project_dto.ProjectFeatureResponse, error) {
	if s.gdrive == nil {
		return professional_project_dto.ProjectFeatureResponse{}, error_helper.Internal(errors.New("google drive is not configured"))
	}
	f, err := s.repository.FindFeatureByID(featureID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return professional_project_dto.ProjectFeatureResponse{}, error_helper.NotFound("feature not found")
		}
		return professional_project_dto.ProjectFeatureResponse{}, error_helper.Internal(err)
	}
	uploaded, err := s.gdrive.UploadImage(fileHeader, "professional-project-feature")
	if err != nil {
		return professional_project_dto.ProjectFeatureResponse{}, error_helper.Internal(err)
	}
	if _, err := s.repository.AddFeatureImage(professional_project_model.ProjectFeatureImage{
		FeatureID: featureID,
		GdriveID:  uploaded.GdriveID,
		FileName:  uploaded.FileName,
		Caption:   gopkg_null.NewString(caption, caption != ""),
	}); err != nil {
		return professional_project_dto.ProjectFeatureResponse{}, error_helper.Internal(err)
	}
	return s.buildFeatureResponse(f)
}

func (s *profProjServiceImpl) DeleteFeatureImage(featureImageID int) error {
	img, err := s.repository.DeleteFeatureImage(featureImageID)
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
