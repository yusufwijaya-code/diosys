package experience_service

import (
	"database/sql"
	"errors"

	"portfolio-api/base/helpers/error_helper"
	"portfolio-api/modules/experience/experience_dto"
	"portfolio-api/modules/experience/experience_model"
	"portfolio-api/modules/experience/experience_repository"
)

// ExperienceService exposes experience business operations scoped to a developer.
type ExperienceService interface {
	GetByUser(userID int) ([]experience_dto.ExperienceResponse, error)
	GetByID(experienceID int) (experience_dto.ExperienceResponse, error)
	Create(userID int, request experience_dto.ExperienceRequest) (experience_dto.ExperienceResponse, error)
	Update(experienceID int, request experience_dto.ExperienceRequest) (experience_dto.ExperienceResponse, error)
	Delete(experienceID int) error
}

type experienceServiceImpl struct {
	repository experience_repository.ExperienceRepository
}

// NewExperienceService builds an ExperienceService.
func NewExperienceService(repository experience_repository.ExperienceRepository) ExperienceService {
	return &experienceServiceImpl{repository: repository}
}

func (s *experienceServiceImpl) buildResponse(experience experience_model.Experience) (experience_dto.ExperienceResponse, error) {
	technologies, err := s.repository.GetTechnologies(experience.ExperienceID)
	if err != nil {
		return experience_dto.ExperienceResponse{}, error_helper.Internal(err)
	}
	responsibilities, err := s.repository.GetResponsibilities(experience.ExperienceID)
	if err != nil {
		return experience_dto.ExperienceResponse{}, error_helper.Internal(err)
	}

	technologyNames := make([]string, 0, len(technologies))
	for _, technology := range technologies {
		technologyNames = append(technologyNames, technology.Name)
	}
	responsibilityTexts := make([]string, 0, len(responsibilities))
	for _, responsibility := range responsibilities {
		responsibilityTexts = append(responsibilityTexts, responsibility.Description)
	}

	return experience_dto.ExperienceResponse{
		ExperienceID:     experience.ExperienceID,
		Position:         experience.Position,
		Company:          experience.Company,
		Period:           experience.Period,
		OrderNo:          experience.OrderNo,
		Technologies:     technologyNames,
		Responsibilities: responsibilityTexts,
	}, nil
}

func (s *experienceServiceImpl) GetByUser(userID int) ([]experience_dto.ExperienceResponse, error) {
	experiences, err := s.repository.FindByUser(userID)
	if err != nil {
		return nil, error_helper.Internal(err)
	}

	responses := make([]experience_dto.ExperienceResponse, 0, len(experiences))
	for _, experience := range experiences {
		response, err := s.buildResponse(experience)
		if err != nil {
			return nil, err
		}
		responses = append(responses, response)
	}
	return responses, nil
}

func (s *experienceServiceImpl) GetByID(experienceID int) (experience_dto.ExperienceResponse, error) {
	experience, err := s.repository.FindByID(experienceID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return experience_dto.ExperienceResponse{}, error_helper.NotFound("experience not found")
		}
		return experience_dto.ExperienceResponse{}, error_helper.Internal(err)
	}
	return s.buildResponse(experience)
}

func (s *experienceServiceImpl) Create(userID int, request experience_dto.ExperienceRequest) (experience_dto.ExperienceResponse, error) {
	id, err := s.repository.Create(experience_model.Experience{
		UserID:   userID,
		Position: request.Position,
		Company:  request.Company,
		Period:   request.Period,
		OrderNo:  request.OrderNo,
	}, request.Technologies, request.Responsibilities)
	if err != nil {
		return experience_dto.ExperienceResponse{}, error_helper.Internal(err)
	}
	return s.GetByID(id)
}

func (s *experienceServiceImpl) Update(experienceID int, request experience_dto.ExperienceRequest) (experience_dto.ExperienceResponse, error) {
	if _, err := s.GetByID(experienceID); err != nil {
		return experience_dto.ExperienceResponse{}, err
	}

	if err := s.repository.Update(experience_model.Experience{
		ExperienceID: experienceID,
		Position:     request.Position,
		Company:      request.Company,
		Period:       request.Period,
		OrderNo:      request.OrderNo,
	}, request.Technologies, request.Responsibilities); err != nil {
		return experience_dto.ExperienceResponse{}, error_helper.Internal(err)
	}
	return s.GetByID(experienceID)
}

func (s *experienceServiceImpl) Delete(experienceID int) error {
	if _, err := s.GetByID(experienceID); err != nil {
		return err
	}
	if err := s.repository.Delete(experienceID); err != nil {
		return error_helper.Internal(err)
	}
	return nil
}
