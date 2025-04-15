package game

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"sync"
)

// GameState содержит все состояние игры
type GameState struct {
	mu           sync.Mutex        // Для безопасного доступа из разных горутин
	Score        int               // Текущий счет
	Record       int               // Рекордный счет
	PlayerX      float32           // X-позиция игрока
	PlayerY      float32           // Y-позиция игрока
	CurrentTrack int               // Номер текущей трассы (1, 2, 3...)
	GameActive   bool              // Флаг активности игры
	KeysPressed  map[fyne.KeyName]bool // Нажатые клавиши
	WindowSize   fyne.Size         // Размер окна
}

// GameObjects содержит все графические объекты игры
type GameObjects struct {
	PlayerCar   *canvas.Image      // Машина игрока
	EnemyCars   []*canvas.Image    // Вражеские машины
	Obstacles   []*canvas.Image    // Препятствия
	Background  *canvas.Image      // Фон (трасса)
	ScoreLabel  *widget.Label      // Текст счета
	RecordLabel *widget.Label      // Текст рекорда
	GameOverBox *fyne.Container    // Контейнер для экрана "Game Over"
}

// TrackConfig конфигурация трассы
type TrackConfig struct {
	EnemyCount   int     // Количество врагов
	ObstacleCount int    // Количество препятствий
	EnemySpeed   float32 // Скорость врагов
	PlayerSpeed  float32 // Скорость игрока
	Background   string  // Путь к изображению трассы
}

// GetState возвращает текущее состояние игры (потокобезопасно)
func (gs *GameState) GetState() GameState {
	gs.mu.Lock()
	defer gs.mu.Unlock()
	return *gs
}

// SetScore обновляет счет (потокобезопасно)
func (gs *GameState) SetScore(newScore int) {
	gs.mu.Lock()
	defer gs.mu.Unlock()
	gs.Score = newScore
}

// GetTrackConfig возвращает конфигурацию для текущей трассы
func GetTrackConfig(trackNumber int) TrackConfig {
	// Базовые настройки
	config := TrackConfig{
		EnemyCount:   3,
		ObstacleCount: 5,
		EnemySpeed:   2.0,
		PlayerSpeed:  5.0,
		Background:   "assets/tracks/track1.png",
	}

	// Настройки для разных трасс
	switch trackNumber {
	case 1:
		// Стандартные настройки (уже заданы)
	case 2:
		config.EnemyCount = 4
		config.ObstacleCount = 6
		config.EnemySpeed = 2.5
		config.Background = "assets/tracks/track2.png"
	case 3:
		config.EnemyCount = 5
		config.ObstacleCount = 7
		config.EnemySpeed = 3.0
		config.Background = "assets/tracks/track3.png"
	}

	return config
}
