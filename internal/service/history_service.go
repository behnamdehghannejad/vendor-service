package service

import (
	"github.com/behnamdehghannejad/vendorservice/internal/domain"
	"github.com/behnamdehghannejad/vendorservice/internal/port"
)

type HistoryService struct {
	repository port.HistoryRepository
}

func NewHistoryService(repository port.HistoryRepository) *HistoryService {
	return &HistoryService{repository: repository}
}

func (s *HistoryService) Update(history domain.History) error {
	return s.repository.Update(history)
}

func (s *HistoryService) Search(search domain.SearchHistory) ([]domain.History, error) {
	return s.repository.Filter(search)
}
