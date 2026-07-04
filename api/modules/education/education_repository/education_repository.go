package education_repository

import (
	"portfolio-api/modules/education/education_model"

	"github.com/jmoiron/sqlx"
)

// EducationRepository handles persistence for education entries and achievements.
type EducationRepository interface {
	FindByUser(userID int) ([]education_model.Education, error)
	FindByID(educationID int) (education_model.Education, error)
	GetAchievements(educationID int) ([]education_model.EducationAchievement, error)
	Create(education education_model.Education, achievements []string) (int, error)
	Update(education education_model.Education, achievements []string) error
	Delete(educationID int) error
}

type educationRepositoryImpl struct {
	db *sqlx.DB
}

// NewEducationRepository builds an EducationRepository.
func NewEducationRepository(db *sqlx.DB) EducationRepository {
	return &educationRepositoryImpl{db: db}
}

func (r *educationRepositoryImpl) FindByUser(userID int) ([]education_model.Education, error) {
	educations := []education_model.Education{}
	query := `SELECT educationID, userID, degree, institution, year, type, orderNo, createdDate, editedDate
		FROM ms_education WHERE userID = ? ORDER BY orderNo ASC, educationID ASC`
	err := r.db.Select(&educations, query, userID)
	return educations, err
}

func (r *educationRepositoryImpl) FindByID(educationID int) (education_model.Education, error) {
	var education education_model.Education
	query := `SELECT educationID, userID, degree, institution, year, type, orderNo, createdDate, editedDate
		FROM ms_education WHERE educationID = ? LIMIT 1`
	err := r.db.Get(&education, query, educationID)
	return education, err
}

func (r *educationRepositoryImpl) GetAchievements(educationID int) ([]education_model.EducationAchievement, error) {
	achievements := []education_model.EducationAchievement{}
	query := `SELECT educationAchievementID, educationID, description, orderNo
		FROM ms_education_achievement WHERE educationID = ? ORDER BY orderNo ASC, educationAchievementID ASC`
	err := r.db.Select(&achievements, query, educationID)
	return achievements, err
}

func (r *educationRepositoryImpl) Create(education education_model.Education, achievements []string) (int, error) {
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

	result, err := tx.Exec(`INSERT INTO ms_education (userID, degree, institution, year, type, orderNo) VALUES (?, ?, ?, ?, ?, ?)`,
		education.UserID, education.Degree, education.Institution, education.Year, education.Type, education.OrderNo)
	if err != nil {
		return 0, err
	}
	id, _ := result.LastInsertId()
	educationID := int(id)

	if err := insertAchievements(tx, educationID, achievements); err != nil {
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}
	committed = true
	return educationID, nil
}

func (r *educationRepositoryImpl) Update(education education_model.Education, achievements []string) error {
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

	if _, err := tx.Exec(`UPDATE ms_education SET degree = ?, institution = ?, year = ?, type = ?, orderNo = ? WHERE educationID = ?`,
		education.Degree, education.Institution, education.Year, education.Type, education.OrderNo, education.EducationID); err != nil {
		return err
	}

	if _, err := tx.Exec(`DELETE FROM ms_education_achievement WHERE educationID = ?`, education.EducationID); err != nil {
		return err
	}

	if err := insertAchievements(tx, education.EducationID, achievements); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	committed = true
	return nil
}

func (r *educationRepositoryImpl) Delete(educationID int) error {
	_, err := r.db.Exec(`DELETE FROM ms_education WHERE educationID = ?`, educationID)
	return err
}

func insertAchievements(tx *sqlx.Tx, educationID int, achievements []string) error {
	for i, description := range achievements {
		if _, err := tx.Exec(`INSERT INTO ms_education_achievement (educationID, description, orderNo) VALUES (?, ?, ?)`,
			educationID, description, i); err != nil {
			return err
		}
	}
	return nil
}
