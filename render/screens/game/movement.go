package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/nuriofernandez/car-around-the-city/render/screens/game/driver"
	"github.com/nuriofernandez/car-around-the-city/render/screens/game/vehicle"
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

	// TMP reset altitude
	if rl.IsKeyDown(rl.KeyR) {
		vehicle.PedVehicle.Body.Position.Location.Y = 10
	}
}
