package education_model

import "gopkg.in/guregu/null.v4"

const (
	TableName            = "ms_education"
	TableAchievementName = "ms_education_achievement"
)

// Education represents a developer's education entry.
type Education struct {
	EducationID int       `db:"educationID" json:"educationID"`
	UserID      int       `db:"userID"      json:"userID"`
	Degree      string    `db:"degree"      json:"degree"`
	Institution string    `db:"institution" json:"institution"`
	Year        string    `db:"year"        json:"year"`
	Type        string    `db:"type"        json:"type"`
	OrderNo     int       `db:"orderNo"     json:"orderNo"`
	CreatedDate null.Time `db:"createdDate" json:"createdDate"`
	EditedDate  null.Time `db:"editedDate"  json:"editedDate"`
}

// EducationAchievement is a single achievement line of an education entry.
type EducationAchievement struct {
	EducationAchievementID int    `db:"educationAchievementID" json:"educationAchievementID"`
	EducationID            int    `db:"educationID"            json:"educationID"`
	Description            string `db:"description"            json:"description"`
	OrderNo                int    `db:"orderNo"                json:"orderNo"`
}
