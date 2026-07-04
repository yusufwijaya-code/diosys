package auth_dto

import "portfolio-api/modules/user/user_dto"

// LoginResponse carries the issued access token and the authenticated user.
type LoginResponse struct {
	AccessToken string                `json:"accessToken"`
	TokenType   string                `json:"tokenType"`
	ExpiresIn   int                   `json:"expiresIn"`
	User        user_dto.UserResponse `json:"user"`
}
