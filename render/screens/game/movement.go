package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/nuriofernandez/car-around-the-city/render/screens/game/chat"
	"github.com/nuriofernandez/car-around-the-city/render/screens/game/driver"
	"github.com/nuriofernandez/car-around-the-city/render/screens/game/vehicle"
)

func Movement() {
	// If chat is open, ignore inputs
	if chat.ChatOpen {
		return
	}

	if rl.IsKeyDown(rl.KeyW) {
		if vehicle.PedVehicle.IsGrounded {
			driver.Accelerate(0.005)
		}
	}

	if rl.IsKeyDown(rl.KeyS) {
		if vehicle.PedVehicle.IsGrounded {
			driver.Accelerate(-0.005)
		}
	}

	if rl.IsKeyDown(rl.KeySpace) {
		if vehicle.PedVehicle.IsGrounded {
			driver.Break()
		}
	}

	if rl.IsKeyDown(rl.KeyA) {
		driver.Steer(-2)
	}

	if rl.IsKeyDown(rl.KeyD) {
		driver.Steer(2)
	}

	// TMP reset altitude
	if rl.IsKeyPressed(rl.KeyR) {
		chat.Add("Updated vehicle position! (Y = 10)")
		vehicle.PedVehicle.Body.Position.Location.Y = 10
	}
}
