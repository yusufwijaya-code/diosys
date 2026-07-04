package pricing_service

import (
	"database/sql"
	"errors"

	"portfolio-api/base/helpers/error_helper"
	"portfolio-api/modules/pricing/pricing_dto"
	"portfolio-api/modules/pricing/pricing_model"
	"portfolio-api/modules/pricing/pricing_repository"

	"gopkg.in/guregu/null.v4"
)

// PricingService exposes pricing plan operations.
type PricingService interface {
	GetPublic() ([]pricing_dto.PricePlanResponse, error)
	GetAll() ([]pricing_dto.PricePlanResponse, error)
	GetByID(planID int) (pricing_dto.PricePlanResponse, error)
	Create(request pricing_dto.PricePlanRequest) (pricing_dto.PricePlanResponse, error)
	Update(planID int, request pricing_dto.PricePlanRequest) (pricing_dto.PricePlanResponse, error)
	Delete(planID int) error
}

type pricingServiceImpl struct {
	repository pricing_repository.PricingRepository
}

// NewPricingService builds a PricingService.
func NewPricingService(repository pricing_repository.PricingRepository) PricingService {
	return &pricingServiceImpl{repository: repository}
}

func (s *pricingServiceImpl) buildResponse(plan pricing_model.PricePlan) (pricing_dto.PricePlanResponse, error) {
	features, err := s.repository.GetFeatures(plan.PlanID)
	if err != nil {
		return pricing_dto.PricePlanResponse{}, error_helper.Internal(err)
	}

	featureResponses := make([]pricing_dto.FeatureResponse, 0, len(features))
	for _, f := range features {
		featureResponses = append(featureResponses, pricing_dto.FeatureResponse{
			FeatureID:  f.FeatureID,
			Text:       f.Text,
			IsIncluded: f.IsIncluded == 1,
			OrderNo:    f.OrderNo,
		})
	}

	var origPrice *float64
	if plan.OriginalPrice.Valid {
		v := plan.OriginalPrice.Float64
		origPrice = &v
	}
	var discPct *int
	if plan.DiscountPercent.Valid {
		v := int(plan.DiscountPercent.Int64)
		discPct = &v
	}

	return pricing_dto.PricePlanResponse{
		PlanID:          plan.PlanID,
		Title:           plan.Title,
		Subtitle:        plan.Subtitle.String,
		Price:           plan.Price,
		Currency:        plan.Currency,
		BillingPeriod:   plan.BillingPeriod.String,
		Badge:           plan.Badge.String,
		OriginalPrice:   origPrice,
		DiscountPercent: discPct,
		IsFeatured:      plan.IsFeatured == 1,
		FlagActive:      plan.FlagActive == 1,
		OrderNo:         plan.OrderNo,
		Features:        featureResponses,
	}, nil
}

func (s *pricingServiceImpl) GetPublic() ([]pricing_dto.PricePlanResponse, error) {
	return s.getList(true)
}

func (s *pricingServiceImpl) GetAll() ([]pricing_dto.PricePlanResponse, error) {
	return s.getList(false)
}

func (s *pricingServiceImpl) getList(activeOnly bool) ([]pricing_dto.PricePlanResponse, error) {
	plans, err := s.repository.FindAll(activeOnly)
	if err != nil {
		return nil, error_helper.Internal(err)
	}
	responses := make([]pricing_dto.PricePlanResponse, 0, len(plans))
	for _, plan := range plans {
		r, err := s.buildResponse(plan)
		if err != nil {
			return nil, err
		}
		responses = append(responses, r)
	}
	return responses, nil
}

func (s *pricingServiceImpl) GetByID(planID int) (pricing_dto.PricePlanResponse, error) {
	plan, err := s.repository.FindByID(planID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return pricing_dto.PricePlanResponse{}, error_helper.NotFound("pricing plan not found")
		}
		return pricing_dto.PricePlanResponse{}, error_helper.Internal(err)
	}
	return s.buildResponse(plan)
}

func (s *pricingServiceImpl) Create(request pricing_dto.PricePlanRequest) (pricing_dto.PricePlanResponse, error) {
	flagActive := 1
	if request.FlagActive != nil && !*request.FlagActive {
		flagActive = 0
	}
	id, err := s.repository.Create(toModel(0, request, flagActive), toFeatures(request.Features))
	if err != nil {
		return pricing_dto.PricePlanResponse{}, error_helper.Internal(err)
	}
	return s.GetByID(id)
}

func (s *pricingServiceImpl) Update(planID int, request pricing_dto.PricePlanRequest) (pricing_dto.PricePlanResponse, error) {
	existing, err := s.repository.FindByID(planID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return pricing_dto.PricePlanResponse{}, error_helper.NotFound("pricing plan not found")
		}
		return pricing_dto.PricePlanResponse{}, error_helper.Internal(err)
	}
	flagActive := existing.FlagActive
	if request.FlagActive != nil {
		if *request.FlagActive {
			flagActive = 1
		} else {
			flagActive = 0
		}
	}
	m := toModel(planID, request, flagActive)
	if err := s.repository.Update(m, toFeatures(request.Features)); err != nil {
		return pricing_dto.PricePlanResponse{}, error_helper.Internal(err)
	}
	return s.GetByID(planID)
}

func (s *pricingServiceImpl) Delete(planID int) error {
	if _, err := s.GetByID(planID); err != nil {
		return err
	}
	if err := s.repository.Delete(planID); err != nil {
		return error_helper.Internal(err)
	}
	return nil
}

func toModel(planID int, r pricing_dto.PricePlanRequest, flagActive int) pricing_model.PricePlan {
	isFeatured := 0
	if r.IsFeatured {
		isFeatured = 1
	}
	currency := r.Currency
	if currency == "" {
		currency = "IDR"
	}
	return pricing_model.PricePlan{
		PlanID:          planID,
		Title:           r.Title,
		Subtitle:        null.NewString(r.Subtitle, r.Subtitle != ""),
		Price:           r.Price,
		Currency:        currency,
		BillingPeriod:   null.NewString(r.BillingPeriod, r.BillingPeriod != ""),
		Badge:           null.NewString(r.Badge, r.Badge != ""),
		OriginalPrice:   pricing_repository.NullFloat(r.OriginalPrice),
		DiscountPercent: pricing_repository.NullInt(r.DiscountPercent),
		IsFeatured:      isFeatured,
		FlagActive:      flagActive,
		OrderNo:         r.OrderNo,
	}
}

func toFeatures(items []pricing_dto.FeatureRequest) []pricing_model.PriceFeature {
	features := make([]pricing_model.PriceFeature, 0, len(items))
	for _, item := range items {
		isIncluded := 0
		if item.IsIncluded {
			isIncluded = 1
		}
		features = append(features, pricing_model.PriceFeature{Text: item.Text, IsIncluded: isIncluded})
	}
	return features
}
