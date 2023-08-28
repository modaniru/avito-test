package services

import (
	"bytes"
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"strconv"

	"github.com/modaniru/avito/internal/entity"
	"github.com/modaniru/avito/internal/storage"
	yandexdrive "github.com/modaniru/avito/internal/yandex_drive"
)

type HistoryService struct {
	historyStorage storage.History
	yandexDisk     yandexdrive.Disk
}

var (
	ErrServiceUnavailible = errors.New("history service unavailible")
)

func NewHistoryService(historyStorage storage.History, yandexDisk yandexdrive.Disk) *HistoryService {
	return &HistoryService{historyStorage: historyStorage, yandexDisk: yandexDisk}
}

func (h *HistoryService) GetHistoryByDate(ctx context.Context, date string) (string, error) {
	op := "internal.service.services.HistoryService.GetHistoryByDate"

	if !h.yandexDisk.IsAvailible() {
		return "", ErrServiceUnavailible
	}
	history, err := h.historyStorage.GetHistoryByDate(ctx, date)
	if err != nil {
		return "", fmt.Errorf("%s get history by date error: %w", op, err)
	}
	b, err := h.encodeHistoryToCSV(history)
	if err != nil {
		return "", fmt.Errorf("%s encode history error: %w", op, err)
	}
	link, err := h.yandexDisk.CreateFile("report_"+date, b)
	if err != nil {
		return "", fmt.Errorf("%s create file error: %w", op, err)
	}
	return link, nil
}

func (h *HistoryService) encodeHistoryToCSV(history []entity.History) ([]byte, error) {
	b := bytes.Buffer{}
	w := csv.NewWriter(&b)

	columns := []string{"ID", "user_id", "segment_name", "operation", "operation_time"}
	err := w.Write(columns)
	if err != nil {
		return nil, err
	}

	for i, h := range history {
		row := []string{strconv.Itoa(i + 1), strconv.Itoa(h.UserId), h.SegmentName, h.Operation, h.OperationTime}
		err := w.Write(row)
		if err != nil {
			return nil, err
		}
	}
	w.Flush()

	return b.Bytes(), nil
}
