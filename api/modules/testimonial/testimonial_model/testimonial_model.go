package testimonial_model

import "gopkg.in/guregu/null.v4"

const TableName = "ms_testimonial"

type Testimonial struct {
	TestimonialID   int         `db:"testimonialID"   json:"testimonialID"`
	ClientName      string      `db:"clientName"      json:"clientName"`
	ClientRole      null.String `db:"clientRole"      json:"clientRole"`
	ClientCompany   null.String `db:"clientCompany"   json:"clientCompany"`
	TestimonialText string      `db:"testimonialText" json:"testimonialText"`
	Rating          int         `db:"rating"          json:"rating"`
	PhotoGdriveID   null.String `db:"photoGdriveID"   json:"photoGdriveID"`
	PhotoFileName   null.String `db:"photoFileName"   json:"photoFileName"`
	FlagActive      int         `db:"flagActive"      json:"flagActive"`
	OrderNo         int         `db:"orderNo"         json:"orderNo"`
	CreatedDate     null.Time   `db:"createdDate"     json:"createdDate"`
	EditedDate      null.Time   `db:"editedDate"      json:"editedDate"`
}
