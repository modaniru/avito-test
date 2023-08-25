package repos

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSaveSegment(t *testing.T) {
	storage := NewSegmentStorage(db)
	insertSegments(storage, t)
	defer clearDB()
}

func TestGetSegmentsWhenEmpty(t *testing.T) {
	storage := NewSegmentStorage(db)
	res, err := storage.GetSegments(context.Background())
	require.NoError(t, err)
	require.Empty(t, res)
}

func TestGetSegments(t *testing.T) {
	storage := NewSegmentStorage(db)
	insertSegments(storage, t)
	defer clearDB()

	res, err := storage.GetSegments(context.Background())
	require.NoError(t, err)
	require.Len(t, res, 4)
}

// TODO удалить несуществующий сегмент
func TestDeleteSegment(t *testing.T) {
	storage := NewSegmentStorage(db)
	insertSegments(storage, t)
	defer clearDB()
	name := "AVITO_VOICE_MESSAGES"

	err := storage.DeleteSegment(context.Background(), name)
	require.NoError(t, err)

	res, err := storage.GetSegments(context.Background())
	require.NoError(t, err)

	for _, i := range res {
		require.NotEqual(t, name, i.Name)
	}
}

func insertSegments(storage *SegmentStorage, t *testing.T) {
	segments := []string{"AVITO_VOICE_MESSAGES", "AVITO_PERFORMANCE_VAS", "AVITO_DISCOUNT_30", "AVITO_DISCOUNT_50"}
	for _, segment := range segments {
		id, err := storage.SaveSegment(context.Background(), segment)
		require.NoError(t, err)
		require.NotZero(t, id)
	}
}
