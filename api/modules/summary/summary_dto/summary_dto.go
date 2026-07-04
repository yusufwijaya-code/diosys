package summary_dto

import "portfolio-api/modules/summary/summary_model"

// SummaryRequest is the upsert payload for a developer's summary and its children.
type SummaryRequest struct {
	Content string               `json:"content" binding:"required"`
	Stats   []SummaryStatRequest `json:"stats"`
	Facts   []SummaryFactRequest `json:"facts"`
}

// SummaryStatRequest is a single statistic payload.
type SummaryStatRequest struct {
	Number string `json:"number"`
	Label  string `json:"label"`
}

// SummaryFactRequest is a single fact payload.
type SummaryFactRequest struct {
	Icon string `json:"icon"`
	Text string `json:"text"`
}

// SummaryResponse is the full summary aggregate returned to clients.
type SummaryResponse struct {
	SummaryID int                         `json:"summaryID"`
	UserID    int                         `json:"userID"`
	Content   string                      `json:"content"`
	Stats     []summary_model.SummaryStat `json:"stats"`
	Facts     []summary_model.SummaryFact `json:"facts"`
}
