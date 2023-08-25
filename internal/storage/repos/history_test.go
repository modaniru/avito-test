package repos

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/modaniru/avito/internal/entity"
	"github.com/stretchr/testify/require"
)

func TestGetEmptyHistory(t *testing.T) {
	storage := NewHistoryStorage(db)
	list, err := storage.GetHistory(context.Background())
	require.NoError(t, err)
	require.Len(t, list, 0)
}

func TestGetHistory(t *testing.T) {
	historyStorage := NewHistoryStorage(db)
	userStorage := NewUserStorage(db)
	segmentStorage := NewSegmentStorage(db)
	defer clearDB()

	insertUsers(userStorage, t)
	insertSegments(segmentStorage, t)

	exceptedHistory := []entity.History{{UserId: 1, SegmentName: "AVITO_VOICE_MESSAGES", Operation: FollowOperation}}
	segments := []string{"AVITO_VOICE_MESSAGES"}
	err := userStorage.FollowToSegments(context.Background(), 1, segments)
	require.NoError(t, err)
	history, err := historyStorage.GetHistory(context.Background())
	require.Len(t, history, len(exceptedHistory))
	for i := range history {
		require.Equal(t, exceptedHistory[i].UserId, history[i].UserId)
		require.Equal(t, exceptedHistory[i].SegmentName, history[i].SegmentName)
		require.Equal(t, exceptedHistory[i].Operation, history[i].Operation)
	}
	time.Sleep(time.Second)
	err = userStorage.UnFollowToSegments(context.Background(), 1, []string{"AVITO_VOICE_MESSAGES"})
	exceptedHistory = append([]entity.History{{UserId: 1, SegmentName: "AVITO_VOICE_MESSAGES", Operation: UnFollowOperation}}, exceptedHistory...)
	history, err = historyStorage.GetHistory(context.Background())
	fmt.Println(history)
	require.Len(t, history, len(exceptedHistory))
	for i := range history {
		require.Equal(t, exceptedHistory[i].UserId, history[i].UserId)
		require.Equal(t, exceptedHistory[i].SegmentName, history[i].SegmentName)
		require.Equal(t, exceptedHistory[i].Operation, history[i].Operation)
	}
}
