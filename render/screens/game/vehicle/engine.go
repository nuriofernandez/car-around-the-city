package vehicle

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/nuriofernandez/car-around-the-city/render/screens/game/driver"
	"math"
)

func EngineSim() {
	wheelsRotation = float32(math.Mod(float64(wheelsRotation+(driver.Acceleration*40)), 360.0))

	yawRadians := float64((180 - carYaw) * rl.Deg2rad)

	// Compute direction vector based on yaw
	dx := float32(math.Cos(yawRadians))
	dz := float32(math.Sin(yawRadians))

	// Apply acceleration in facing direction
	carPosition.Z -= dx * driver.Acceleration
	carPosition.X -= dz * driver.Acceleration

	reduceAcceleration(0.002)
}

func reduceAcceleration(reduction float32) {
	if driver.Acceleration > 0 {
		driver.Acceleration = float32(math.Max(float64(driver.Acceleration-reduction), 0))
	} else {
		driver.Acceleration = float32(math.Min(float64(driver.Acceleration+reduction), 0))
	}
}
