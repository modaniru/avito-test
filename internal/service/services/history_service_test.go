package services_test

import (
	"context"
	"errors"
	"testing"

	"github.com/modaniru/avito/internal/entity"
	"github.com/modaniru/avito/internal/service/services"
	"github.com/modaniru/avito/internal/storage/mocks"
	diskMock "github.com/modaniru/avito/internal/yandex_drive/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestGetHistoryByDate(t *testing.T) {
	disk := diskMock.NewDisk(t)
	storage := mocks.NewHistory(t)
	historyService := services.NewHistoryService(storage, disk)
	date := "2023-08-28"

	history := []entity.History{{UserId: 1, SegmentName: "test", Operation: "oper", OperationTime: "2023-08-28"}}

	storage.On("GetHistoryByDate", mock.Anything, "2023-08-28").Once().Return(history, nil)
	disk.On("CreateFile", "report_"+date, mock.Anything).Once().Return("link", nil)
	disk.On("IsAvailible").Once().Return(true)

	link, err := historyService.GetHistoryByDate(context.Background(), date)
	require.NoError(t, err)
	require.Equal(t, "link", link)
}

func TestGetHistoryByDateWithStorageError(t *testing.T) {
	disk := diskMock.NewDisk(t)
	storage := mocks.NewHistory(t)
	historyService := services.NewHistoryService(storage, disk)
	date := "2023-08-28"

	history := []entity.History{{UserId: 1, SegmentName: "test", Operation: "oper", OperationTime: "2023-08-28"}}

	disk.On("IsAvailible").Once().Return(true)
	storage.On("GetHistoryByDate", mock.Anything, "2023-08-28").Once().Return(history, errors.New("error"))

	link, err := historyService.GetHistoryByDate(context.Background(), date)
	require.Error(t, err)
	require.Equal(t, "", link)

}

func TestGetHistoryByDateWithYandexError(t *testing.T) {
	disk := diskMock.NewDisk(t)
	storage := mocks.NewHistory(t)
	historyService := services.NewHistoryService(storage, disk)
	date := "2023-08-28"

	history := []entity.History{{UserId: 1, SegmentName: "test", Operation: "oper", OperationTime: "2023-08-28"}}

	storage.On("GetHistoryByDate", mock.Anything, "2023-08-28").Once().Return(history, nil)
	disk.On("CreateFile", "report_"+date, mock.Anything).Once().Return("", errors.New("error"))
	disk.On("IsAvailible").Once().Return(true)

	link, err := historyService.GetHistoryByDate(context.Background(), date)
	require.Error(t, err)
	require.Equal(t, "", link)
}

func TestGetHistoryByDateWithUnavailibleService(t *testing.T) {
	disk := diskMock.NewDisk(t)
	storage := mocks.NewHistory(t)
	historyService := services.NewHistoryService(storage, disk)
	date := "2023-08-28"

	disk.On("IsAvailible").Once().Return(false)

	link, err := historyService.GetHistoryByDate(context.Background(), date)
	require.Error(t, err)
	require.Equal(t, "", link)
}
