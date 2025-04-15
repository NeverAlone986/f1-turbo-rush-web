package game

import "fyne.io/fyne/v2"

func ChangeTrack(trackNum int) {
	state.mu.Lock()
	defer state.mu.Unlock()
	
	state.CurrentTrack = trackNum
	state.GameActive = false // Остановить текущую игру
}

func GetTrackConfig(trackNumber int) TrackConfig {
	config := TrackConfig{
		EnemyCount:    3,
		ObstacleCount: 5,
		EnemySpeed:    2.0,
		PlayerSpeed:   5.0,
		Background:    "assets/tracks/track1.png",
	}

	switch trackNumber {
	case 2:
		config = TrackConfig{
			EnemyCount:    4,
			ObstacleCount: 6,
			EnemySpeed:    2.5,
			PlayerSpeed:   5.0,
			Background:    "assets/tracks/track2.png",
		}
	case 3:
		config = TrackConfig{
			EnemyCount:    5,
			ObstacleCount: 7,
			EnemySpeed:    3.0,
			PlayerSpeed:   5.5,
			Background:    "assets/tracks/track3.png",
		}
	}

	return config
}
