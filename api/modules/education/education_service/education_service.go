package education_service

import (
	"database/sql"
	"errors"

	"portfolio-api/base/helpers/error_helper"
	"portfolio-api/modules/education/education_dto"
	"portfolio-api/modules/education/education_model"
	"portfolio-api/modules/education/education_repository"
)

// EducationService exposes education business operations scoped to a developer.
type EducationService interface {
	GetByUser(userID int) ([]education_dto.EducationResponse, error)
	GetByID(educationID int) (education_dto.EducationResponse, error)
	Create(userID int, request education_dto.EducationRequest) (education_dto.EducationResponse, error)
	Update(educationID int, request education_dto.EducationRequest) (education_dto.EducationResponse, error)
	Delete(educationID int) error
}

type educationServiceImpl struct {
	repository education_repository.EducationRepository
}

// NewEducationService builds an EducationService.
func NewEducationService(repository education_repository.EducationRepository) EducationService {
	return &educationServiceImpl{repository: repository}
}

func (s *educationServiceImpl) buildResponse(education education_model.Education) (education_dto.EducationResponse, error) {
	achievements, err := s.repository.GetAchievements(education.EducationID)
	if err != nil {
		return education_dto.EducationResponse{}, error_helper.Internal(err)
	}

	achievementTexts := make([]string, 0, len(achievements))
	for _, achievement := range achievements {
		achievementTexts = append(achievementTexts, achievement.Description)
	}

	return education_dto.EducationResponse{
		EducationID:  education.EducationID,
		Degree:       education.Degree,
		Institution:  education.Institution,
		Year:         education.Year,
		Type:         education.Type,
		OrderNo:      education.OrderNo,
		Achievements: achievementTexts,
	}, nil
}

func (s *educationServiceImpl) GetByUser(userID int) ([]education_dto.EducationResponse, error) {
	educations, err := s.repository.FindByUser(userID)
	if err != nil {
		return nil, error_helper.Internal(err)
	}

	responses := make([]education_dto.EducationResponse, 0, len(educations))
	for _, education := range educations {
		response, err := s.buildResponse(education)
		if err != nil {
			return nil, err
		}
		responses = append(responses, response)
	}
	return responses, nil
}

func (s *educationServiceImpl) GetByID(educationID int) (education_dto.EducationResponse, error) {
	education, err := s.repository.FindByID(educationID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return education_dto.EducationResponse{}, error_helper.NotFound("education not found")
		}
		return education_dto.EducationResponse{}, error_helper.Internal(err)
	}
	return s.buildResponse(education)
}

func (s *educationServiceImpl) Create(userID int, request education_dto.EducationRequest) (education_dto.EducationResponse, error) {
	id, err := s.repository.Create(education_model.Education{
		UserID:      userID,
		Degree:      request.Degree,
		Institution: request.Institution,
		Year:        request.Year,
		Type:        request.Type,
		OrderNo:     request.OrderNo,
	}, request.Achievements)
	if err != nil {
		return education_dto.EducationResponse{}, error_helper.Internal(err)
	}
	return s.GetByID(id)
}

func (s *educationServiceImpl) Update(educationID int, request education_dto.EducationRequest) (education_dto.EducationResponse, error) {
	if _, err := s.GetByID(educationID); err != nil {
		return education_dto.EducationResponse{}, err
	}

	if err := s.repository.Update(education_model.Education{
		EducationID: educationID,
		Degree:      request.Degree,
		Institution: request.Institution,
		Year:        request.Year,
		Type:        request.Type,
		OrderNo:     request.OrderNo,
	}, request.Achievements); err != nil {
		return education_dto.EducationResponse{}, error_helper.Internal(err)
	}
	return s.GetByID(educationID)
}

func (s *educationServiceImpl) Delete(educationID int) error {
	if _, err := s.GetByID(educationID); err != nil {
		return err
	}
	if err := s.repository.Delete(educationID); err != nil {
		return error_helper.Internal(err)
	}
	return nil
}
