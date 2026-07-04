package setting_service

import (
	"portfolio-api/base/helpers/error_helper"
	"portfolio-api/modules/setting/setting_dto"
	"portfolio-api/modules/setting/setting_model"
	"portfolio-api/modules/setting/setting_repository"
)

// SettingService exposes system setting operations.
type SettingService interface {
	GetAll() ([]setting_model.Setting, error)
	GetPublicMap() (map[string]string, error)
	Update(request setting_dto.SettingUpdateRequest) (map[string]string, error)
}

type settingServiceImpl struct {
	repository setting_repository.SettingRepository
}

// NewSettingService builds a SettingService.
func NewSettingService(repository setting_repository.SettingRepository) SettingService {
	return &settingServiceImpl{repository: repository}
}

func (s *settingServiceImpl) GetAll() ([]setting_model.Setting, error) {
	settings, err := s.repository.FindAll()
	if err != nil {
		return nil, error_helper.Internal(err)
	}
	return settings, nil
}

func (s *settingServiceImpl) GetPublicMap() (map[string]string, error) {
	settings, err := s.repository.FindAll()
	if err != nil {
		return nil, error_helper.Internal(err)
	}
	result := make(map[string]string, len(settings))
	for _, setting := range settings {
		result[setting.SettingKey] = setting.SettingValue.String
	}
	return result, nil
}

func (s *settingServiceImpl) Update(request setting_dto.SettingUpdateRequest) (map[string]string, error) {
	for _, item := range request.Settings {
		if item.SettingKey == "" {
			continue
		}
		if err := s.repository.Upsert(item.SettingKey, item.SettingValue); err != nil {
			return nil, error_helper.Internal(err)
		}
	}
	return s.GetPublicMap()
}
