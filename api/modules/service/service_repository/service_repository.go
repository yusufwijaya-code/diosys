package service_repository

import (
	"portfolio-api/modules/service/service_model"

	"github.com/jmoiron/sqlx"
)

const serviceColumns = `serviceID, title, description, icon, orderNo, flagActive, createdDate, editedDate`

// ServiceRepository handles persistence for agency services.
type ServiceRepository interface {
	FindAll(activeOnly bool) ([]service_model.Service, error)
	FindByID(serviceID int) (service_model.Service, error)
	Create(service service_model.Service) (int, error)
	Update(service service_model.Service) error
	Delete(serviceID int) error
}

type serviceRepositoryImpl struct {
	db *sqlx.DB
}

// NewServiceRepository builds a ServiceRepository.
func NewServiceRepository(db *sqlx.DB) ServiceRepository {
	return &serviceRepositoryImpl{db: db}
}

func (r *serviceRepositoryImpl) FindAll(activeOnly bool) ([]service_model.Service, error) {
	services := []service_model.Service{}
	query := `SELECT ` + serviceColumns + ` FROM ms_service`
	if activeOnly {
		query += ` WHERE flagActive = 1`
	}
	query += ` ORDER BY orderNo ASC, serviceID ASC`
	err := r.db.Select(&services, query)
	return services, err
}

func (r *serviceRepositoryImpl) FindByID(serviceID int) (service_model.Service, error) {
	var service service_model.Service
	query := `SELECT ` + serviceColumns + ` FROM ms_service WHERE serviceID = ? LIMIT 1`
	err := r.db.Get(&service, query, serviceID)
	return service, err
}

func (r *serviceRepositoryImpl) Create(service service_model.Service) (int, error) {
	query := `INSERT INTO ms_service (title, description, icon, orderNo, flagActive) VALUES (?, ?, ?, ?, ?)`
	result, err := r.db.Exec(query, service.Title, service.Description, service.Icon, service.OrderNo, service.FlagActive)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	return int(id), err
}

func (r *serviceRepositoryImpl) Update(service service_model.Service) error {
	query := `UPDATE ms_service SET title = ?, description = ?, icon = ?, orderNo = ?, flagActive = ? WHERE serviceID = ?`
	_, err := r.db.Exec(query, service.Title, service.Description, service.Icon, service.OrderNo, service.FlagActive, service.ServiceID)
	return err
}

func (r *serviceRepositoryImpl) Delete(serviceID int) error {
	_, err := r.db.Exec(`DELETE FROM ms_service WHERE serviceID = ?`, serviceID)
	return err
}
