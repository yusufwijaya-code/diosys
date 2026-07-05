package user_repository

import (
	"portfolio-api/modules/user/user_model"

	"github.com/jmoiron/sqlx"
	"gopkg.in/guregu/null.v4"
)

const userColumns = `userID, username, email, googleSub, fullName, jobTitle, intro, bio,
	specialization, phone, website, githubUrl, linkedinUrl, instagramUrl, cvFileName, cvGdriveID,
	location, photoFileName, photoGdriveID, userRoleID, isAdmin, flagActive, orderNo, createdDate, editedDate`

// UserRepository handles persistence for Diosys accounts / developer profiles.
type UserRepository interface {
	FindByID(userID int) (user_model.User, error)
	FindByEmail(email string) (user_model.User, error)
	FindByUsername(username string) (user_model.User, error)
	FindAllDevelopers() ([]user_model.User, error)
	Count() (int, error)
	Create(user user_model.User) (int, error)
	Update(user user_model.User) error
	Delete(userID int) error
	UpdateGoogleSub(userID int, googleSub string) error
	UpdatePhoto(userID int, fileName, gdriveID string) error
	UpdateCV(userID int, fileName, gdriveID string) error
	// CollectGdriveIDsForUser returns every Google Drive file ID belonging to a
	// developer and their entire portfolio so they can be purged on account deletion.
	CollectGdriveIDsForUser(userID int) ([]string, error)
}

type userRepositoryImpl struct {
	db *sqlx.DB
}

// NewUserRepository builds a UserRepository backed by the given database.
func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepositoryImpl{db: db}
}

func (r *userRepositoryImpl) FindByID(userID int) (user_model.User, error) {
	var user user_model.User
	query := `SELECT ` + userColumns + ` FROM ms_user WHERE userID = ? LIMIT 1`
	err := r.db.Get(&user, query, userID)
	return user, err
}

func (r *userRepositoryImpl) FindByEmail(email string) (user_model.User, error) {
	var user user_model.User
	query := `SELECT ` + userColumns + ` FROM ms_user WHERE email = ? LIMIT 1`
	err := r.db.Get(&user, query, email)
	return user, err
}

func (r *userRepositoryImpl) FindByUsername(username string) (user_model.User, error) {
	var user user_model.User
	query := `SELECT ` + userColumns + ` FROM ms_user WHERE username = ? LIMIT 1`
	err := r.db.Get(&user, query, username)
	return user, err
}

func (r *userRepositoryImpl) FindAllDevelopers() ([]user_model.User, error) {
	users := []user_model.User{}
	query := `SELECT ` + userColumns + ` FROM ms_user WHERE flagActive = 1
		ORDER BY orderNo ASC, userID ASC`
	err := r.db.Select(&users, query)
	return users, err
}

func (r *userRepositoryImpl) Count() (int, error) {
	var count int
	err := r.db.Get(&count, "SELECT COUNT(*) FROM ms_user")
	return count, err
}

func (r *userRepositoryImpl) Create(user user_model.User) (int, error) {
	query := `INSERT INTO ms_user
		(username, email, googleSub, fullName, jobTitle, intro, bio, specialization,
		 phone, website, githubUrl, linkedinUrl, instagramUrl, location, userRoleID, isAdmin, flagActive, orderNo)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	result, err := r.db.Exec(query,
		user.Username, user.Email, user.GoogleSub, user.FullName, user.JobTitle, user.Intro,
		user.Bio, user.Specialization, user.Phone, user.Website,
		user.GithubUrl, user.LinkedinUrl, user.InstagramUrl,
		user.Location, user.UserRoleID, user.IsAdmin, user.FlagActive, user.OrderNo)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	return int(id), err
}

func (r *userRepositoryImpl) Update(user user_model.User) error {
	query := `UPDATE ms_user SET
		username = ?, email = ?, fullName = ?, jobTitle = ?, intro = ?, bio = ?,
		specialization = ?, phone = ?, website = ?, githubUrl = ?, linkedinUrl = ?,
		instagramUrl = ?, location = ?, flagActive = ?, orderNo = ?
		WHERE userID = ?`
	_, err := r.db.Exec(query,
		user.Username, user.Email, user.FullName, user.JobTitle, user.Intro, user.Bio,
		user.Specialization, user.Phone, user.Website,
		user.GithubUrl, user.LinkedinUrl, user.InstagramUrl,
		user.Location, user.FlagActive, user.OrderNo, user.UserID)
	return err
}

func (r *userRepositoryImpl) Delete(userID int) error {
	_, err := r.db.Exec(`DELETE FROM ms_user WHERE userID = ?`, userID)
	return err
}

func (r *userRepositoryImpl) UpdateGoogleSub(userID int, googleSub string) error {
	_, err := r.db.Exec(`UPDATE ms_user SET googleSub = ? WHERE userID = ?`,
		null.NewString(googleSub, googleSub != ""), userID)
	return err
}

func (r *userRepositoryImpl) UpdatePhoto(userID int, fileName, gdriveID string) error {
	_, err := r.db.Exec(`UPDATE ms_user SET photoFileName = ?, photoGdriveID = ? WHERE userID = ?`,
		null.NewString(fileName, fileName != ""), null.NewString(gdriveID, gdriveID != ""), userID)
	return err
}

func (r *userRepositoryImpl) UpdateCV(userID int, fileName, gdriveID string) error {
	_, err := r.db.Exec(`UPDATE ms_user SET cvFileName = ?, cvGdriveID = ? WHERE userID = ?`,
		null.NewString(fileName, fileName != ""), null.NewString(gdriveID, gdriveID != ""), userID)
	return err
}

func (r *userRepositoryImpl) CollectGdriveIDsForUser(userID int) ([]string, error) {
	query := `
		SELECT photoGdriveID FROM ms_user
			WHERE userID = ? AND photoGdriveID IS NOT NULL AND photoGdriveID != ''
		UNION ALL
		SELECT cvGdriveID FROM ms_user
			WHERE userID = ? AND cvGdriveID IS NOT NULL AND cvGdriveID != ''
		UNION ALL
		SELECT thumbnailGdriveID FROM ms_project
			WHERE userID = ? AND thumbnailGdriveID IS NOT NULL AND thumbnailGdriveID != ''
		UNION ALL
		SELECT pi.gdriveID FROM ms_project_image pi
			INNER JOIN ms_project p ON p.projectID = pi.projectID
			WHERE p.userID = ? AND pi.gdriveID IS NOT NULL AND pi.gdriveID != ''
		UNION ALL
		SELECT pfi.gdriveID FROM ms_project_feature_image pfi
			INNER JOIN ms_project_feature pf ON pf.projectFeatureID = pfi.projectFeatureID
			INNER JOIN ms_project p ON p.projectID = pf.projectID
			WHERE p.userID = ? AND pfi.gdriveID != ''
		UNION ALL
		SELECT thumbnailGdriveID FROM ms_professional_project
			WHERE userID = ? AND thumbnailGdriveID IS NOT NULL AND thumbnailGdriveID != ''
		UNION ALL
		SELECT ppfi.gdriveID FROM ms_professional_project_feature_image ppfi
			INNER JOIN ms_professional_project_feature ppf ON ppf.featureID = ppfi.featureID
			INNER JOIN ms_professional_project pp ON pp.professionalProjectID = ppf.professionalProjectID
			WHERE pp.userID = ? AND ppfi.gdriveID != ''`

	var ids []string
	err := r.db.Select(&ids, query, userID, userID, userID, userID, userID, userID, userID)
	return ids, err
}
