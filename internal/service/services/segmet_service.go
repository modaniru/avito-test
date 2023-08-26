package services

import (
	"context"
	"fmt"

	"github.com/modaniru/avito/internal/entity"
	"github.com/modaniru/avito/internal/storage"
)

type SegmentService struct {
	segmentStorage storage.Segment
}

func NewSegmentService(segmentStorage storage.Segment) *SegmentService {
	return &SegmentService{segmentStorage: segmentStorage}
}

func (s *SegmentService) SaveSegment(ctx context.Context, name string) (int, error) {
	op := "internal.service.services.SegmentService.SaveSegment"

	id, err := s.segmentStorage.SaveSegment(ctx, name)
	if err != nil {
		return 0, fmt.Errorf("%s save segment error %w: ", op, err)
	}
	return id, err
}

func (s *SegmentService) GetSegments(ctx context.Context) ([]entity.Segment, error) {
	op := "internal.service.services.SegmentService.GetSegments"

	segments, err := s.segmentStorage.GetSegments(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s get segments error %w: ", op, err)
	}
	return segments, err
}

func (s *SegmentService) DeleteSegment(ctx context.Context, name string) error {
	op := "internal.service.services.SegmentService.DeleteSegment"

	err := s.segmentStorage.DeleteSegment(ctx, name)
	if err != nil {
		return fmt.Errorf("%s delete segment error %w: ", op, err)
	}
	return err
}
