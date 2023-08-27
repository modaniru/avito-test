package services

import (
	"context"
	"fmt"

	"github.com/modaniru/avito/internal/entity"
	"github.com/modaniru/avito/internal/storage"
)

type HistoryService struct {
	historyStorage storage.History
}

func NewHistoryService(historyStorage storage.History) *HistoryService {
	return &HistoryService{historyStorage: historyStorage}
}

func (h *HistoryService) GetHistory(ctx context.Context) ([]entity.History, error) {
	op := "internal.service.services.HistoryService.GetHistory"

	history, err := h.historyStorage.GetHistory(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s get history error: %w", op, err)
	}
	return history, nil
}

func (h *HistoryService) GetHistoryByDate(ctx context.Context, date string) ([]entity.History, error) {
	op := "internal.service.services.HistoryService.GetHistoryByDate"

	history, err := h.historyStorage.GetHistoryByDate(ctx, date)
	if err != nil {
		return nil, fmt.Errorf("%s get history by date error: %w", op, err)
	}
	return history, nil
}
