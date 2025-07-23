package entity

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/nuriofernandez/car-around-the-city/render/screens/game/models/position"
)

type Entity struct {
	Object   rl.Model
	Position position.Position
}
