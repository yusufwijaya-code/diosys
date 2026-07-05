package professional_project_dto

// ProfessionalProjectRequest is the create payload.
type ProfessionalProjectRequest struct {
	Title   string `json:"title"   binding:"required"`
	Company string `json:"company" binding:"required"`
	Summary string `json:"summary"`
	OrderNo int    `json:"orderNo"`
}

// ProjectFeatureRequest is the create payload for a feature.
type ProjectFeatureRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	OrderNo     int    `json:"orderNo"`
}

// ProjectFeatureImageResponse is a single image attached to a feature.
type ProjectFeatureImageResponse struct {
	FeatureImageID int    `json:"featureImageID"`
	Url            string `json:"url"`
	Caption        string `json:"caption"`
	OrderNo        int    `json:"orderNo"`
}

// ProjectFeatureResponse is a feature with its images.
type ProjectFeatureResponse struct {
	FeatureID   int                           `json:"featureID"`
	Title       string                        `json:"title"`
	Description string                        `json:"description"`
	Images      []ProjectFeatureImageResponse `json:"images"`
	OrderNo     int                           `json:"orderNo"`
}

// ProfessionalProjectCardResponse is the compact card view.
type ProfessionalProjectCardResponse struct {
	ProfessionalProjectID int    `json:"professionalProjectID"`
	Title                 string `json:"title"`
	Company               string `json:"company"`
	Summary               string `json:"summary"`
	ThumbnailUrl          string `json:"thumbnailUrl"`
	OrderNo               int    `json:"orderNo"`
}

// ProfessionalProjectResponse is the full detail view.
type ProfessionalProjectResponse struct {
	ProfessionalProjectID int                      `json:"professionalProjectID"`
	UserID                int                      `json:"userID"`
	Title                 string                   `json:"title"`
	Company               string                   `json:"company"`
	Summary               string                   `json:"summary"`
	ThumbnailUrl          string                   `json:"thumbnailUrl"`
	Features              []ProjectFeatureResponse `json:"features"`
	OwnerPhone            string                   `json:"ownerPhone"`
	OrderNo               int                      `json:"orderNo"`
}
