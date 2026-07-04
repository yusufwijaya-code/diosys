package summary_service

import (
	"database/sql"
	"errors"

	"portfolio-api/base/helpers/error_helper"
	"portfolio-api/modules/summary/summary_dto"
	"portfolio-api/modules/summary/summary_model"
	"portfolio-api/modules/summary/summary_repository"
)

// SummaryService exposes summary business operations scoped to a developer.
type SummaryService interface {
	GetByUser(userID int) (summary_dto.SummaryResponse, error)
	Save(userID int, request summary_dto.SummaryRequest) (summary_dto.SummaryResponse, error)
}

type summaryServiceImpl struct {
	repository summary_repository.SummaryRepository
}

// NewSummaryService builds a SummaryService.
func NewSummaryService(repository summary_repository.SummaryRepository) SummaryService {
	return &summaryServiceImpl{repository: repository}
}

func (s *summaryServiceImpl) GetByUser(userID int) (summary_dto.SummaryResponse, error) {
	summary, err := s.repository.FindByUser(userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return summary_dto.SummaryResponse{
				UserID: userID,
				Stats:  []summary_model.SummaryStat{},
				Facts:  []summary_model.SummaryFact{},
			}, nil
		}
		return summary_dto.SummaryResponse{}, error_helper.Internal(err)
	}

	stats, err := s.repository.GetStats(summary.SummaryID)
	if err != nil {
		return summary_dto.SummaryResponse{}, error_helper.Internal(err)
	}
	facts, err := s.repository.GetFacts(summary.SummaryID)
	if err != nil {
		return summary_dto.SummaryResponse{}, error_helper.Internal(err)
	}

	return summary_dto.SummaryResponse{
		SummaryID: summary.SummaryID,
		UserID:    summary.UserID,
		Content:   summary.Content,
		Stats:     stats,
		Facts:     facts,
	}, nil
}

func (s *summaryServiceImpl) Save(userID int, request summary_dto.SummaryRequest) (summary_dto.SummaryResponse, error) {
	stats := make([]summary_model.SummaryStat, 0, len(request.Stats))
	for _, stat := range request.Stats {
		stats = append(stats, summary_model.SummaryStat{Number: stat.Number, Label: stat.Label})
	}
	facts := make([]summary_model.SummaryFact, 0, len(request.Facts))
	for _, fact := range request.Facts {
		facts = append(facts, summary_model.SummaryFact{Icon: fact.Icon, Text: fact.Text})
	}

	if _, err := s.repository.Save(userID, request.Content, stats, facts); err != nil {
		return summary_dto.SummaryResponse{}, error_helper.Internal(err)
	}
	return s.GetByUser(userID)
}
