package game

import (
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

const (
	objectSize = 50
)

type GameState struct {
	mu          sync.Mutex
	Score       int
	Record      int
	PlayerX     float32
	PlayerY     float32
	CurrentTrack int
	KeysPressed map[fyne.KeyName]bool
}

var (
	state      GameState
	playerCar  *canvas.Image
	enemyCars  []*canvas.Image
	obstacles  []*canvas.Image
	scoreLabel *widget.Label
	recordLabel *widget.Label
)

func InitGame() {
	state = GameState{
		PlayerX:     175,
		PlayerY:     500,
		KeysPressed: make(map[fyne.KeyName]bool),
		CurrentTrack: 1,
	}
	loadRecord()
}

func loadRecord() {
	if data, err := os.ReadFile("record.txt"); err == nil {
		if r, err := strconv.Atoi(string(data)); err == nil {
			state.Record = r
		}
	}
}

func saveRecord() {
	state.mu.Lock()
	defer state.mu.Unlock()
	
	if state.Score > state.Record {
		state.Record = state.Score
		os.WriteFile("record.txt", []byte(strconv.Itoa(state.Record)), 0644)
	}
}

// Остальные функции игры аналогичны оригинальному коду, но адаптированы для веба
// и используют state.mu для синхронизации доступа к состоянию
