package user_model

import "gopkg.in/guregu/null.v4"

const TableName = "ms_user"

// User represents a Diosys account. Every account is a developer profile; the
// whitelisted administrator additionally has isAdmin set.
type User struct {
	UserID         int         `db:"userID"         json:"userID"`
	Username       string      `db:"username"       json:"username"`
	Email          string      `db:"email"          json:"email"`
	GoogleSub      null.String `db:"googleSub"      json:"-"`
	FullName       string      `db:"fullName"       json:"fullName"`
	JobTitle       null.String `db:"jobTitle"       json:"jobTitle"`
	Intro          null.String `db:"intro"          json:"intro"`
	Bio            null.String `db:"bio"            json:"bio"`
	Specialization null.String `db:"specialization" json:"specialization"`
	Phone          null.String `db:"phone"          json:"phone"`
	Website        null.String `db:"website"        json:"website"`
	Location       null.String `db:"location"       json:"location"`
	PhotoFileName  null.String `db:"photoFileName"  json:"photoFileName"`
	PhotoGdriveID  null.String `db:"photoGdriveID"  json:"photoGdriveID"`
	UserRoleID     null.Int    `db:"userRoleID"     json:"userRoleID"`
	IsAdmin        int         `db:"isAdmin"        json:"isAdmin"`
	FlagActive     int         `db:"flagActive"     json:"flagActive"`
	OrderNo        int         `db:"orderNo"        json:"orderNo"`
	CreatedDate    null.Time   `db:"createdDate"    json:"createdDate"`
	EditedDate     null.Time   `db:"editedDate"     json:"editedDate"`
}
