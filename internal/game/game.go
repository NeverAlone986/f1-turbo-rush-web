package game

import (
	"os"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

const (
	recordFile = "record.txt"
)

var (
	state    GameState
	objects  GameObjects
)

func InitGame() {
	state = GameState{
		PlayerX:     175,
		PlayerY:     500,
		KeysPressed: make(map[fyne.KeyName]bool),
		CurrentTrack: 1,
		GameActive:   true,
		WindowSize:  fyne.NewSize(400, 600),
	}
	loadRecord()
}

func loadRecord() {
	if data, err := os.ReadFile(recordFile); err == nil {
		if r, err := strconv.Atoi(string(data)); err == nil {
			state.SetScore(0)
			state.Record = r
		}
	}
}

func saveRecord() {
	if state.Score > state.Record {
		state.Record = state.Score
		os.WriteFile(recordFile, []byte(strconv.Itoa(state.Record)), 0644)
	}
}

func SetupGame(window fyne.Window) {
	config := GetTrackConfig(state.CurrentTrack)
	
	objects.Background = canvas.NewImageFromFile(config.Background)
	objects.Background.Resize(state.WindowSize)
	
	objects.PlayerCar = canvas.NewImageFromFile("assets/player/player.png")
	objects.PlayerCar.Resize(fyne.NewSize(50, 100))
	objects.PlayerCar.Move(fyne.NewPos(state.PlayerX, state.PlayerY))

	initEnemies(config.EnemyCount)
	initObstacles(config.ObstacleCount)

	objects.ScoreLabel = widget.NewLabelWithStyle("Score: 0", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	objects.RecordLabel = widget.NewLabelWithStyle("Record: "+strconv.Itoa(state.Record), fyne.TextAlignTrailing, fyne.TextStyle{Bold: true})

	gameContent := container.NewWithoutLayout(objects.Background, objects.PlayerCar, objects.ScoreLabel, objects.RecordLabel)
	for _, car := range objects.EnemyCars {
		gameContent.Add(car)
	}
	for _, obstacle := range objects.Obstacles {
		gameContent.Add(obstacle)
	}

	window.SetContent(gameContent)
	addKeyboardControl(window)
	go gameLoop(window, config)
}

// ... остальные функции (isOverlapping, getValidPosition и т.д.) остаются аналогичными, 
// но используют state и objects вместо глобальных переменных ...
