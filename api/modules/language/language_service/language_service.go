package language_service

import (
	"database/sql"
	"errors"

	"portfolio-api/base/helpers/error_helper"
	"portfolio-api/modules/language/language_dto"
	"portfolio-api/modules/language/language_model"
	"portfolio-api/modules/language/language_repository"
)

// LanguageService exposes language operations scoped to a developer.
type LanguageService interface {
	GetByUser(userID int) ([]language_model.Language, error)
	GetByID(languageID int) (language_model.Language, error)
	Create(userID int, request language_dto.LanguageRequest) (language_model.Language, error)
	Update(languageID int, request language_dto.LanguageRequest) (language_model.Language, error)
	Delete(languageID int) error
}

type languageServiceImpl struct {
	repository language_repository.LanguageRepository
}

// NewLanguageService builds a LanguageService.
func NewLanguageService(repository language_repository.LanguageRepository) LanguageService {
	return &languageServiceImpl{repository: repository}
}

func (s *languageServiceImpl) GetByUser(userID int) ([]language_model.Language, error) {
	languages, err := s.repository.FindByUser(userID)
	if err != nil {
		return nil, error_helper.Internal(err)
	}
	return languages, nil
}

func (s *languageServiceImpl) GetByID(languageID int) (language_model.Language, error) {
	language, err := s.repository.FindByID(languageID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return language_model.Language{}, error_helper.NotFound("language not found")
		}
		return language_model.Language{}, error_helper.Internal(err)
	}
	return language, nil
}

func (s *languageServiceImpl) Create(userID int, request language_dto.LanguageRequest) (language_model.Language, error) {
	id, err := s.repository.Create(language_model.Language{
		UserID:  userID,
		Name:    request.Name,
		Level:   request.Level,
		Icon:    request.Icon,
		OrderNo: request.OrderNo,
	})
	if err != nil {
		return language_model.Language{}, error_helper.Internal(err)
	}
	return s.GetByID(id)
}

func (s *languageServiceImpl) Update(languageID int, request language_dto.LanguageRequest) (language_model.Language, error) {
	if _, err := s.GetByID(languageID); err != nil {
		return language_model.Language{}, err
	}

	if err := s.repository.Update(language_model.Language{
		LanguageID: languageID,
		Name:       request.Name,
		Level:      request.Level,
		Icon:       request.Icon,
		OrderNo:    request.OrderNo,
	}); err != nil {
		return language_model.Language{}, error_helper.Internal(err)
	}
	return s.GetByID(languageID)
}

func (s *languageServiceImpl) Delete(languageID int) error {
	if _, err := s.GetByID(languageID); err != nil {
		return err
	}
	if err := s.repository.Delete(languageID); err != nil {
		return error_helper.Internal(err)
	}
	return nil
}
