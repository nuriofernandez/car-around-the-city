package vehicle

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

func EngineSim() {
	wheelsRotation = float32(math.Mod(float64(wheelsRotation+(acceleration*40)), 360.0))

	yawRadians := float64((180 - carYaw) * rl.Deg2rad)

	// Compute direction vector based on yaw
	dx := float32(math.Cos(yawRadians))
	dz := float32(math.Sin(yawRadians))

	// Apply acceleration in facing direction
	carPosition.Z -= dx * acceleration
	carPosition.X -= dz * acceleration
}
