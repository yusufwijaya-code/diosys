package user_service

import (
	"database/sql"
	"errors"
	"log"

	"portfolio-api/base/helpers/error_helper"
	"portfolio-api/base/helpers/gdrive_helper"
	"portfolio-api/modules/user/user_dto"
	"portfolio-api/modules/user/user_model"
	"portfolio-api/modules/user/user_repository"

	"gopkg.in/guregu/null.v4"
)

// UserService exposes account / developer-profile operations.
type UserService interface {
	GetByID(userID int) (user_dto.UserResponse, error)
	EnsureWhitelistedAdmin(email, username, fullName string) error
}

type userServiceImpl struct {
	repository user_repository.UserRepository
}

// NewUserService builds a UserService.
func NewUserService(repository user_repository.UserRepository) UserService {
	return &userServiceImpl{repository: repository}
}

func (s *userServiceImpl) GetByID(userID int) (user_dto.UserResponse, error) {
	user, err := s.repository.FindByID(userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user_dto.UserResponse{}, error_helper.NotFound("user not found")
		}
		return user_dto.UserResponse{}, error_helper.Internal(err)
	}

	return user_dto.UserResponse{
		UserID:     user.UserID,
		Username:   user.Username,
		Email:      user.Email,
		FullName:   user.FullName,
		JobTitle:   user.JobTitle.String,
		PhotoUrl:   gdrive_helper.PublicURL(user.PhotoGdriveID.String),
		IsAdmin:    user.IsAdmin,
		FlagActive: user.FlagActive,
	}, nil
}

// EnsureWhitelistedAdmin seeds the single whitelisted administrator (who is also
// the first developer profile) when it does not exist yet.
func (s *userServiceImpl) EnsureWhitelistedAdmin(email, username, fullName string) error {
	_, err := s.repository.FindByEmail(email)
	if err == nil {
		return nil
	}
	if !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	if _, err := s.repository.Create(user_model.User{
		Username:   username,
		Email:      email,
		FullName:   fullName,
		JobTitle:   null.StringFrom("Founder & Full-Stack Developer"),
		Intro:      null.StringFrom("Building premium web, mobile, and AI products at Diosys."),
		UserRoleID: null.IntFrom(1),
		IsAdmin:    1,
		FlagActive: 1,
		OrderNo:    0,
	}); err != nil {
		return err
	}

	log.Printf("Whitelisted admin/developer seeded (email: %s)", email)
	return nil
}
