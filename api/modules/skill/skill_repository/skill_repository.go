package skill_repository

import (
	"portfolio-api/modules/skill/skill_model"

	"github.com/jmoiron/sqlx"
)

// SkillRepository handles persistence for skills.
type SkillRepository interface {
	FindByUser(userID int) ([]skill_model.Skill, error)
	FindByID(skillID int) (skill_model.Skill, error)
	Create(skill skill_model.Skill) (int, error)
	Update(skill skill_model.Skill) error
	Delete(skillID int) error
}

type skillRepositoryImpl struct {
	db *sqlx.DB
}

// NewSkillRepository builds a SkillRepository.
func NewSkillRepository(db *sqlx.DB) SkillRepository {
	return &skillRepositoryImpl{db: db}
}

func (r *skillRepositoryImpl) FindByUser(userID int) ([]skill_model.Skill, error) {
	skills := []skill_model.Skill{}
	query := `SELECT skillID, userID, name, level, category, orderNo, createdDate, editedDate
		FROM ms_skill WHERE userID = ? ORDER BY orderNo ASC, skillID ASC`
	err := r.db.Select(&skills, query, userID)
	return skills, err
}

func (r *skillRepositoryImpl) FindByID(skillID int) (skill_model.Skill, error) {
	var skill skill_model.Skill
	query := `SELECT skillID, userID, name, level, category, orderNo, createdDate, editedDate
		FROM ms_skill WHERE skillID = ? LIMIT 1`
	err := r.db.Get(&skill, query, skillID)
	return skill, err
}

func (r *skillRepositoryImpl) Create(skill skill_model.Skill) (int, error) {
	query := `INSERT INTO ms_skill (userID, name, level, category, orderNo) VALUES (?, ?, ?, ?, ?)`
	result, err := r.db.Exec(query, skill.UserID, skill.Name, skill.Level, skill.Category, skill.OrderNo)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	return int(id), err
}

func (r *skillRepositoryImpl) Update(skill skill_model.Skill) error {
	query := `UPDATE ms_skill SET name = ?, level = ?, category = ?, orderNo = ? WHERE skillID = ?`
	_, err := r.db.Exec(query, skill.Name, skill.Level, skill.Category, skill.OrderNo, skill.SkillID)
	return err
}

func (r *skillRepositoryImpl) Delete(skillID int) error {
	_, err := r.db.Exec(`DELETE FROM ms_skill WHERE skillID = ?`, skillID)
	return err
}
