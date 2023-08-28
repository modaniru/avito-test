package services_test

import (
	"context"
	"errors"
	"testing"

	"github.com/modaniru/avito/internal/entity"
	"github.com/modaniru/avito/internal/service/services"
	"github.com/modaniru/avito/internal/storage/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestSaveSegment(t *testing.T) {
	segmentStorage := mocks.NewSegment(t)
	segmentService := services.NewSegmentService(segmentStorage)

	segmentName := "segment_name"
	returningId := 1
	segmentStorage.On("SaveSegment", mock.Anything, segmentName).Once().Return(returningId, nil)

	id, err := segmentService.SaveSegment(context.Background(), segmentName)
	require.NoError(t, err)
	require.Equal(t, returningId, id)
}

func TestSaveSegmentWithError(t *testing.T) {
	segmentStorage := mocks.NewSegment(t)
	segmentService := services.NewSegmentService(segmentStorage)

	segmentName := "segment_name"
	segmentStorage.On("SaveSegment", mock.Anything, segmentName).Once().Return(0, errors.New("test"))

	_, err := segmentService.SaveSegment(context.Background(), segmentName)
	require.Error(t, err)
}

func TestGetSegments(t *testing.T) {
	segmentStorage := mocks.NewSegment(t)
	segmentService := services.NewSegmentService(segmentStorage)

	excepted := []entity.Segment{{Id: 1, Name: "segment_1"}, {Id: 2, Name: "segment_2"}}
	segmentStorage.On("GetSegments", mock.Anything).Once().Return(excepted, nil)

	list, err := segmentService.GetSegments(context.Background())
	require.NoError(t, err)
	require.Equal(t, excepted, list)
}

func TestGetSegmentsWithError(t *testing.T) {
	segmentStorage := mocks.NewSegment(t)
	segmentService := services.NewSegmentService(segmentStorage)

	segmentStorage.On("GetSegments", mock.Anything).Once().Return(nil, errors.New("error"))

	list, err := segmentService.GetSegments(context.Background())
	require.Error(t, err)
	require.Nil(t, list)
}

func TestDeleteSegment(t *testing.T) {
	segmentStorage := mocks.NewSegment(t)
	segmentService := services.NewSegmentService(segmentStorage)

	segmentName := "segment_name"
	segmentStorage.On("DeleteSegment", mock.Anything, segmentName).Once().Return(nil)

	err := segmentService.DeleteSegment(context.Background(), segmentName)
	require.NoError(t, err)
}

func TestDeleteSegmentWithError(t *testing.T) {
	segmentStorage := mocks.NewSegment(t)
	segmentService := services.NewSegmentService(segmentStorage)

	segmentName := "segment_name"
	segmentStorage.On("DeleteSegment", mock.Anything, segmentName).Once().Return(errors.New("test"))

	err := segmentService.DeleteSegment(context.Background(), segmentName)
	require.Error(t, err)
}
