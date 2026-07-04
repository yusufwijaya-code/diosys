package service_dto

// ServiceRequest is the create/update payload for an agency service.
type ServiceRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	OrderNo     int    `json:"orderNo"`
	FlagActive  *bool  `json:"flagActive"`
}
