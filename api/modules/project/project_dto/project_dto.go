package project_dto

// ProjectRequest is the create/update payload for a project.
type ProjectRequest struct {
	Title           string   `json:"title" binding:"required"`
	Summary         string   `json:"summary"`
	Body            string   `json:"body"`
	Client          string   `json:"client"`
	ProjectLink     string   `json:"projectLink"`
	RepoLink        string   `json:"repoLink"`
	ProjectStatusID *int     `json:"projectStatusID"`
	IsFeatured      bool     `json:"isFeatured"`
	OrderNo         int      `json:"orderNo"`
	Features        []string `json:"features"`
	Technologies    []string `json:"technologies"`
}

// ProjectImageRequest carries optional metadata for a gallery image upload.
type ProjectImageRequest struct {
	Caption string `json:"caption"`
}

// ProjectImageResponse is a single gallery image returned to clients.
type ProjectImageResponse struct {
	ProjectImageID int    `json:"projectImageID"`
	FileName       string `json:"fileName"`
	GdriveID       string `json:"gdriveID"`
	Url            string `json:"url"`
	Caption        string `json:"caption"`
	DisplayOrder   int    `json:"displayOrder"`
}

// ProjectResponse is the aggregate returned to clients.
type ProjectResponse struct {
	ProjectID         int                    `json:"projectID"`
	UserID            int                    `json:"userID"`
	OwnerUsername     string                 `json:"ownerUsername"`
	OwnerFullName     string                 `json:"ownerFullName"`
	Title             string                 `json:"title"`
	Summary           string                 `json:"summary"`
	Body              string                 `json:"body"`
	Client            string                 `json:"client"`
	ProjectLink       string                 `json:"projectLink"`
	RepoLink          string                 `json:"repoLink"`
	ProjectStatusID   *int                   `json:"projectStatusID"`
	IsFeatured        bool                   `json:"isFeatured"`
	ThumbnailFileName string                 `json:"thumbnailFileName"`
	ThumbnailGdriveID string                 `json:"thumbnailGdriveID"`
	ThumbnailUrl      string                 `json:"thumbnailUrl"`
	OrderNo           int                    `json:"orderNo"`
	Features          []string               `json:"features"`
	Technologies      []string               `json:"technologies"`
	Images            []ProjectImageResponse `json:"images"`
}
