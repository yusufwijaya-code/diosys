package certificate_repository

import (
	"portfolio-api/modules/certificate/certificate_model"

	"github.com/jmoiron/sqlx"
)

// CertificateRepository handles persistence for certificates.
type CertificateRepository interface {
	FindByUser(userID int) ([]certificate_model.Certificate, error)
	FindByID(certificateID int) (certificate_model.Certificate, error)
	Create(certificate certificate_model.Certificate) (int, error)
	Update(certificate certificate_model.Certificate) error
	Delete(certificateID int) error
}

type certificateRepositoryImpl struct {
	db *sqlx.DB
}

// NewCertificateRepository builds a CertificateRepository.
func NewCertificateRepository(db *sqlx.DB) CertificateRepository {
	return &certificateRepositoryImpl{db: db}
}

func (r *certificateRepositoryImpl) FindByUser(userID int) ([]certificate_model.Certificate, error) {
	certificates := []certificate_model.Certificate{}
	query := `SELECT certificateID, userID, name, issuer, period, link, orderNo, createdDate, editedDate
		FROM ms_certificate WHERE userID = ? ORDER BY orderNo ASC, certificateID ASC`
	err := r.db.Select(&certificates, query, userID)
	return certificates, err
}

func (r *certificateRepositoryImpl) FindByID(certificateID int) (certificate_model.Certificate, error) {
	var certificate certificate_model.Certificate
	query := `SELECT certificateID, userID, name, issuer, period, link, orderNo, createdDate, editedDate
		FROM ms_certificate WHERE certificateID = ? LIMIT 1`
	err := r.db.Get(&certificate, query, certificateID)
	return certificate, err
}

func (r *certificateRepositoryImpl) Create(certificate certificate_model.Certificate) (int, error) {
	query := `INSERT INTO ms_certificate (userID, name, issuer, period, link, orderNo) VALUES (?, ?, ?, ?, ?, ?)`
	result, err := r.db.Exec(query, certificate.UserID, certificate.Name, certificate.Issuer, certificate.Period, certificate.Link, certificate.OrderNo)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	return int(id), err
}

func (r *certificateRepositoryImpl) Update(certificate certificate_model.Certificate) error {
	query := `UPDATE ms_certificate SET name = ?, issuer = ?, period = ?, link = ?, orderNo = ? WHERE certificateID = ?`
	_, err := r.db.Exec(query, certificate.Name, certificate.Issuer, certificate.Period, certificate.Link, certificate.OrderNo, certificate.CertificateID)
	return err
}

func (r *certificateRepositoryImpl) Delete(certificateID int) error {
	_, err := r.db.Exec(`DELETE FROM ms_certificate WHERE certificateID = ?`, certificateID)
	return err
}
