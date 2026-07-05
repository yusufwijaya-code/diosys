package testimonial_dto

// TestimonialRequest is the create/update payload for a testimonial.
type TestimonialRequest struct {
	ClientName      string `json:"clientName" binding:"required"`
	ClientRole      string `json:"clientRole"`
	ClientCompany   string `json:"clientCompany"`
	TestimonialText string `json:"testimonialText" binding:"required"`
	Rating          int    `json:"rating"`
	OrderNo         int    `json:"orderNo"`
	FlagActive      *bool  `json:"flagActive"`
}

// TestimonialResponse is the public-facing testimonial with a derived photo URL.
type TestimonialResponse struct {
	TestimonialID   int    `json:"testimonialID"`
	ClientName      string `json:"clientName"`
	ClientRole      string `json:"clientRole"`
	ClientCompany   string `json:"clientCompany"`
	TestimonialText string `json:"testimonialText"`
	Rating          int    `json:"rating"`
	PhotoURL        string `json:"photoUrl"`
	FlagActive      int    `json:"flagActive"`
	OrderNo         int    `json:"orderNo"`
}
