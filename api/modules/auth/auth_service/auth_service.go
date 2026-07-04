package auth_service

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"portfolio-api/base/helpers/error_helper"
	"portfolio-api/base/helpers/gdrive_helper"
	"portfolio-api/base/helpers/jwt_helper"
	"portfolio-api/config"
	"portfolio-api/modules/auth/auth_dto"
	"portfolio-api/modules/user/user_dto"
	"portfolio-api/modules/user/user_repository"

	"golang.org/x/oauth2"
	googleoauth "golang.org/x/oauth2/google"
	"google.golang.org/api/idtoken"
)

// AuthService handles Google-based authentication for the CMS administrator.
type AuthService interface {
	BuildAuthURL(state string) string
	HandleCallback(ctx context.Context, code string) (auth_dto.LoginResponse, error)
}

type authServiceImpl struct {
	userRepository user_repository.UserRepository
	config         config.AppConfig
}

// NewAuthService builds an AuthService.
func NewAuthService(userRepository user_repository.UserRepository, cfg config.AppConfig) AuthService {
	return &authServiceImpl{userRepository: userRepository, config: cfg}
}

func (s *authServiceImpl) oauthConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     s.config.GoogleClientID,
		ClientSecret: s.config.GoogleClientSecret,
		RedirectURL:  s.config.GoogleRedirectURI,
		Scopes:       []string{"openid", "email", "profile"},
		Endpoint:     googleoauth.Endpoint,
	}
}

// BuildAuthURL returns the Google consent screen URL for the redirect flow.
func (s *authServiceImpl) BuildAuthURL(state string) string {
	return s.oauthConfig().AuthCodeURL(
		state,
		oauth2.AccessTypeOnline,
		oauth2.SetAuthURLParam("prompt", "select_account"),
	)
}

// HandleCallback exchanges the authorization code, validates the resulting ID
// token, enforces the whitelist and issues a Diosys JWT.
func (s *authServiceImpl) HandleCallback(ctx context.Context, code string) (auth_dto.LoginResponse, error) {
	token, err := s.oauthConfig().Exchange(ctx, code)
	if err != nil {
		return auth_dto.LoginResponse{}, error_helper.Unauthorized("failed to exchange google authorization code")
	}

	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok || rawIDToken == "" {
		return auth_dto.LoginResponse{}, error_helper.Unauthorized("google did not return an id token")
	}

	payload, err := idtoken.Validate(ctx, rawIDToken, s.config.GoogleClientID)
	if err != nil {
		return auth_dto.LoginResponse{}, error_helper.Unauthorized("invalid google credential")
	}

	return s.issueForPayload(payload.Claims)
}

func (s *authServiceImpl) issueForPayload(claims map[string]interface{}) (auth_dto.LoginResponse, error) {
	email, _ := claims["email"].(string)
	email = strings.ToLower(strings.TrimSpace(email))
	if verified, ok := claims["email_verified"].(bool); ok && !verified {
		return auth_dto.LoginResponse{}, error_helper.Forbidden("google email is not verified")
	}

	// Strict whitelist: only the configured administrator may sign in.
	if email == "" || email != strings.ToLower(s.config.AdminEmail) {
		return auth_dto.LoginResponse{}, error_helper.Forbidden("this account is not authorized to sign in")
	}

	user, err := s.userRepository.FindByEmail(email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return auth_dto.LoginResponse{}, error_helper.Forbidden("this account is not authorized to sign in")
		}
		return auth_dto.LoginResponse{}, error_helper.Internal(err)
	}

	if user.FlagActive != 1 {
		return auth_dto.LoginResponse{}, error_helper.Forbidden("account is inactive")
	}

	if sub, ok := claims["sub"].(string); ok && sub != "" && user.GoogleSub.String != sub {
		_ = s.userRepository.UpdateGoogleSub(user.UserID, sub)
	}

	tokenString, err := jwt_helper.GenerateToken(user.UserID, user.Username, s.config.JwtSecret, s.config.JwtTokenLifespanMinutes)
	if err != nil {
		return auth_dto.LoginResponse{}, error_helper.Internal(err)
	}

	return auth_dto.LoginResponse{
		AccessToken: tokenString,
		TokenType:   "Bearer",
		ExpiresIn:   s.config.JwtTokenLifespanMinutes * 60,
		User: user_dto.UserResponse{
			UserID:     user.UserID,
			Username:   user.Username,
			Email:      user.Email,
			FullName:   user.FullName,
			JobTitle:   user.JobTitle.String,
			PhotoUrl:   gdrive_helper.PublicURL(user.PhotoGdriveID.String),
			IsAdmin:    user.IsAdmin,
			FlagActive: user.FlagActive,
		},
	}, nil
}
