package developer_service

import (
	"database/sql"
	"errors"
	"mime/multipart"

	"portfolio-api/base/helpers/error_helper"
	"portfolio-api/base/helpers/gdrive_helper"
	"portfolio-api/modules/certificate/certificate_service"
	"portfolio-api/modules/developer/developer_dto"
	"portfolio-api/modules/education/education_service"
	"portfolio-api/modules/experience/experience_service"
	"portfolio-api/modules/project/project_service"
	"portfolio-api/modules/skill/skill_service"
	"portfolio-api/modules/summary/summary_service"
	"portfolio-api/modules/user/user_model"
	"portfolio-api/modules/user/user_repository"

	"gopkg.in/guregu/null.v4"
)

// DeveloperService aggregates a developer's profile and manages developer records.
type DeveloperService interface {
	GetDirectory() ([]developer_dto.DeveloperCardResponse, error)
	GetProfileByUsername(username string) (developer_dto.DeveloperProfileResponse, error)
	List() ([]developer_dto.DeveloperResponse, error)
	GetByID(userID int) (developer_dto.DeveloperResponse, error)
	Create(request developer_dto.DeveloperRequest) (developer_dto.DeveloperResponse, error)
	Update(userID int, request developer_dto.DeveloperRequest) (developer_dto.DeveloperResponse, error)
	Delete(userID int) error
	UploadPhoto(userID int, fileHeader *multipart.FileHeader) (developer_dto.DeveloperResponse, error)
	UploadCV(userID int, fileHeader *multipart.FileHeader) (developer_dto.DeveloperResponse, error)
}

type developerServiceImpl struct {
	userRepository     user_repository.UserRepository
	summaryService     summary_service.SummaryService
	experienceService  experience_service.ExperienceService
	educationService   education_service.EducationService
	certificateService certificate_service.CertificateService
	skillService       skill_service.SkillService
	projectService     project_service.ProjectService
	gdrive             *gdrive_helper.Client
}

// NewDeveloperService builds a DeveloperService.
func NewDeveloperService(
	userRepository user_repository.UserRepository,
	summaryService summary_service.SummaryService,
	experienceService experience_service.ExperienceService,
	educationService education_service.EducationService,
	certificateService certificate_service.CertificateService,
	skillService skill_service.SkillService,
	projectService project_service.ProjectService,
	gdrive *gdrive_helper.Client,
) DeveloperService {
	return &developerServiceImpl{
		userRepository:     userRepository,
		summaryService:     summaryService,
		experienceService:  experienceService,
		educationService:   educationService,
		certificateService: certificateService,
		skillService:       skillService,
		projectService:     projectService,
		gdrive:             gdrive,
	}
}

func (s *developerServiceImpl) mapToResponse(user user_model.User) developer_dto.DeveloperResponse {
	return developer_dto.DeveloperResponse{
		UserID:         user.UserID,
		Username:       user.Username,
		Email:          user.Email,
		FullName:       user.FullName,
		JobTitle:       user.JobTitle.String,
		Intro:          user.Intro.String,
		Bio:            user.Bio.String,
		Specialization: user.Specialization.String,
		Phone:          user.Phone.String,
		Website:        user.Website.String,
		GithubUrl:      user.GithubUrl.String,
		LinkedinUrl:    user.LinkedinUrl.String,
		InstagramUrl:   user.InstagramUrl.String,
		CvUrl:          gdrive_helper.DownloadURL(user.CvGdriveID.String),
		Location:       user.Location.String,
		PhotoUrl:       gdrive_helper.PublicURL(user.PhotoGdriveID.String),
		IsAdmin:        user.IsAdmin,
		FlagActive:     user.FlagActive,
		OrderNo:        user.OrderNo,
	}
}

func (s *developerServiceImpl) GetDirectory() ([]developer_dto.DeveloperCardResponse, error) {
	users, err := s.userRepository.FindAllDevelopers()
	if err != nil {
		return nil, error_helper.Internal(err)
	}

	cards := make([]developer_dto.DeveloperCardResponse, 0, len(users))
	for _, user := range users {
		cards = append(cards, developer_dto.DeveloperCardResponse{
			UserID:   user.UserID,
			Username: user.Username,
			FullName: user.FullName,
			JobTitle: user.JobTitle.String,
			Intro:    user.Intro.String,
			PhotoUrl: gdrive_helper.PublicURL(user.PhotoGdriveID.String),
		})
	}
	return cards, nil
}

func (s *developerServiceImpl) GetProfileByUsername(username string) (developer_dto.DeveloperProfileResponse, error) {
	user, err := s.userRepository.FindByUsername(username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return developer_dto.DeveloperProfileResponse{}, error_helper.NotFound("developer not found")
		}
		return developer_dto.DeveloperProfileResponse{}, error_helper.Internal(err)
	}

	summary, err := s.summaryService.GetByUser(user.UserID)
	if err != nil {
		return developer_dto.DeveloperProfileResponse{}, err
	}
	experiences, err := s.experienceService.GetByUser(user.UserID)
	if err != nil {
		return developer_dto.DeveloperProfileResponse{}, err
	}
	educations, err := s.educationService.GetByUser(user.UserID)
	if err != nil {
		return developer_dto.DeveloperProfileResponse{}, err
	}
	certificates, err := s.certificateService.GetByUser(user.UserID)
	if err != nil {
		return developer_dto.DeveloperProfileResponse{}, err
	}
	skills, err := s.skillService.GetByUser(user.UserID)
	if err != nil {
		return developer_dto.DeveloperProfileResponse{}, err
	}
	projects, err := s.projectService.GetByUser(user.UserID)
	if err != nil {
		return developer_dto.DeveloperProfileResponse{}, err
	}

	return developer_dto.DeveloperProfileResponse{
		Developer:    s.mapToResponse(user),
		Summary:      summary,
		Experiences:  experiences,
		Educations:   educations,
		Certificates: certificates,
		Skills:       skills,
		Projects:     projects,
	}, nil
}

func (s *developerServiceImpl) List() ([]developer_dto.DeveloperResponse, error) {
	users, err := s.userRepository.FindAllDevelopers()
	if err != nil {
		return nil, error_helper.Internal(err)
	}
	responses := make([]developer_dto.DeveloperResponse, 0, len(users))
	for _, user := range users {
		responses = append(responses, s.mapToResponse(user))
	}
	return responses, nil
}

func (s *developerServiceImpl) GetByID(userID int) (developer_dto.DeveloperResponse, error) {
	user, err := s.userRepository.FindByID(userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return developer_dto.DeveloperResponse{}, error_helper.NotFound("developer not found")
		}
		return developer_dto.DeveloperResponse{}, error_helper.Internal(err)
	}
	return s.mapToResponse(user), nil
}

func (s *developerServiceImpl) Create(request developer_dto.DeveloperRequest) (developer_dto.DeveloperResponse, error) {
	flagActive := 1
	if request.FlagActive != nil && !*request.FlagActive {
		flagActive = 0
	}

	id, err := s.userRepository.Create(user_model.User{
		Username:       request.Username,
		Email:          request.Email,
		FullName:       request.FullName,
		JobTitle:       null.NewString(request.JobTitle, request.JobTitle != ""),
		Intro:          null.NewString(request.Intro, request.Intro != ""),
		Bio:            null.NewString(request.Bio, request.Bio != ""),
		Specialization: null.NewString(request.Specialization, request.Specialization != ""),
		Phone:          null.NewString(request.Phone, request.Phone != ""),
		Website:        null.NewString(request.Website, request.Website != ""),
		GithubUrl:      null.NewString(request.GithubUrl, request.GithubUrl != ""),
		LinkedinUrl:    null.NewString(request.LinkedinUrl, request.LinkedinUrl != ""),
		InstagramUrl:   null.NewString(request.InstagramUrl, request.InstagramUrl != ""),
		Location:       null.NewString(request.Location, request.Location != ""),
		UserRoleID:     null.IntFrom(2),
		IsAdmin:        0,
		FlagActive:     flagActive,
		OrderNo:        request.OrderNo,
	})
	if err != nil {
		return developer_dto.DeveloperResponse{}, error_helper.Internal(err)
	}
	return s.GetByID(id)
}

func (s *developerServiceImpl) Update(userID int, request developer_dto.DeveloperRequest) (developer_dto.DeveloperResponse, error) {
	existing, err := s.userRepository.FindByID(userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return developer_dto.DeveloperResponse{}, error_helper.NotFound("developer not found")
		}
		return developer_dto.DeveloperResponse{}, error_helper.Internal(err)
	}

	flagActive := existing.FlagActive
	if request.FlagActive != nil {
		if *request.FlagActive {
			flagActive = 1
		} else {
			flagActive = 0
		}
	}

	existing.Username = request.Username
	existing.Email = request.Email
	existing.FullName = request.FullName
	existing.JobTitle = null.NewString(request.JobTitle, request.JobTitle != "")
	existing.Intro = null.NewString(request.Intro, request.Intro != "")
	existing.Bio = null.NewString(request.Bio, request.Bio != "")
	existing.Specialization = null.NewString(request.Specialization, request.Specialization != "")
	existing.Phone = null.NewString(request.Phone, request.Phone != "")
	existing.Website = null.NewString(request.Website, request.Website != "")
	existing.GithubUrl = null.NewString(request.GithubUrl, request.GithubUrl != "")
	existing.LinkedinUrl = null.NewString(request.LinkedinUrl, request.LinkedinUrl != "")
	existing.InstagramUrl = null.NewString(request.InstagramUrl, request.InstagramUrl != "")
	existing.Location = null.NewString(request.Location, request.Location != "")
	existing.FlagActive = flagActive
	existing.OrderNo = request.OrderNo

	if err := s.userRepository.Update(existing); err != nil {
		return developer_dto.DeveloperResponse{}, error_helper.Internal(err)
	}
	return s.GetByID(userID)
}

func (s *developerServiceImpl) Delete(userID int) error {
	user, err := s.userRepository.FindByID(userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return error_helper.NotFound("developer not found")
		}
		return error_helper.Internal(err)
	}
	if user.IsAdmin == 1 {
		return error_helper.Forbidden("the administrator account cannot be deleted")
	}
	if err := s.userRepository.Delete(userID); err != nil {
		return error_helper.Internal(err)
	}
	if s.gdrive != nil && user.PhotoGdriveID.Valid && user.PhotoGdriveID.String != "" {
		_ = s.gdrive.DeleteFile(user.PhotoGdriveID.String)
	}
	return nil
}

func (s *developerServiceImpl) UploadPhoto(userID int, fileHeader *multipart.FileHeader) (developer_dto.DeveloperResponse, error) {
	if s.gdrive == nil {
		return developer_dto.DeveloperResponse{}, error_helper.Internal(errors.New("google drive is not configured"))
	}

	user, err := s.userRepository.FindByID(userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return developer_dto.DeveloperResponse{}, error_helper.NotFound("developer not found")
		}
		return developer_dto.DeveloperResponse{}, error_helper.Internal(err)
	}

	uploaded, err := s.gdrive.UploadImage(fileHeader, "developer")
	if err != nil {
		return developer_dto.DeveloperResponse{}, error_helper.Internal(err)
	}

	if user.PhotoGdriveID.Valid && user.PhotoGdriveID.String != "" {
		_ = s.gdrive.DeleteFile(user.PhotoGdriveID.String)
	}

	if err := s.userRepository.UpdatePhoto(userID, uploaded.FileName, uploaded.GdriveID); err != nil {
		return developer_dto.DeveloperResponse{}, error_helper.Internal(err)
	}
	return s.GetByID(userID)
}

func (s *developerServiceImpl) UploadCV(userID int, fileHeader *multipart.FileHeader) (developer_dto.DeveloperResponse, error) {
	if s.gdrive == nil {
		return developer_dto.DeveloperResponse{}, error_helper.Internal(errors.New("google drive is not configured"))
	}

	user, err := s.userRepository.FindByID(userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return developer_dto.DeveloperResponse{}, error_helper.NotFound("developer not found")
		}
		return developer_dto.DeveloperResponse{}, error_helper.Internal(err)
	}

	uploaded, err := s.gdrive.UploadImage(fileHeader, "cv")
	if err != nil {
		return developer_dto.DeveloperResponse{}, error_helper.Internal(err)
	}

	if user.CvGdriveID.Valid && user.CvGdriveID.String != "" {
		_ = s.gdrive.DeleteFile(user.CvGdriveID.String)
	}

	if err := s.userRepository.UpdateCV(userID, uploaded.FileName, uploaded.GdriveID); err != nil {
		return developer_dto.DeveloperResponse{}, error_helper.Internal(err)
	}
	return s.GetByID(userID)
}
