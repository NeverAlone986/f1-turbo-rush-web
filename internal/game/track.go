package game

func GetTrackConfig(trackNumber int) TrackConfig {
	baseConfig := TrackConfig{
		EnemyCount:    3,
		ObstacleCount: 5,
		EnemySpeed:    2.0,
		PlayerSpeed:   5.0,
		Background:    "assets/tracks/track1.png",
	}

	switch trackNumber {
	case 2:
		return TrackConfig{
			EnemyCount:    4,
			ObstacleCount: 6,
			EnemySpeed:    2.5,
			PlayerSpeed:   5.0,
			Background:    "assets/tracks/track2.png",
		}
	case 3:
		return TrackConfig{
			EnemyCount:    5,
			ObstacleCount: 7,
			EnemySpeed:    3.0,
			PlayerSpeed:   5.5,
			Background:    "assets/tracks/track3.png",
		}
	default:
		return baseConfig
	}
}
