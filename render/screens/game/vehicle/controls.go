package vehicle

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	MAX_SPEED_FORWARD = 0.7
	MAX_SPEED_REVERSE = -0.3
)

var carPosition = rl.NewVector3(-9.45, -1, 0)
var acceleration = float32(0)
var steering = float32(0)

func GetCarPos() rl.Vector3 {
	return carPosition
}

func Accelerate(speed float32) {
	acceleration = MaxSpeed(acceleration + speed)
}

func Break() {
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

func Steer(direction float32) {
	steering = MaxSteering(steering + direction)
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
