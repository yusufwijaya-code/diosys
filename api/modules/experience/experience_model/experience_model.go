package experience_model

import "gopkg.in/guregu/null.v4"

const (
	TableName               = "ms_experience"
	TableTechnologyName     = "ms_experience_technology"
	TableResponsibilityName = "ms_experience_responsibility"
)

// Experience represents a developer's work experience entry.
type Experience struct {
	ExperienceID int       `db:"experienceID" json:"experienceID"`
	UserID       int       `db:"userID"       json:"userID"`
	Position     string    `db:"position"     json:"position"`
	Company      string    `db:"company"      json:"company"`
	Period       string    `db:"period"       json:"period"`
	OrderNo      int       `db:"orderNo"      json:"orderNo"`
	CreatedDate  null.Time `db:"createdDate"  json:"createdDate"`
	EditedDate   null.Time `db:"editedDate"   json:"editedDate"`
}

// ExperienceTechnology is a technology tag attached to an experience.
type ExperienceTechnology struct {
	ExperienceTechnologyID int    `db:"experienceTechnologyID" json:"experienceTechnologyID"`
	ExperienceID           int    `db:"experienceID"           json:"experienceID"`
	Name                   string `db:"name"                   json:"name"`
	OrderNo                int    `db:"orderNo"                json:"orderNo"`
}

// ExperienceResponsibility is a single responsibility line of an experience.
type ExperienceResponsibility struct {
	ExperienceResponsibilityID int    `db:"experienceResponsibilityID" json:"experienceResponsibilityID"`
	ExperienceID               int    `db:"experienceID"               json:"experienceID"`
	Description                string `db:"description"                json:"description"`
	OrderNo                    int    `db:"orderNo"                    json:"orderNo"`
}
