package history

import (
	"time"
)

type PlaybackEntry struct {
	UserID      int
	ContentID   int
	ContentType string
	Timestamp   time.Time
}

var playbackHistory []PlaybackEntry

func AddPlayback(userID, contentID int, contentType string) {
	playbackHistory = append(playbackHistory, PlaybackEntry{
		UserID:      userID,
		ContentID:   contentID,
		ContentType: contentType,
		Timestamp:   time.Now(),
	})
}

func GetHistory(userID int) []PlaybackEntry {
	var userHistory []PlaybackEntry
	for i := len(playbackHistory) - 1; i >= 0; i-- {
		if playbackHistory[i].UserID == userID {
			userHistory = append(userHistory, playbackHistory[i])
			if len(userHistory) >= 10 {
				break
			}
		}
	}
	return userHistory
}