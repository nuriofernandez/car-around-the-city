package vehicle

import "math"

func EngineSim() {
	wheelsRotation = float32(math.Mod(float64(wheelsRotation+(acceleration*40)), 360.0))
	carPosition.Z += acceleration
}
