package setting_repository

import (
	"portfolio-api/modules/setting/setting_model"

	"github.com/jmoiron/sqlx"
)

// SettingRepository handles persistence for system settings.
type SettingRepository interface {
	FindAll() ([]setting_model.Setting, error)
	Upsert(settingKey, settingValue string) error
}

type settingRepositoryImpl struct {
	db *sqlx.DB
}

// NewSettingRepository builds a SettingRepository.
func NewSettingRepository(db *sqlx.DB) SettingRepository {
	return &settingRepositoryImpl{db: db}
}

func (r *settingRepositoryImpl) FindAll() ([]setting_model.Setting, error) {
	settings := []setting_model.Setting{}
	query := `SELECT settingID, settingKey, settingValue, description, editedDate
		FROM ms_system_setting ORDER BY settingKey ASC`
	err := r.db.Select(&settings, query)
	return settings, err
}

func (r *settingRepositoryImpl) Upsert(settingKey, settingValue string) error {
	query := `INSERT INTO ms_system_setting (settingKey, settingValue) VALUES (?, ?)
		ON DUPLICATE KEY UPDATE settingValue = VALUES(settingValue)`
	_, err := r.db.Exec(query, settingKey, settingValue)
	return err
}
