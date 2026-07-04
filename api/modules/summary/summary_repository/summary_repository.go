package summary_repository

import (
	"database/sql"
	"errors"

	"portfolio-api/modules/summary/summary_model"

	"github.com/jmoiron/sqlx"
)

// SummaryRepository handles persistence for a developer's summary aggregate.
type SummaryRepository interface {
	FindByUser(userID int) (summary_model.Summary, error)
	GetStats(summaryID int) ([]summary_model.SummaryStat, error)
	GetFacts(summaryID int) ([]summary_model.SummaryFact, error)
	Save(userID int, content string, stats []summary_model.SummaryStat, facts []summary_model.SummaryFact) (int, error)
}

type summaryRepositoryImpl struct {
	db *sqlx.DB
}

// NewSummaryRepository builds a SummaryRepository.
func NewSummaryRepository(db *sqlx.DB) SummaryRepository {
	return &summaryRepositoryImpl{db: db}
}

func (r *summaryRepositoryImpl) FindByUser(userID int) (summary_model.Summary, error) {
	var summary summary_model.Summary
	query := `SELECT summaryID, userID, content, createdDate, editedDate
		FROM ms_summary WHERE userID = ? LIMIT 1`
	err := r.db.Get(&summary, query, userID)
	return summary, err
}

func (r *summaryRepositoryImpl) GetStats(summaryID int) ([]summary_model.SummaryStat, error) {
	stats := []summary_model.SummaryStat{}
	query := `SELECT summaryStatID, summaryID, number, label, orderNo FROM ms_summary_stat
		WHERE summaryID = ? ORDER BY orderNo ASC, summaryStatID ASC`
	err := r.db.Select(&stats, query, summaryID)
	return stats, err
}

func (r *summaryRepositoryImpl) GetFacts(summaryID int) ([]summary_model.SummaryFact, error) {
	facts := []summary_model.SummaryFact{}
	query := `SELECT summaryFactID, summaryID, icon, text, orderNo FROM ms_summary_fact
		WHERE summaryID = ? ORDER BY orderNo ASC, summaryFactID ASC`
	err := r.db.Select(&facts, query, summaryID)
	return facts, err
}

func (r *summaryRepositoryImpl) Save(userID int, content string, stats []summary_model.SummaryStat, facts []summary_model.SummaryFact) (int, error) {
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

	var summaryID int
	err = tx.Get(&summaryID, `SELECT summaryID FROM ms_summary WHERE userID = ? LIMIT 1`, userID)
	if errors.Is(err, sql.ErrNoRows) {
		result, insertErr := tx.Exec(`INSERT INTO ms_summary (userID, content) VALUES (?, ?)`, userID, content)
		if insertErr != nil {
			return 0, insertErr
		}
		id, _ := result.LastInsertId()
		summaryID = int(id)
	} else if err != nil {
		return 0, err
	} else {
		if _, updateErr := tx.Exec(`UPDATE ms_summary SET content = ? WHERE summaryID = ?`, content, summaryID); updateErr != nil {
			return 0, updateErr
		}
	}

	for _, table := range []string{"ms_summary_stat", "ms_summary_fact"} {
		if _, delErr := tx.Exec("DELETE FROM "+table+" WHERE summaryID = ?", summaryID); delErr != nil {
			return 0, delErr
		}
	}

	for i, stat := range stats {
		if _, insErr := tx.Exec(`INSERT INTO ms_summary_stat (summaryID, number, label, orderNo) VALUES (?, ?, ?, ?)`,
			summaryID, stat.Number, stat.Label, i); insErr != nil {
			return 0, insErr
		}
	}
	for i, fact := range facts {
		if _, insErr := tx.Exec(`INSERT INTO ms_summary_fact (summaryID, icon, text, orderNo) VALUES (?, ?, ?, ?)`,
			summaryID, fact.Icon, fact.Text, i); insErr != nil {
			return 0, insErr
		}
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}
	committed = true
	return summaryID, nil
}
