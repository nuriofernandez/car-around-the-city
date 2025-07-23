package driver

import "math"

const (
	MAX_STEERING = 28
	MIN_STEERING = -28
)

var Steering = float32(0)

func Steer(direction float32) {
	Steering = maxSteering(Steering + direction)
}

func ReduceSteering(reduction float32) {
	if Steering > 0 {
		Steering = float32(math.Max(float64(Steering-reduction), 0))
	} else {
		Steering = float32(math.Min(float64(Steering+reduction), 0))
	}
}

func maxSteering(angle float32) float32 {
	if angle > MAX_STEERING {
		return MAX_STEERING
	}
	if angle < MIN_STEERING {
		return MIN_STEERING
	}
	return angle
}
