package game

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"math/rand"
)

func ChangeTrack(trackNum int) {
	state.mu.Lock()
	defer state.mu.Unlock()
	
	state.CurrentTrack = trackNum
	// Здесь нужно сбросить позиции машин и препятствий для новой трассы
}

func getTrackImage() *canvas.Image {
	trackPath := "assets/tracks/track" + strconv.Itoa(state.CurrentTrack) + ".png"
	return canvas.NewImageFromFile(trackPath)
}

// Функции для работы с трассами
