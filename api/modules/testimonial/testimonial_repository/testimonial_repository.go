package testimonial_repository

import (
	"portfolio-api/modules/testimonial/testimonial_model"

	"github.com/jmoiron/sqlx"
	"gopkg.in/guregu/null.v4"
)

const cols = `testimonialID, clientName, clientRole, clientCompany, testimonialText, rating, photoGdriveID, photoFileName, flagActive, orderNo, createdDate, editedDate`

type TestimonialRepository interface {
	FindAll(activeOnly bool) ([]testimonial_model.Testimonial, error)
	FindByID(id int) (testimonial_model.Testimonial, error)
	Create(t testimonial_model.Testimonial) (int, error)
	Update(t testimonial_model.Testimonial) error
	UpdatePhoto(id int, gdriveID, fileName string) error
	Delete(id int) error
}

type testimonialRepositoryImpl struct{ db *sqlx.DB }

func NewTestimonialRepository(db *sqlx.DB) TestimonialRepository {
	return &testimonialRepositoryImpl{db: db}
}

func (r *testimonialRepositoryImpl) FindAll(activeOnly bool) ([]testimonial_model.Testimonial, error) {
	rows := []testimonial_model.Testimonial{}
	q := `SELECT ` + cols + ` FROM ms_testimonial`
	if activeOnly {
		q += ` WHERE flagActive = 1`
	}
	q += ` ORDER BY orderNo ASC, testimonialID ASC`
	return rows, r.db.Select(&rows, q)
}

func (r *testimonialRepositoryImpl) FindByID(id int) (testimonial_model.Testimonial, error) {
	var t testimonial_model.Testimonial
	return t, r.db.Get(&t, `SELECT `+cols+` FROM ms_testimonial WHERE testimonialID = ? LIMIT 1`, id)
}

func (r *testimonialRepositoryImpl) Create(t testimonial_model.Testimonial) (int, error) {
	res, err := r.db.Exec(
		`INSERT INTO ms_testimonial (clientName, clientRole, clientCompany, testimonialText, rating, flagActive, orderNo) VALUES (?, ?, ?, ?, ?, ?, ?)`,
		t.ClientName, t.ClientRole, t.ClientCompany, t.TestimonialText, t.Rating, t.FlagActive, t.OrderNo,
	)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	return int(id), err
}

func (r *testimonialRepositoryImpl) Update(t testimonial_model.Testimonial) error {
	_, err := r.db.Exec(
		`UPDATE ms_testimonial SET clientName=?, clientRole=?, clientCompany=?, testimonialText=?, rating=?, flagActive=?, orderNo=? WHERE testimonialID=?`,
		t.ClientName, t.ClientRole, t.ClientCompany, t.TestimonialText, t.Rating, t.FlagActive, t.OrderNo, t.TestimonialID,
	)
	return err
}

func (r *testimonialRepositoryImpl) UpdatePhoto(id int, gdriveID, fileName string) error {
	_, err := r.db.Exec(
		`UPDATE ms_testimonial SET photoGdriveID=?, photoFileName=? WHERE testimonialID=?`,
		null.NewString(gdriveID, gdriveID != ""), null.NewString(fileName, fileName != ""), id,
	)
	return err
}

func (r *testimonialRepositoryImpl) Delete(id int) error {
	_, err := r.db.Exec(`DELETE FROM ms_testimonial WHERE testimonialID=?`, id)
	return err
}
