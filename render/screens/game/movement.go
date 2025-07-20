package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	MAX_SPEED_FORWARD = 0.7
	MAX_SPEED_REVERSE = -0.3
)

var carPosition = rl.NewVector3(-9.45, -1, 0)

var acceleration = float32(0.1)

func Movement() {
	EngineSim()
	if rl.IsKeyDown(rl.KeyW) {
		acceleration = MaxSpeed(acceleration + 0.02)
	}

	if rl.IsKeyDown(rl.KeyS) {
		acceleration = MaxSpeed(acceleration - 0.02)
	}

	if rl.IsKeyDown(rl.KeySpace) {
		if acceleration < 0.1 && acceleration > -0.1 {
			acceleration = 0
		}
		if acceleration < -0.1 { // REVERSE
			acceleration = MaxSpeed(acceleration + 0.03)
		}
		if acceleration > 0.1 { // FORWARD
			acceleration = MaxSpeed(acceleration - 0.03)
		}
	}
}

func EngineSim() {
	if acceleration == 0 {
		return
	}
	carPosition.Z += acceleration
}

func MaxSpeed(speed float32) float32 {
	if speed > MAX_SPEED_FORWARD {
		return MAX_SPEED_FORWARD
	}
	if speed < MAX_SPEED_REVERSE {
		return MAX_SPEED_REVERSE
	}
	return speed
}
