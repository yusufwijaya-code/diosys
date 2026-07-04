package experience_repository

import (
	"portfolio-api/modules/experience/experience_model"

	"github.com/jmoiron/sqlx"
)

// ExperienceRepository handles persistence for experiences and their children.
type ExperienceRepository interface {
	FindByUser(userID int) ([]experience_model.Experience, error)
	FindByID(experienceID int) (experience_model.Experience, error)
	GetTechnologies(experienceID int) ([]experience_model.ExperienceTechnology, error)
	GetResponsibilities(experienceID int) ([]experience_model.ExperienceResponsibility, error)
	Create(experience experience_model.Experience, technologies, responsibilities []string) (int, error)
	Update(experience experience_model.Experience, technologies, responsibilities []string) error
	Delete(experienceID int) error
}

type experienceRepositoryImpl struct {
	db *sqlx.DB
}

// NewExperienceRepository builds an ExperienceRepository.
func NewExperienceRepository(db *sqlx.DB) ExperienceRepository {
	return &experienceRepositoryImpl{db: db}
}

func (r *experienceRepositoryImpl) FindByUser(userID int) ([]experience_model.Experience, error) {
	experiences := []experience_model.Experience{}
	query := `SELECT experienceID, userID, position, company, period, orderNo, createdDate, editedDate
		FROM ms_experience WHERE userID = ? ORDER BY orderNo ASC, experienceID ASC`
	err := r.db.Select(&experiences, query, userID)
	return experiences, err
}

func (r *experienceRepositoryImpl) FindByID(experienceID int) (experience_model.Experience, error) {
	var experience experience_model.Experience
	query := `SELECT experienceID, userID, position, company, period, orderNo, createdDate, editedDate
		FROM ms_experience WHERE experienceID = ? LIMIT 1`
	err := r.db.Get(&experience, query, experienceID)
	return experience, err
}

func (r *experienceRepositoryImpl) GetTechnologies(experienceID int) ([]experience_model.ExperienceTechnology, error) {
	technologies := []experience_model.ExperienceTechnology{}
	query := `SELECT experienceTechnologyID, experienceID, name, orderNo
		FROM ms_experience_technology WHERE experienceID = ? ORDER BY orderNo ASC, experienceTechnologyID ASC`
	err := r.db.Select(&technologies, query, experienceID)
	return technologies, err
}

func (r *experienceRepositoryImpl) GetResponsibilities(experienceID int) ([]experience_model.ExperienceResponsibility, error) {
	responsibilities := []experience_model.ExperienceResponsibility{}
	query := `SELECT experienceResponsibilityID, experienceID, description, orderNo
		FROM ms_experience_responsibility WHERE experienceID = ? ORDER BY orderNo ASC, experienceResponsibilityID ASC`
	err := r.db.Select(&responsibilities, query, experienceID)
	return responsibilities, err
}

func (r *experienceRepositoryImpl) Create(experience experience_model.Experience, technologies, responsibilities []string) (int, error) {
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

	result, err := tx.Exec(`INSERT INTO ms_experience (userID, position, company, period, orderNo) VALUES (?, ?, ?, ?, ?)`,
		experience.UserID, experience.Position, experience.Company, experience.Period, experience.OrderNo)
	if err != nil {
		return 0, err
	}
	id, _ := result.LastInsertId()
	experienceID := int(id)

	if err := insertChildren(tx, experienceID, technologies, responsibilities); err != nil {
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}
	committed = true
	return experienceID, nil
}

func (r *experienceRepositoryImpl) Update(experience experience_model.Experience, technologies, responsibilities []string) error {
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

	if _, err := tx.Exec(`UPDATE ms_experience SET position = ?, company = ?, period = ?, orderNo = ? WHERE experienceID = ?`,
		experience.Position, experience.Company, experience.Period, experience.OrderNo, experience.ExperienceID); err != nil {
		return err
	}

	if _, err := tx.Exec(`DELETE FROM ms_experience_technology WHERE experienceID = ?`, experience.ExperienceID); err != nil {
		return err
	}
	if _, err := tx.Exec(`DELETE FROM ms_experience_responsibility WHERE experienceID = ?`, experience.ExperienceID); err != nil {
		return err
	}

	if err := insertChildren(tx, experience.ExperienceID, technologies, responsibilities); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	committed = true
	return nil
}

func (r *experienceRepositoryImpl) Delete(experienceID int) error {
	_, err := r.db.Exec(`DELETE FROM ms_experience WHERE experienceID = ?`, experienceID)
	return err
}

func insertChildren(tx *sqlx.Tx, experienceID int, technologies, responsibilities []string) error {
	for i, name := range technologies {
		if _, err := tx.Exec(`INSERT INTO ms_experience_technology (experienceID, name, orderNo) VALUES (?, ?, ?)`,
			experienceID, name, i); err != nil {
			return err
		}
	}
	for i, description := range responsibilities {
		if _, err := tx.Exec(`INSERT INTO ms_experience_responsibility (experienceID, description, orderNo) VALUES (?, ?, ?)`,
			experienceID, description, i); err != nil {
			return err
		}
	}
	return nil
}
