package game

import rl "github.com/gen2brain/raylib-go/raylib"

var carPosition = rl.NewVector3(0, 0, 0)

var carEngine = true
var speed = float32(0.1)

func Movement() {
	EngineSim()
	if rl.IsKeyDown(rl.KeyW) {
		carEngine = true
		if speed <= 1 && speed >= 0 {
			speed += 0.1
		}
	}

	if rl.IsKeyDown(rl.KeyS) {
		speed -= 0.1
		if speed <= 0.1 {
			carEngine = false
			speed = 0
		}
	}
}

func EngineSim() {
	if carEngine {
		carPosition.Z += speed
	}
}
