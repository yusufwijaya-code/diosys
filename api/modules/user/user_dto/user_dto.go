package user_dto

// UserResponse is the authenticated-account representation returned by /auth/me.
type UserResponse struct {
	UserID     int    `json:"userID"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	FullName   string `json:"fullName"`
	JobTitle   string `json:"jobTitle"`
	PhotoUrl   string `json:"photoUrl"`
	IsAdmin    int    `json:"isAdmin"`
	FlagActive int    `json:"flagActive"`
}
