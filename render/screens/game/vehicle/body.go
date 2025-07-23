package vehicle

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/nuriofernandez/car-around-the-city/render/screens/game/driver"
	"math"
)

var (
	pitch  = float32(0)
	roll   = float32(0)
	carYaw = float32(0)
)

func HandleDirection() {
	// Calculate next vehicle location properties
	calculateBodyRotation()

	// Update vehicle location
	rotation := rl.Vector3{
		X: rl.Deg2rad * pitch,
		Y: rl.Deg2rad * carYaw,
		Z: rl.Deg2rad * roll,
	}
	carModel.Transform = rl.MatrixRotateXYZ(rotation)
}

func calculateBodyRotation() {
	if driver.Acceleration == 0 || driver.Steering == 0 {
		return
	}

	// forward
	if driver.Acceleration > 0 {
		if driver.Steering > 0 {
			carYaw += (driver.Steering / 28) * yawMultiplierByAcceleration()
		}
		if driver.Steering < 0 {
			carYaw -= (-driver.Steering / 28) * yawMultiplierByAcceleration()
		}
		driver.ReduceSteering(1)
	}

	// backwards
	if driver.Acceleration < 0 {
		if driver.Steering > 0 {
			carYaw -= (driver.Steering / 28) * yawMultiplierByAcceleration()
		}
		if driver.Steering < 0 {
			carYaw += (-driver.Steering / 28) * yawMultiplierByAcceleration()
		}
		driver.ReduceSteering(1)
	}

	reduceAcceleration(0.001)
}

func smoothStep(x, edge0, edge1 float64) float64 {
	t := (x - edge0) / (edge1 - edge0)
	t = math.Max(0, math.Min(1, t)) // Clamp t to [0, 1]
	return t * t * (3 - 2*t)
}

func yawMultiplierByAcceleration() float32 {
	// forward
	if driver.Acceleration > 0 {
		value1 := smoothStep(float64(driver.Acceleration), 0, 0.15)
		value2 := smoothStep(float64(driver.Acceleration), 0.16, 0.5)

		// Blend values smoothly
		return float32((1-value1)*0.2 + value1*((1-value2)*1.5+value2*2.0))
	}

	//backwards
	if driver.Acceleration < 0 {
		value1 := smoothStep(float64(driver.Acceleration), 0, -0.15)
		value2 := smoothStep(float64(driver.Acceleration), -0.16, -0.5)

		// Blend values smoothly
		return float32((1-value1)*0.2 + value1*((1-value2)*1.5+value2*2.0))
	}

	// otherwise, driver.Acceleration is 0
	return 0
}
