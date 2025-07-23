package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/nuriofernandez/car-around-the-city/render/screens/game/driver"
)

func Movement() {
	if rl.IsKeyDown(rl.KeyW) {
		driver.Accelerate(0.005)
	}

	if rl.IsKeyDown(rl.KeyS) {
		driver.Accelerate(-0.005)
	}

	if rl.IsKeyDown(rl.KeySpace) {
		driver.Break()
	}

	if rl.IsKeyDown(rl.KeyA) {
		driver.Steer(-2)
	}

	if rl.IsKeyDown(rl.KeyD) {
		driver.Steer(2)
	}
}
