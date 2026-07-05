package testimonial_service

import (
	"database/sql"
	"errors"
	"mime/multipart"

	"portfolio-api/base/helpers/error_helper"
	"portfolio-api/base/helpers/gdrive_helper"
	"portfolio-api/modules/testimonial/testimonial_dto"
	"portfolio-api/modules/testimonial/testimonial_model"
	"portfolio-api/modules/testimonial/testimonial_repository"

	"gopkg.in/guregu/null.v4" //nolint:depguard
)

type TestimonialService interface {
	GetPublic() ([]testimonial_dto.TestimonialResponse, error)
	GetAll() ([]testimonial_dto.TestimonialResponse, error)
	GetByID(id int) (testimonial_dto.TestimonialResponse, error)
	Create(req testimonial_dto.TestimonialRequest) (testimonial_dto.TestimonialResponse, error)
	Update(id int, req testimonial_dto.TestimonialRequest) (testimonial_dto.TestimonialResponse, error)
	UploadPhoto(id int, file *multipart.FileHeader) (testimonial_dto.TestimonialResponse, error)
	Delete(id int) error
}

type testimonialServiceImpl struct {
	repo   testimonial_repository.TestimonialRepository
	gdrive *gdrive_helper.Client
}

func NewTestimonialService(repo testimonial_repository.TestimonialRepository, gdrive *gdrive_helper.Client) TestimonialService {
	return &testimonialServiceImpl{repo: repo, gdrive: gdrive}
}

func toResponse(t testimonial_model.Testimonial) testimonial_dto.TestimonialResponse {
	return testimonial_dto.TestimonialResponse{
		TestimonialID:   t.TestimonialID,
		ClientName:      t.ClientName,
		ClientRole:      t.ClientRole.ValueOrZero(),
		ClientCompany:   t.ClientCompany.ValueOrZero(),
		TestimonialText: t.TestimonialText,
		Rating:          t.Rating,
		PhotoURL:        gdrive_helper.PublicURL(t.PhotoGdriveID.ValueOrZero()),
		FlagActive:      t.FlagActive,
		OrderNo:         t.OrderNo,
	}
}

func (s *testimonialServiceImpl) GetPublic() ([]testimonial_dto.TestimonialResponse, error) {
	rows, err := s.repo.FindAll(true)
	if err != nil {
		return nil, error_helper.Internal(err)
	}
	out := make([]testimonial_dto.TestimonialResponse, len(rows))
	for i, r := range rows {
		out[i] = toResponse(r)
	}
	return out, nil
}

func (s *testimonialServiceImpl) GetAll() ([]testimonial_dto.TestimonialResponse, error) {
	rows, err := s.repo.FindAll(false)
	if err != nil {
		return nil, error_helper.Internal(err)
	}
	out := make([]testimonial_dto.TestimonialResponse, len(rows))
	for i, r := range rows {
		out[i] = toResponse(r)
	}
	return out, nil
}

func (s *testimonialServiceImpl) GetByID(id int) (testimonial_dto.TestimonialResponse, error) {
	t, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return testimonial_dto.TestimonialResponse{}, error_helper.NotFound("testimonial not found")
		}
		return testimonial_dto.TestimonialResponse{}, error_helper.Internal(err)
	}
	return toResponse(t), nil
}

func flagVal(v *bool, fallback int) int {
	if v == nil {
		return fallback
	}
	if *v {
		return 1
	}
	return 0
}

func ratingVal(v int) int {
	if v < 1 || v > 5 {
		return 5
	}
	return v
}

func (s *testimonialServiceImpl) Create(req testimonial_dto.TestimonialRequest) (testimonial_dto.TestimonialResponse, error) {
	id, err := s.repo.Create(testimonial_model.Testimonial{
		ClientName:      req.ClientName,
		ClientRole:      null.NewString(req.ClientRole, req.ClientRole != ""),
		ClientCompany:   null.NewString(req.ClientCompany, req.ClientCompany != ""),
		TestimonialText: req.TestimonialText,
		Rating:          ratingVal(req.Rating),
		FlagActive:      flagVal(req.FlagActive, 1),
		OrderNo:         req.OrderNo,
	})
	if err != nil {
		return testimonial_dto.TestimonialResponse{}, error_helper.Internal(err)
	}
	return s.GetByID(id)
}

func (s *testimonialServiceImpl) Update(id int, req testimonial_dto.TestimonialRequest) (testimonial_dto.TestimonialResponse, error) {
	existing, err := s.GetByID(id)
	if err != nil {
		return testimonial_dto.TestimonialResponse{}, err
	}
	if err := s.repo.Update(testimonial_model.Testimonial{
		TestimonialID:   id,
		ClientName:      req.ClientName,
		ClientRole:      null.NewString(req.ClientRole, req.ClientRole != ""),
		ClientCompany:   null.NewString(req.ClientCompany, req.ClientCompany != ""),
		TestimonialText: req.TestimonialText,
		Rating:          ratingVal(req.Rating),
		FlagActive:      flagVal(req.FlagActive, existing.FlagActive),
		OrderNo:         req.OrderNo,
	}); err != nil {
		return testimonial_dto.TestimonialResponse{}, error_helper.Internal(err)
	}
	return s.GetByID(id)
}

func (s *testimonialServiceImpl) UploadPhoto(id int, file *multipart.FileHeader) (testimonial_dto.TestimonialResponse, error) {
	if _, err := s.GetByID(id); err != nil {
		return testimonial_dto.TestimonialResponse{}, err
	}
	if s.gdrive == nil {
		return testimonial_dto.TestimonialResponse{}, error_helper.Internal(errors.New("google drive not configured"))
	}
	result, err := s.gdrive.UploadImage(file, "testimonial")
	if err != nil {
		return testimonial_dto.TestimonialResponse{}, error_helper.Internal(err)
	}
	if err := s.repo.UpdatePhoto(id, result.GdriveID, result.FileName); err != nil {
		return testimonial_dto.TestimonialResponse{}, error_helper.Internal(err)
	}
	return s.GetByID(id)
}

func (s *testimonialServiceImpl) Delete(id int) error {
	if _, err := s.GetByID(id); err != nil {
		return err
	}
	return error_helper.Internal(s.repo.Delete(id))
}
