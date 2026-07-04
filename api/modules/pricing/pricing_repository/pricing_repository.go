package pricing_repository

import (
	"portfolio-api/modules/pricing/pricing_model"

	"github.com/jmoiron/sqlx"
	"gopkg.in/guregu/null.v4"
)

const planColumns = `planID, title, subtitle, price, currency, billingPeriod, badge,
	originalPrice, discountPercent, isFeatured, flagActive, orderNo, createdDate, editedDate`

// PricingRepository handles persistence for pricing plans and their features.
type PricingRepository interface {
	FindAll(activeOnly bool) ([]pricing_model.PricePlan, error)
	FindByID(planID int) (pricing_model.PricePlan, error)
	GetFeatures(planID int) ([]pricing_model.PriceFeature, error)
	Create(plan pricing_model.PricePlan, features []pricing_model.PriceFeature) (int, error)
	Update(plan pricing_model.PricePlan, features []pricing_model.PriceFeature) error
	Delete(planID int) error
}

type pricingRepositoryImpl struct {
	db *sqlx.DB
}

// NewPricingRepository builds a PricingRepository.
func NewPricingRepository(db *sqlx.DB) PricingRepository {
	return &pricingRepositoryImpl{db: db}
}

func (r *pricingRepositoryImpl) FindAll(activeOnly bool) ([]pricing_model.PricePlan, error) {
	plans := []pricing_model.PricePlan{}
	query := `SELECT ` + planColumns + ` FROM ms_price_plan`
	if activeOnly {
		query += ` WHERE flagActive = 1`
	}
	query += ` ORDER BY orderNo ASC, planID ASC`
	err := r.db.Select(&plans, query)
	return plans, err
}

func (r *pricingRepositoryImpl) FindByID(planID int) (pricing_model.PricePlan, error) {
	var plan pricing_model.PricePlan
	query := `SELECT ` + planColumns + ` FROM ms_price_plan WHERE planID = ? LIMIT 1`
	err := r.db.Get(&plan, query, planID)
	return plan, err
}

func (r *pricingRepositoryImpl) GetFeatures(planID int) ([]pricing_model.PriceFeature, error) {
	features := []pricing_model.PriceFeature{}
	query := `SELECT featureID, planID, text, isIncluded, orderNo FROM ms_price_feature
		WHERE planID = ? ORDER BY orderNo ASC, featureID ASC`
	err := r.db.Select(&features, query, planID)
	return features, err
}

func (r *pricingRepositoryImpl) Create(plan pricing_model.PricePlan, features []pricing_model.PriceFeature) (int, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return 0, err
	}
	committed := false
	defer func() {
		if !committed {
			_ = tx.Rollback()
		}
	}()

	result, err := tx.Exec(`INSERT INTO ms_price_plan
		(title, subtitle, price, currency, billingPeriod, badge, originalPrice, discountPercent, isFeatured, flagActive, orderNo)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		plan.Title, plan.Subtitle, plan.Price, plan.Currency, plan.BillingPeriod, plan.Badge,
		plan.OriginalPrice, plan.DiscountPercent, plan.IsFeatured, plan.FlagActive, plan.OrderNo)
	if err != nil {
		return 0, err
	}
	id, _ := result.LastInsertId()
	planID := int(id)

	if err := insertFeatures(tx, planID, features); err != nil {
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}
	committed = true
	return planID, nil
}

func (r *pricingRepositoryImpl) Update(plan pricing_model.PricePlan, features []pricing_model.PriceFeature) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}
	committed := false
	defer func() {
		if !committed {
			_ = tx.Rollback()
		}
	}()

	if _, err := tx.Exec(`UPDATE ms_price_plan SET
		title=?, subtitle=?, price=?, currency=?, billingPeriod=?, badge=?,
		originalPrice=?, discountPercent=?, isFeatured=?, flagActive=?, orderNo=?
		WHERE planID=?`,
		plan.Title, plan.Subtitle, plan.Price, plan.Currency, plan.BillingPeriod, plan.Badge,
		plan.OriginalPrice, plan.DiscountPercent, plan.IsFeatured, plan.FlagActive, plan.OrderNo, plan.PlanID); err != nil {
		return err
	}

	if _, err := tx.Exec(`DELETE FROM ms_price_feature WHERE planID = ?`, plan.PlanID); err != nil {
		return err
	}

	if err := insertFeatures(tx, plan.PlanID, features); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	committed = true
	return nil
}

func (r *pricingRepositoryImpl) Delete(planID int) error {
	_, err := r.db.Exec(`DELETE FROM ms_price_plan WHERE planID = ?`, planID)
	return err
}

func insertFeatures(tx *sqlx.Tx, planID int, features []pricing_model.PriceFeature) error {
	for i, f := range features {
		if _, err := tx.Exec(`INSERT INTO ms_price_feature (planID, text, isIncluded, orderNo) VALUES (?, ?, ?, ?)`,
			planID, f.Text, f.IsIncluded, i); err != nil {
			return err
		}
	}
	return nil
}

func NullFloat(v *float64) null.Float {
	if v == nil {
		return null.Float{}
	}
	return null.FloatFrom(*v)
}

func NullInt(v *int) null.Int {
	if v == nil {
		return null.Int{}
	}
	return null.IntFrom(int64(*v))
}
