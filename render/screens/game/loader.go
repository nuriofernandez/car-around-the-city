package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/nuriofernandez/car-around-the-city/render/screens/game/audio"
	"github.com/nuriofernandez/car-around-the-city/render/screens/game/vehicle"
)

func Load() {
	LoadMap()
	vehicle.Load()
	rl.DisableCursor()
	audio.Load()
}

func UnLoad() {
	rl.EnableCursor()
}
