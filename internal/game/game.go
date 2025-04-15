package game

import (
	"math/rand"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

const (
	objectSize  = 50
	updateDelay = 10 * time.Millisecond
)

// InitGame инициализирует состояние игры
func InitGame() {
	state = &GameState{
		PlayerX:     175,
		PlayerY:     500,
		KeysPressed: make(map[fyne.KeyName]bool),
		CurrentTrack: 1,
		GameActive:   true,
		WindowSize:  fyne.NewSize(400, 600),
	}
	objects = &GameObjects{}
	//loadRecord()
}

func getValidPosition(existing []*canvas.Image) (float32, float32) {
	for {
		x := float32(rand.Intn(int(state.WindowSize.Width) - objectSize))
		y := float32(-rand.Intn(300) - 100)
		if !isOverlapping(x, y, existing) {
			return x, y
		}
	}
}

func isOverlapping(x, y float32, objs []*canvas.Image) bool {
	for _, obj := range objs {
		pos := obj.Position()
		if x < pos.X+objectSize && x+objectSize > pos.X &&
			y < pos.Y+objectSize && y+objectSize > pos.Y {
			return true
		}
	}
	return false
}

func initEnemies(count int) {
	objects.EnemyCars = make([]*canvas.Image, count)
	for i := 0; i < count; i++ {
		x, y := getValidPosition(nil)
		enemy := canvas.NewImageFromFile("assets/enemies/enemy" + strconv.Itoa(i+1) + ".png")
		enemy.Resize(fyne.NewSize(objectSize, 100))
		enemy.Move(fyne.NewPos(x, y))
		objects.EnemyCars[i] = enemy
	}
}

func initObstacles(count int) {
	objects.Obstacles = make([]*canvas.Image, count)
	for i := 0; i < count; i++ {
		x, y := getValidPosition(objects.EnemyCars)
		obstacle := canvas.NewImageFromFile("assets/obstacles/obstacle" + strconv.Itoa(i+1) + ".png")
		obstacle.Resize(fyne.NewSize(objectSize, objectSize))
		obstacle.Move(fyne.NewPos(x, y))
		objects.Obstacles[i] = obstacle
	}
}

func addKeyboardControl(window fyne.Window) {
	if deskCanvas, ok := window.Canvas().(desktop.Canvas); ok {
		deskCanvas.SetOnKeyDown(func(e *fyne.KeyEvent) {
			state.KeysPressed[e.Name] = true
		})
		deskCanvas.SetOnKeyUp(func(e *fyne.KeyEvent) {
			delete(state.KeysPressed, e.Name)
		})
	}
}

func gameLoop(window fyne.Window, config TrackConfig) {
	ticker := time.NewTicker(updateDelay)
	defer ticker.Stop()

	for range ticker.C {
		if !state.GameActive {
			return
		}

		state.SetScore(state.Score + 1)
		objects.ScoreLabel.SetText("Score: " + strconv.Itoa(state.Score))

		for _, enemy := range objects.EnemyCars {
			pos := enemy.Position()
			enemy.Move(fyne.NewPos(pos.X, pos.Y+config.EnemySpeed))
		}

		window.Content().Refresh()
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

	gameContent := container.NewWithoutLayout(
		objects.Background,
		objects.PlayerCar,
		objects.ScoreLabel,
		objects.RecordLabel,
	)

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

// GetState возвращает текущее состояние игры (для API)
func GetState() GameState {
	return state.GetState()
}

// ChangeTrack меняет текущую трассу (для API)
func ChangeTrack(trackNum int) {
	state.mu.Lock()
	defer state.mu.Unlock()
	state.CurrentTrack = trackNum
	state.GameActive = false // Остановить текущую игру
}
