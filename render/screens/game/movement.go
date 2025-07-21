package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/nuriofernandez/car-around-the-city/render/screens/game/vehicle"
)

func Movement() {
	if rl.IsKeyDown(rl.KeyW) {
		vehicle.Accelerate(0.02)
	}

	if rl.IsKeyDown(rl.KeyS) {
		vehicle.Accelerate(-0.02)
	}

	if rl.IsKeyDown(rl.KeySpace) {
		vehicle.Break()
	}

	if rl.IsKeyDown(rl.KeyA) {
		vehicle.Steer(-2)
	}

	if rl.IsKeyDown(rl.KeyD) {
		vehicle.Steer(2)
	}
}
