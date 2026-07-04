package summary_model

import "gopkg.in/guregu/null.v4"

const (
	TableName     = "ms_summary"
	TableStatName = "ms_summary_stat"
	TableFactName = "ms_summary_fact"
)

// Summary holds a developer's "about me" record.
type Summary struct {
	SummaryID   int       `db:"summaryID"   json:"summaryID"`
	UserID      int       `db:"userID"      json:"userID"`
	Content     string    `db:"content"     json:"content"`
	CreatedDate null.Time `db:"createdDate" json:"createdDate"`
	EditedDate  null.Time `db:"editedDate"  json:"editedDate"`
}

// SummaryStat is a highlighted numeric statistic (e.g. "2+ Years Experience").
type SummaryStat struct {
	SummaryStatID int    `db:"summaryStatID" json:"summaryStatID"`
	SummaryID     int    `db:"summaryID"     json:"summaryID"`
	Number        string `db:"number"        json:"number"`
	Label         string `db:"label"         json:"label"`
	OrderNo       int    `db:"orderNo"       json:"orderNo"`
}

// SummaryFact is a short highlighted fact with an icon.
type SummaryFact struct {
	SummaryFactID int    `db:"summaryFactID" json:"summaryFactID"`
	SummaryID     int    `db:"summaryID"     json:"summaryID"`
	Icon          string `db:"icon"          json:"icon"`
	Text          string `db:"text"          json:"text"`
	OrderNo       int    `db:"orderNo"       json:"orderNo"`
}
