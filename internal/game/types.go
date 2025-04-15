package game

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"sync"
)

var (
	state   *GameState
	objects *GameObjects
)

type GameState struct {
	mu           sync.Mutex
	Score        int
	Record       int
	PlayerX      float32
	PlayerY      float32
	CurrentTrack int
	GameActive   bool
	KeysPressed  map[fyne.KeyName]bool
	WindowSize   fyne.Size
}

type GameObjects struct {
	PlayerCar   *canvas.Image
	EnemyCars   []*canvas.Image
	Obstacles   []*canvas.Image
	Background  *canvas.Image
	ScoreLabel  *widget.Label
	RecordLabel *widget.Label
}

type TrackConfig struct {
	EnemyCount    int
	ObstacleCount int
	EnemySpeed    float32
	PlayerSpeed   float32
	Background    string
}

func (gs *GameState) GetState() GameState {
	gs.mu.Lock()
	defer gs.mu.Unlock()
	return *gs
}

func (gs *GameState) SetScore(newScore int) {
	gs.mu.Lock()
	defer gs.mu.Unlock()
	gs.Score = newScore
}
