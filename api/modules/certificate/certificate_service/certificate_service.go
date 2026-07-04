package certificate_service

import (
	"database/sql"
	"errors"

	"portfolio-api/base/helpers/error_helper"
	"portfolio-api/modules/certificate/certificate_dto"
	"portfolio-api/modules/certificate/certificate_model"
	"portfolio-api/modules/certificate/certificate_repository"

	"gopkg.in/guregu/null.v4"
)

// CertificateService exposes certificate operations scoped to a developer.
type CertificateService interface {
	GetByUser(userID int) ([]certificate_model.Certificate, error)
	GetByID(certificateID int) (certificate_model.Certificate, error)
	Create(userID int, request certificate_dto.CertificateRequest) (certificate_model.Certificate, error)
	Update(certificateID int, request certificate_dto.CertificateRequest) (certificate_model.Certificate, error)
	Delete(certificateID int) error
}

type certificateServiceImpl struct {
	repository certificate_repository.CertificateRepository
}

// NewCertificateService builds a CertificateService.
func NewCertificateService(repository certificate_repository.CertificateRepository) CertificateService {
	return &certificateServiceImpl{repository: repository}
}

func (s *certificateServiceImpl) GetByUser(userID int) ([]certificate_model.Certificate, error) {
	certificates, err := s.repository.FindByUser(userID)
	if err != nil {
		return nil, error_helper.Internal(err)
	}
	return certificates, nil
}

func (s *certificateServiceImpl) GetByID(certificateID int) (certificate_model.Certificate, error) {
	certificate, err := s.repository.FindByID(certificateID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return certificate_model.Certificate{}, error_helper.NotFound("certificate not found")
		}
		return certificate_model.Certificate{}, error_helper.Internal(err)
	}
	return certificate, nil
}

func (s *certificateServiceImpl) Create(userID int, request certificate_dto.CertificateRequest) (certificate_model.Certificate, error) {
	id, err := s.repository.Create(certificate_model.Certificate{
		UserID:  userID,
		Name:    request.Name,
		Issuer:  request.Issuer,
		Period:  request.Period,
		Link:    null.NewString(request.Link, request.Link != ""),
		OrderNo: request.OrderNo,
	})
	if err != nil {
		return certificate_model.Certificate{}, error_helper.Internal(err)
	}
	return s.GetByID(id)
}

func (s *certificateServiceImpl) Update(certificateID int, request certificate_dto.CertificateRequest) (certificate_model.Certificate, error) {
	if _, err := s.GetByID(certificateID); err != nil {
		return certificate_model.Certificate{}, err
	}

	if err := s.repository.Update(certificate_model.Certificate{
		CertificateID: certificateID,
		Name:          request.Name,
		Issuer:        request.Issuer,
		Period:        request.Period,
		Link:          null.NewString(request.Link, request.Link != ""),
		OrderNo:       request.OrderNo,
	}); err != nil {
		return certificate_model.Certificate{}, error_helper.Internal(err)
	}
	return s.GetByID(certificateID)
}

func (s *certificateServiceImpl) Delete(certificateID int) error {
	if _, err := s.GetByID(certificateID); err != nil {
		return err
	}
	if err := s.repository.Delete(certificateID); err != nil {
		return error_helper.Internal(err)
	}
	return nil
}
