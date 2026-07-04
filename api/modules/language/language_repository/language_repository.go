package language_repository

import (
	"portfolio-api/modules/language/language_model"

	"github.com/jmoiron/sqlx"
)

// LanguageRepository handles persistence for languages.
type LanguageRepository interface {
	FindByUser(userID int) ([]language_model.Language, error)
	FindByID(languageID int) (language_model.Language, error)
	Create(language language_model.Language) (int, error)
	Update(language language_model.Language) error
	Delete(languageID int) error
}

type languageRepositoryImpl struct {
	db *sqlx.DB
}

// NewLanguageRepository builds a LanguageRepository.
func NewLanguageRepository(db *sqlx.DB) LanguageRepository {
	return &languageRepositoryImpl{db: db}
}

func (r *languageRepositoryImpl) FindByUser(userID int) ([]language_model.Language, error) {
	languages := []language_model.Language{}
	query := `SELECT languageID, userID, name, level, icon, orderNo, createdDate, editedDate
		FROM ms_language WHERE userID = ? ORDER BY orderNo ASC, languageID ASC`
	err := r.db.Select(&languages, query, userID)
	return languages, err
}

func (r *languageRepositoryImpl) FindByID(languageID int) (language_model.Language, error) {
	var language language_model.Language
	query := `SELECT languageID, userID, name, level, icon, orderNo, createdDate, editedDate
		FROM ms_language WHERE languageID = ? LIMIT 1`
	err := r.db.Get(&language, query, languageID)
	return language, err
}

func (r *languageRepositoryImpl) Create(language language_model.Language) (int, error) {
	query := `INSERT INTO ms_language (userID, name, level, icon, orderNo) VALUES (?, ?, ?, ?, ?)`
	result, err := r.db.Exec(query, language.UserID, language.Name, language.Level, language.Icon, language.OrderNo)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	return int(id), err
}

func (r *languageRepositoryImpl) Update(language language_model.Language) error {
	query := `UPDATE ms_language SET name = ?, level = ?, icon = ?, orderNo = ? WHERE languageID = ?`
	_, err := r.db.Exec(query, language.Name, language.Level, language.Icon, language.OrderNo, language.LanguageID)
	return err
}

func (r *languageRepositoryImpl) Delete(languageID int) error {
	_, err := r.db.Exec(`DELETE FROM ms_language WHERE languageID = ?`, languageID)
	return err
}
