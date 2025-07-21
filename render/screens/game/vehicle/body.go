package vehicle

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	pitch  = float32(0)
	roll   = float32(0)
	carYaw = float32(0)
)

func HandleDirection() {
	rotation := rl.Vector3{
		X: rl.Deg2rad * pitch,
		Y: rl.Deg2rad * carYaw,
		Z: rl.Deg2rad * roll,
	}
	carModel.Transform = rl.MatrixRotateXYZ(rotation)

	steerDirection()
}

func steerDirection() {
	if acceleration == 0 || steering == 0 {
		return
	}

	// forward
	if acceleration > 0 {
		if steering > 0 {
			carYaw += (steering / 28) * 2
			steering--
		}
		if steering < 0 {
			carYaw -= (-steering / 28) * 2
			steering++
		}
	}

	// backwards
	if acceleration < 0 {
		if steering > 0 {
			carYaw -= (steering / 28) * 2
			steering--
		}
		if steering < 0 {
			carYaw += (-steering / 28) * 2
			steering++
		}
	}

	reduceAcceleration(0.003)
}
