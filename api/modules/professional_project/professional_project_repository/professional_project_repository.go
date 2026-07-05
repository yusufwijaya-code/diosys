package professional_project_repository

import (
	"portfolio-api/modules/professional_project/professional_project_model"

	"github.com/jmoiron/sqlx"
)

type ProfessionalProjectRepository interface {
	FindByUser(userID int) ([]professional_project_model.ProfessionalProject, error)
	FindByID(projectID int) (professional_project_model.ProfessionalProject, error)
	FindByIDWithOwner(projectID int) (professional_project_model.ProfessionalProjectWithOwner, error)
	Create(project professional_project_model.ProfessionalProject) (int, error)
	UpdateThumbnail(projectID int, fileName, gdriveID string) error
	Delete(projectID int) error
	GetFeatures(projectID int) ([]professional_project_model.ProjectFeature, error)
	FindFeatureByID(featureID int) (professional_project_model.ProjectFeature, error)
	AddFeature(feature professional_project_model.ProjectFeature) (int, error)
	DeleteFeature(featureID int) error
	GetFeatureImages(featureID int) ([]professional_project_model.ProjectFeatureImage, error)
	AddFeatureImage(image professional_project_model.ProjectFeatureImage) (int, error)
	DeleteFeatureImage(featureImageID int) (professional_project_model.ProjectFeatureImage, error)
}

type profProjRepositoryImpl struct {
	db *sqlx.DB
}

func NewProfessionalProjectRepository(db *sqlx.DB) ProfessionalProjectRepository {
	return &profProjRepositoryImpl{db: db}
}

func (r *profProjRepositoryImpl) FindByUser(userID int) ([]professional_project_model.ProfessionalProject, error) {
	rows := []professional_project_model.ProfessionalProject{}
	err := r.db.Select(&rows,
		`SELECT professionalProjectID, userID, title, company, summary,
		        thumbnailGdriveID, thumbnailFileName, orderNo, createdDate
		 FROM ms_professional_project WHERE userID = ? ORDER BY orderNo ASC, professionalProjectID ASC`,
		userID)
	return rows, err
}

func (r *profProjRepositoryImpl) FindByID(projectID int) (professional_project_model.ProfessionalProject, error) {
	var row professional_project_model.ProfessionalProject
	err := r.db.Get(&row,
		`SELECT professionalProjectID, userID, title, company, summary,
		        thumbnailGdriveID, thumbnailFileName, orderNo, createdDate
		 FROM ms_professional_project WHERE professionalProjectID = ? LIMIT 1`,
		projectID)
	return row, err
}

func (r *profProjRepositoryImpl) FindByIDWithOwner(projectID int) (professional_project_model.ProfessionalProjectWithOwner, error) {
	var row professional_project_model.ProfessionalProjectWithOwner
	err := r.db.Get(&row,
		`SELECT p.professionalProjectID, p.userID, p.title, p.company, p.summary,
		        p.thumbnailGdriveID, p.thumbnailFileName, p.orderNo, p.createdDate,
		        u.phone AS ownerPhone
		 FROM ms_professional_project p
		 INNER JOIN ms_user u ON u.userID = p.userID
		 WHERE p.professionalProjectID = ? LIMIT 1`,
		projectID)
	return row, err
}

func (r *profProjRepositoryImpl) Create(p professional_project_model.ProfessionalProject) (int, error) {
	result, err := r.db.Exec(
		`INSERT INTO ms_professional_project (userID, title, company, summary, orderNo) VALUES (?, ?, ?, ?, ?)`,
		p.UserID, p.Title, p.Company, p.Summary, p.OrderNo)
	if err != nil {
		return 0, err
	}
	id, _ := result.LastInsertId()
	return int(id), nil
}

func (r *profProjRepositoryImpl) UpdateThumbnail(projectID int, fileName, gdriveID string) error {
	_, err := r.db.Exec(
		`UPDATE ms_professional_project SET thumbnailFileName = ?, thumbnailGdriveID = ? WHERE professionalProjectID = ?`,
		fileName, gdriveID, projectID)
	return err
}

func (r *profProjRepositoryImpl) Delete(projectID int) error {
	_, err := r.db.Exec(`DELETE FROM ms_professional_project WHERE professionalProjectID = ?`, projectID)
	return err
}

func (r *profProjRepositoryImpl) GetFeatures(projectID int) ([]professional_project_model.ProjectFeature, error) {
	rows := []professional_project_model.ProjectFeature{}
	err := r.db.Select(&rows,
		`SELECT featureID, professionalProjectID, title, description, orderNo
		 FROM ms_professional_project_feature WHERE professionalProjectID = ? ORDER BY orderNo ASC, featureID ASC`,
		projectID)
	return rows, err
}

func (r *profProjRepositoryImpl) FindFeatureByID(featureID int) (professional_project_model.ProjectFeature, error) {
	var f professional_project_model.ProjectFeature
	err := r.db.Get(&f,
		`SELECT featureID, professionalProjectID, title, description, orderNo
		 FROM ms_professional_project_feature WHERE featureID = ? LIMIT 1`,
		featureID)
	return f, err
}

func (r *profProjRepositoryImpl) AddFeature(f professional_project_model.ProjectFeature) (int, error) {
	result, err := r.db.Exec(
		`INSERT INTO ms_professional_project_feature (professionalProjectID, title, description, orderNo) VALUES (?, ?, ?, ?)`,
		f.ProfessionalProjectID, f.Title, f.Description, f.OrderNo)
	if err != nil {
		return 0, err
	}
	id, _ := result.LastInsertId()
	return int(id), nil
}

func (r *profProjRepositoryImpl) DeleteFeature(featureID int) error {
	_, err := r.db.Exec(`DELETE FROM ms_professional_project_feature WHERE featureID = ?`, featureID)
	return err
}

func (r *profProjRepositoryImpl) GetFeatureImages(featureID int) ([]professional_project_model.ProjectFeatureImage, error) {
	rows := []professional_project_model.ProjectFeatureImage{}
	err := r.db.Select(&rows,
		`SELECT featureImageID, featureID, gdriveID, fileName, caption, orderNo
		 FROM ms_professional_project_feature_image WHERE featureID = ? ORDER BY orderNo ASC, featureImageID ASC`,
		featureID)
	return rows, err
}

func (r *profProjRepositoryImpl) AddFeatureImage(img professional_project_model.ProjectFeatureImage) (int, error) {
	result, err := r.db.Exec(
		`INSERT INTO ms_professional_project_feature_image (featureID, gdriveID, fileName, caption, orderNo) VALUES (?, ?, ?, ?, ?)`,
		img.FeatureID, img.GdriveID, img.FileName, img.Caption, img.OrderNo)
	if err != nil {
		return 0, err
	}
	id, _ := result.LastInsertId()
	return int(id), nil
}

func (r *profProjRepositoryImpl) DeleteFeatureImage(featureImageID int) (professional_project_model.ProjectFeatureImage, error) {
	var img professional_project_model.ProjectFeatureImage
	err := r.db.Get(&img,
		`SELECT featureImageID, featureID, gdriveID, fileName, caption, orderNo
		 FROM ms_professional_project_feature_image WHERE featureImageID = ? LIMIT 1`,
		featureImageID)
	if err != nil {
		return img, err
	}
	_, err = r.db.Exec(`DELETE FROM ms_professional_project_feature_image WHERE featureImageID = ?`, featureImageID)
	return img, err
}
