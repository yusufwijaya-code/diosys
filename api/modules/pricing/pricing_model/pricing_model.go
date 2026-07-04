package pricing_model

import "gopkg.in/guregu/null.v4"

const (
	TablePlanName    = "ms_price_plan"
	TableFeatureName = "ms_price_feature"
)

// PricePlan is a single pricing tier shown on the website.
type PricePlan struct {
	PlanID          int          `db:"planID"          json:"planID"`
	Title           string       `db:"title"           json:"title"`
	Subtitle        null.String  `db:"subtitle"        json:"subtitle"`
	Price           float64      `db:"price"           json:"price"`
	Currency        string       `db:"currency"        json:"currency"`
	BillingPeriod   null.String  `db:"billingPeriod"   json:"billingPeriod"`
	Badge           null.String  `db:"badge"           json:"badge"`
	OriginalPrice   null.Float   `db:"originalPrice"   json:"originalPrice"`
	DiscountPercent null.Int     `db:"discountPercent" json:"discountPercent"`
	IsFeatured      int          `db:"isFeatured"      json:"isFeatured"`
	FlagActive      int          `db:"flagActive"      json:"flagActive"`
	OrderNo         int          `db:"orderNo"         json:"orderNo"`
	CreatedDate     null.Time    `db:"createdDate"     json:"createdDate"`
	EditedDate      null.Time    `db:"editedDate"      json:"editedDate"`
}

// PriceFeature is a single bullet point of a pricing plan.
type PriceFeature struct {
	FeatureID  int    `db:"featureID"  json:"featureID"`
	PlanID     int    `db:"planID"     json:"planID"`
	Text       string `db:"text"       json:"text"`
	IsIncluded int    `db:"isIncluded" json:"isIncluded"`
	OrderNo    int    `db:"orderNo"    json:"orderNo"`
}
