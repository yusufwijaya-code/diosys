package skill_model

import "gopkg.in/guregu/null.v4"

const TableName = "ms_skill"

// Skill represents a single skill entry of a developer.
type Skill struct {
	SkillID     int       `db:"skillID"     json:"skillID"`
	UserID      int       `db:"userID"      json:"userID"`
	Name        string    `db:"name"        json:"name"`
	Level       string    `db:"level"       json:"level"`
	Category    string    `db:"category"    json:"category"`
	OrderNo     int       `db:"orderNo"     json:"orderNo"`
	CreatedDate null.Time `db:"createdDate" json:"createdDate"`
	EditedDate  null.Time `db:"editedDate"  json:"editedDate"`
}
