package service_service

import (
	"database/sql"
	"errors"

	"portfolio-api/base/helpers/error_helper"
	"portfolio-api/modules/service/service_dto"
	"portfolio-api/modules/service/service_model"
	"portfolio-api/modules/service/service_repository"

	"gopkg.in/guregu/null.v4"
)

// Service exposes agency service operations.
type Service interface {
	GetPublic() ([]service_model.Service, error)
	GetAll() ([]service_model.Service, error)
	GetByID(serviceID int) (service_model.Service, error)
	Create(request service_dto.ServiceRequest) (service_model.Service, error)
	Update(serviceID int, request service_dto.ServiceRequest) (service_model.Service, error)
	Delete(serviceID int) error
}

type serviceServiceImpl struct {
	repository service_repository.ServiceRepository
}

// NewService builds a Service.
func NewService(repository service_repository.ServiceRepository) Service {
	return &serviceServiceImpl{repository: repository}
}

func (s *serviceServiceImpl) GetPublic() ([]service_model.Service, error) {
	services, err := s.repository.FindAll(true)
	if err != nil {
		return nil, error_helper.Internal(err)
	}
	return services, nil
}

func (s *serviceServiceImpl) GetAll() ([]service_model.Service, error) {
	services, err := s.repository.FindAll(false)
	if err != nil {
		return nil, error_helper.Internal(err)
	}
	return services, nil
}

func (s *serviceServiceImpl) GetByID(serviceID int) (service_model.Service, error) {
	service, err := s.repository.FindByID(serviceID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return service_model.Service{}, error_helper.NotFound("service not found")
		}
		return service_model.Service{}, error_helper.Internal(err)
	}
	return service, nil
}

func (s *serviceServiceImpl) Create(request service_dto.ServiceRequest) (service_model.Service, error) {
	id, err := s.repository.Create(service_model.Service{
		Title:       request.Title,
		Description: null.NewString(request.Description, request.Description != ""),
		Icon:        null.NewString(request.Icon, request.Icon != ""),
		OrderNo:     request.OrderNo,
		FlagActive:  flagActiveValue(request.FlagActive, 1),
	})
	if err != nil {
		return service_model.Service{}, error_helper.Internal(err)
	}
	return s.GetByID(id)
}

func (s *serviceServiceImpl) Update(serviceID int, request service_dto.ServiceRequest) (service_model.Service, error) {
	existing, err := s.GetByID(serviceID)
	if err != nil {
		return service_model.Service{}, err
	}

	if err := s.repository.Update(service_model.Service{
		ServiceID:   serviceID,
		Title:       request.Title,
		Description: null.NewString(request.Description, request.Description != ""),
		Icon:        null.NewString(request.Icon, request.Icon != ""),
		OrderNo:     request.OrderNo,
		FlagActive:  flagActiveValue(request.FlagActive, existing.FlagActive),
	}); err != nil {
		return service_model.Service{}, error_helper.Internal(err)
	}
	return s.GetByID(serviceID)
}

func (s *serviceServiceImpl) Delete(serviceID int) error {
	if _, err := s.GetByID(serviceID); err != nil {
		return err
	}
	if err := s.repository.Delete(serviceID); err != nil {
		return error_helper.Internal(err)
	}
	return nil
}

func flagActiveValue(value *bool, fallback int) int {
	if value == nil {
		return fallback
	}
	if *value {
		return 1
	}
	return 0
}
