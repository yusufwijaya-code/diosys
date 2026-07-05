package project_model

import "gopkg.in/guregu/null.v4"

const (
	TableName           = "ms_project"
	TableFeatureName    = "ms_project_feature"
	TableTechnologyName = "ms_project_technology"
	TableImageName      = "ms_project_image"
)

// Project represents a developer's portfolio project entry.
type Project struct {
	ProjectID         int         `db:"projectID"         json:"projectID"`
	UserID            int         `db:"userID"            json:"userID"`
	Title             string      `db:"title"             json:"title"`
	Summary           null.String `db:"summary"           json:"summary"`
	Body              null.String `db:"body"              json:"body"`
	Client            null.String `db:"client"            json:"client"`
	ProjectLink       null.String `db:"projectLink"       json:"projectLink"`
	RepoLink          null.String `db:"repoLink"          json:"repoLink"`
	ProjectStatusID   null.Int    `db:"projectStatusID"   json:"projectStatusID"`
	IsFeatured        int         `db:"isFeatured"        json:"isFeatured"`
	ThumbnailFileName null.String `db:"thumbnailFileName" json:"thumbnailFileName"`
	ThumbnailGdriveID null.String `db:"thumbnailGdriveID" json:"thumbnailGdriveID"`
	OrderNo           int         `db:"orderNo"           json:"orderNo"`
	CreatedDate       null.Time   `db:"createdDate"       json:"createdDate"`
	EditedDate        null.Time   `db:"editedDate"        json:"editedDate"`
}

// ProjectFeature is a single feature bullet of a project.
type ProjectFeature struct {
	ProjectFeatureID int         `db:"projectFeatureID" json:"projectFeatureID"`
	ProjectID        int         `db:"projectID"        json:"projectID"`
	Text             string      `db:"text"             json:"text"`
	Description      null.String `db:"description"      json:"description"`
	OrderNo          int         `db:"orderNo"          json:"orderNo"`
}

// ProjectFeatureImage is a gallery image attached to a feature.
type ProjectFeatureImage struct {
	ProjectFeatureImageID int         `db:"projectFeatureImageID" json:"projectFeatureImageID"`
	ProjectFeatureID      int         `db:"projectFeatureID"      json:"projectFeatureID"`
	GdriveID              string      `db:"gdriveID"              json:"gdriveID"`
	FileName              string      `db:"fileName"              json:"fileName"`
	Caption               null.String `db:"caption"               json:"caption"`
	OrderNo               int         `db:"orderNo"               json:"orderNo"`
}

// ProjectTechnology is a technology tag attached to a project.
type ProjectTechnology struct {
	ProjectTechnologyID int    `db:"projectTechnologyID" json:"projectTechnologyID"`
	ProjectID           int    `db:"projectID"           json:"projectID"`
	Name                string `db:"name"                json:"name"`
	OrderNo             int    `db:"orderNo"             json:"orderNo"`
}

// ProjectImage is a gallery image attached to a project.
type ProjectImage struct {
	ProjectImageID int         `db:"projectImageID" json:"projectImageID"`
	ProjectID      int         `db:"projectID"      json:"projectID"`
	FileName       null.String `db:"fileName"       json:"fileName"`
	GdriveID       null.String `db:"gdriveID"       json:"gdriveID"`
	Caption        null.String `db:"caption"        json:"caption"`
	DisplayOrder   int         `db:"displayOrder"   json:"displayOrder"`
}

// ProjectWithOwner is a project joined with its developer for public listings.
type ProjectWithOwner struct {
	Project
	OwnerUsername string      `db:"ownerUsername" json:"ownerUsername"`
	OwnerFullName string      `db:"ownerFullName" json:"ownerFullName"`
	OwnerPhone    null.String `db:"ownerPhone"    json:"ownerPhone"`
}
