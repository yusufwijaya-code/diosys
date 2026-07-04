package pricing_dto

// FeatureRequest is a single feature row payload.
type FeatureRequest struct {
	Text       string `json:"text" binding:"required"`
	IsIncluded bool   `json:"isIncluded"`
}

// PricePlanRequest is the create/update payload for a pricing plan.
type PricePlanRequest struct {
	Title           string           `json:"title" binding:"required"`
	Subtitle        string           `json:"subtitle"`
	Price           float64          `json:"price"`
	Currency        string           `json:"currency"`
	BillingPeriod   string           `json:"billingPeriod"`
	Badge           string           `json:"badge"`
	OriginalPrice   *float64         `json:"originalPrice"`
	DiscountPercent *int             `json:"discountPercent"`
	IsFeatured      bool             `json:"isFeatured"`
	FlagActive      *bool            `json:"flagActive"`
	OrderNo         int              `json:"orderNo"`
	Features        []FeatureRequest `json:"features"`
}

// FeatureResponse is a feature with its included flag.
type FeatureResponse struct {
	FeatureID  int    `json:"featureID"`
	Text       string `json:"text"`
	IsIncluded bool   `json:"isIncluded"`
	OrderNo    int    `json:"orderNo"`
}

// PricePlanResponse is the aggregate returned to clients.
type PricePlanResponse struct {
	PlanID          int               `json:"planID"`
	Title           string            `json:"title"`
	Subtitle        string            `json:"subtitle"`
	Price           float64           `json:"price"`
	Currency        string            `json:"currency"`
	BillingPeriod   string            `json:"billingPeriod"`
	Badge           string            `json:"badge"`
	OriginalPrice   *float64          `json:"originalPrice"`
	DiscountPercent *int              `json:"discountPercent"`
	IsFeatured      bool              `json:"isFeatured"`
	FlagActive      bool              `json:"flagActive"`
	OrderNo         int               `json:"orderNo"`
	Features        []FeatureResponse `json:"features"`
}
