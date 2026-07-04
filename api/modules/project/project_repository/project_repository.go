package project_repository

import (
	"portfolio-api/modules/project/project_model"

	"github.com/jmoiron/sqlx"
	"gopkg.in/guregu/null.v4"
)

const projectColumns = `projectID, userID, title, summary, body, client, projectLink, repoLink,
	projectStatusID, isFeatured, thumbnailFileName, thumbnailGdriveID, orderNo, createdDate, editedDate`

// ProjectRepository handles persistence for projects and their children.
type ProjectRepository interface {
	FindAllPublic() ([]project_model.ProjectWithOwner, error)
	FindByUser(userID int) ([]project_model.Project, error)
	FindByID(projectID int) (project_model.Project, error)
	GetFeatures(projectID int) ([]project_model.ProjectFeature, error)
	GetTechnologies(projectID int) ([]project_model.ProjectTechnology, error)
	GetImages(projectID int) ([]project_model.ProjectImage, error)
	Create(project project_model.Project, features, technologies []string) (int, error)
	Update(project project_model.Project, features, technologies []string) error
	Delete(projectID int) error
	UpdateThumbnail(projectID int, fileName, gdriveID string) error
	AddImage(image project_model.ProjectImage) (int, error)
	FindImageByID(projectImageID int) (project_model.ProjectImage, error)
	DeleteImage(projectImageID int) error
}

type projectRepositoryImpl struct {
	db *sqlx.DB
}

// NewProjectRepository builds a ProjectRepository.
func NewProjectRepository(db *sqlx.DB) ProjectRepository {
	return &projectRepositoryImpl{db: db}
}

func (r *projectRepositoryImpl) FindAllPublic() ([]project_model.ProjectWithOwner, error) {
	projects := []project_model.ProjectWithOwner{}
	query := `SELECT p.projectID, p.userID, p.title, p.summary, p.body, p.client, p.projectLink,
		p.repoLink, p.projectStatusID, p.isFeatured, p.thumbnailFileName, p.thumbnailGdriveID,
		p.orderNo, p.createdDate, p.editedDate, u.username AS ownerUsername, u.fullName AS ownerFullName
		FROM ms_project p
		INNER JOIN ms_user u ON u.userID = p.userID
		WHERE u.flagActive = 1
		ORDER BY p.isFeatured DESC, p.orderNo ASC, p.projectID DESC`
	err := r.db.Select(&projects, query)
	return projects, err
}

func (r *projectRepositoryImpl) FindByUser(userID int) ([]project_model.Project, error) {
	projects := []project_model.Project{}
	query := `SELECT ` + projectColumns + ` FROM ms_project WHERE userID = ?
		ORDER BY orderNo ASC, projectID DESC`
	err := r.db.Select(&projects, query, userID)
	return projects, err
}

func (r *projectRepositoryImpl) FindByID(projectID int) (project_model.Project, error) {
	var project project_model.Project
	query := `SELECT ` + projectColumns + ` FROM ms_project WHERE projectID = ? LIMIT 1`
	err := r.db.Get(&project, query, projectID)
	return project, err
}

func (r *projectRepositoryImpl) GetFeatures(projectID int) ([]project_model.ProjectFeature, error) {
	features := []project_model.ProjectFeature{}
	query := `SELECT projectFeatureID, projectID, text, orderNo FROM ms_project_feature
		WHERE projectID = ? ORDER BY orderNo ASC, projectFeatureID ASC`
	err := r.db.Select(&features, query, projectID)
	return features, err
}

func (r *projectRepositoryImpl) GetTechnologies(projectID int) ([]project_model.ProjectTechnology, error) {
	technologies := []project_model.ProjectTechnology{}
	query := `SELECT projectTechnologyID, projectID, name, orderNo FROM ms_project_technology
		WHERE projectID = ? ORDER BY orderNo ASC, projectTechnologyID ASC`
	err := r.db.Select(&technologies, query, projectID)
	return technologies, err
}

func (r *projectRepositoryImpl) GetImages(projectID int) ([]project_model.ProjectImage, error) {
	images := []project_model.ProjectImage{}
	query := `SELECT projectImageID, projectID, fileName, gdriveID, caption, displayOrder
		FROM ms_project_image WHERE projectID = ? ORDER BY displayOrder ASC, projectImageID ASC`
	err := r.db.Select(&images, query, projectID)
	return images, err
}

func (r *projectRepositoryImpl) Create(project project_model.Project, features, technologies []string) (int, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return 0, err
	}
	committed := false
	defer func() {
		if !committed {
			_ = tx.Rollback()
		}
	}()

	result, err := tx.Exec(`INSERT INTO ms_project
		(userID, title, summary, body, client, projectLink, repoLink, projectStatusID, isFeatured, orderNo)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		project.UserID, project.Title, project.Summary, project.Body, project.Client,
		project.ProjectLink, project.RepoLink, project.ProjectStatusID, project.IsFeatured, project.OrderNo)
	if err != nil {
		return 0, err
	}
	id, _ := result.LastInsertId()
	projectID := int(id)

	if err := insertChildren(tx, projectID, features, technologies); err != nil {
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}
	committed = true
	return projectID, nil
}

func (r *projectRepositoryImpl) Update(project project_model.Project, features, technologies []string) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}
	committed := false
	defer func() {
		if !committed {
			_ = tx.Rollback()
		}
	}()

	if _, err := tx.Exec(`UPDATE ms_project SET title = ?, summary = ?, body = ?, client = ?,
		projectLink = ?, repoLink = ?, projectStatusID = ?, isFeatured = ?, orderNo = ?
		WHERE projectID = ?`,
		project.Title, project.Summary, project.Body, project.Client, project.ProjectLink,
		project.RepoLink, project.ProjectStatusID, project.IsFeatured, project.OrderNo, project.ProjectID); err != nil {
		return err
	}

	if _, err := tx.Exec(`DELETE FROM ms_project_feature WHERE projectID = ?`, project.ProjectID); err != nil {
		return err
	}
	if _, err := tx.Exec(`DELETE FROM ms_project_technology WHERE projectID = ?`, project.ProjectID); err != nil {
		return err
	}

	if err := insertChildren(tx, project.ProjectID, features, technologies); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	committed = true
	return nil
}

func (r *projectRepositoryImpl) Delete(projectID int) error {
	_, err := r.db.Exec(`DELETE FROM ms_project WHERE projectID = ?`, projectID)
	return err
}

func (r *projectRepositoryImpl) UpdateThumbnail(projectID int, fileName, gdriveID string) error {
	query := `UPDATE ms_project SET thumbnailFileName = ?, thumbnailGdriveID = ? WHERE projectID = ?`
	_, err := r.db.Exec(query, null.NewString(fileName, fileName != ""), null.NewString(gdriveID, gdriveID != ""), projectID)
	return err
}

func (r *projectRepositoryImpl) AddImage(image project_model.ProjectImage) (int, error) {
	query := `INSERT INTO ms_project_image (projectID, fileName, gdriveID, caption, displayOrder)
		VALUES (?, ?, ?, ?, ?)`
	result, err := r.db.Exec(query, image.ProjectID, image.FileName, image.GdriveID, image.Caption, image.DisplayOrder)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	return int(id), err
}

func (r *projectRepositoryImpl) FindImageByID(projectImageID int) (project_model.ProjectImage, error) {
	var image project_model.ProjectImage
	query := `SELECT projectImageID, projectID, fileName, gdriveID, caption, displayOrder
		FROM ms_project_image WHERE projectImageID = ? LIMIT 1`
	err := r.db.Get(&image, query, projectImageID)
	return image, err
}

func (r *projectRepositoryImpl) DeleteImage(projectImageID int) error {
	_, err := r.db.Exec(`DELETE FROM ms_project_image WHERE projectImageID = ?`, projectImageID)
	return err
}

func insertChildren(tx *sqlx.Tx, projectID int, features, technologies []string) error {
	for i, text := range features {
		if _, err := tx.Exec(`INSERT INTO ms_project_feature (projectID, text, orderNo) VALUES (?, ?, ?)`,
			projectID, text, i); err != nil {
			return err
		}
	}
	for i, name := range technologies {
		if _, err := tx.Exec(`INSERT INTO ms_project_technology (projectID, name, orderNo) VALUES (?, ?, ?)`,
			projectID, name, i); err != nil {
			return err
		}
	}
	return nil
}
