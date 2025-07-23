package vehicle

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/nuriofernandez/car-around-the-city/render/screens/game/driver"
	"math"
)

func EngineSim() {
	yawRadians := float64((180 - PedVehicle.Body.Position.Rotation.Yaw) * rl.Deg2rad)

	// Compute direction vector based on yaw
	dx := float32(math.Cos(yawRadians))
	dz := float32(math.Sin(yawRadians))

	// Apply acceleration in facing direction
	PedVehicle.Body.Position.Location.Z -= dx * PedVehicle.Acceleration
	PedVehicle.Body.Position.Location.X -= dz * PedVehicle.Acceleration

	reduceAcceleration(0.002)
}

func reduceAcceleration(reduction float32) {
	if driver.Acceleration > 0 {
		driver.Acceleration = float32(math.Max(float64(driver.Acceleration-reduction), 0))
	} else {
		driver.Acceleration = float32(math.Min(float64(driver.Acceleration+reduction), 0))
	}
}
