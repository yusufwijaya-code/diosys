package skill_service

import (
	"database/sql"
	"errors"

	"portfolio-api/base/helpers/error_helper"
	"portfolio-api/modules/skill/skill_dto"
	"portfolio-api/modules/skill/skill_model"
	"portfolio-api/modules/skill/skill_repository"
)

// SkillService exposes skill operations scoped to a developer.
type SkillService interface {
	GetByUser(userID int) ([]skill_model.Skill, error)
	GetByID(skillID int) (skill_model.Skill, error)
	Create(userID int, request skill_dto.SkillRequest) (skill_model.Skill, error)
	Update(skillID int, request skill_dto.SkillRequest) (skill_model.Skill, error)
	Delete(skillID int) error
}

type skillServiceImpl struct {
	repository skill_repository.SkillRepository
}

// NewSkillService builds a SkillService.
func NewSkillService(repository skill_repository.SkillRepository) SkillService {
	return &skillServiceImpl{repository: repository}
}

func (s *skillServiceImpl) GetByUser(userID int) ([]skill_model.Skill, error) {
	skills, err := s.repository.FindByUser(userID)
	if err != nil {
		return nil, error_helper.Internal(err)
	}
	return skills, nil
}

func (s *skillServiceImpl) GetByID(skillID int) (skill_model.Skill, error) {
	skill, err := s.repository.FindByID(skillID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return skill_model.Skill{}, error_helper.NotFound("skill not found")
		}
		return skill_model.Skill{}, error_helper.Internal(err)
	}
	return skill, nil
}

func (s *skillServiceImpl) Create(userID int, request skill_dto.SkillRequest) (skill_model.Skill, error) {
	id, err := s.repository.Create(skill_model.Skill{
		UserID:   userID,
		Name:     request.Name,
		Level:    request.Level,
		Category: request.Category,
		OrderNo:  request.OrderNo,
	})
	if err != nil {
		return skill_model.Skill{}, error_helper.Internal(err)
	}
	return s.GetByID(id)
}

func (s *skillServiceImpl) Update(skillID int, request skill_dto.SkillRequest) (skill_model.Skill, error) {
	if _, err := s.GetByID(skillID); err != nil {
		return skill_model.Skill{}, err
	}

	if err := s.repository.Update(skill_model.Skill{
		SkillID:  skillID,
		Name:     request.Name,
		Level:    request.Level,
		Category: request.Category,
		OrderNo:  request.OrderNo,
	}); err != nil {
		return skill_model.Skill{}, error_helper.Internal(err)
	}
	return s.GetByID(skillID)
}

func (s *skillServiceImpl) Delete(skillID int) error {
	if _, err := s.GetByID(skillID); err != nil {
		return err
	}
	if err := s.repository.Delete(skillID); err != nil {
		return error_helper.Internal(err)
	}
	return nil
}
