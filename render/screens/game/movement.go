package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

const (
	MAX_SPEED_FORWARD = 0.7
	MAX_SPEED_REVERSE = -0.3
)

var wheelsRotation = float32(0)

var carPosition = rl.NewVector3(-9.45, -1, 0)

var acceleration = float32(0.1)

var steering = float32(0)

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

	if rl.IsKeyDown(rl.KeyA) {
		steering = MaxSteering(steering - 2)
	}

	if rl.IsKeyDown(rl.KeyD) {
		steering = MaxSteering(steering + 2)
	}
}

func EngineSim() {
	wheelsRotation = float32(math.Mod(float64(wheelsRotation+(acceleration*40)), 360.0))
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

func MaxSteering(angle float32) float32 {
	if angle > 28 {
		return 28
	}
	if angle < -28 {
		return -28
	}
	return angle
}
