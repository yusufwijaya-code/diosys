package professional_project_model

import "gopkg.in/guregu/null.v4"

type ProfessionalProject struct {
	ProfessionalProjectID int         `db:"professionalProjectID"`
	UserID                int         `db:"userID"`
	Title                 string      `db:"title"`
	Company               string      `db:"company"`
	Summary               null.String `db:"summary"`
	ThumbnailGdriveID     null.String `db:"thumbnailGdriveID"`
	ThumbnailFileName     null.String `db:"thumbnailFileName"`
	OrderNo               int         `db:"orderNo"`
	CreatedDate           null.Time   `db:"createdDate"`
}

type ProjectFeature struct {
	FeatureID             int         `db:"featureID"`
	ProfessionalProjectID int         `db:"professionalProjectID"`
	Title                 string      `db:"title"`
	Description           null.String `db:"description"`
	OrderNo               int         `db:"orderNo"`
}

type ProjectFeatureImage struct {
	FeatureImageID int         `db:"featureImageID"`
	FeatureID      int         `db:"featureID"`
	GdriveID       string      `db:"gdriveID"`
	FileName       string      `db:"fileName"`
	Caption        null.String `db:"caption"`
	OrderNo        int         `db:"orderNo"`
}
