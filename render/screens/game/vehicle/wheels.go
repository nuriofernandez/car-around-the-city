package vehicle

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var wheelsRotation = float32(0)

func CarWheels() {
	driverBack()
	driverFront()

	passengerBack()
	passengerFront()
}

func driverBack() {
	// driver back
	var (
		pitch = float32(-wheelsRotation)
		roll  = float32(0)
		yaw   = float32(carYaw)
	)

	rotation := rl.Vector3{
		X: rl.Deg2rad * pitch,
		Y: rl.Deg2rad * yaw,
		Z: rl.Deg2rad * roll,
	}
	driverBackWheelModel.Transform = rl.MatrixRotateXYZ(rotation)
}

func driverFront() {
	// driver back
	var (
		pitch = float32(-wheelsRotation)
		roll  = float32(0)
		yaw   = float32(carYaw + steering) // Max 28
	)

	rotation := rl.Vector3{
		X: rl.Deg2rad * pitch,
		Y: rl.Deg2rad * yaw,
		Z: rl.Deg2rad * roll,
	}
	driverFrontWheelModel.Transform = rl.MatrixRotateXYZ(rotation)
}

func passengerBack() {
	// driver back
	var (
		pitch = float32(-wheelsRotation)
		roll  = float32(-180)
		yaw   = float32(-carYaw + 180)
	)

	rotation := rl.Vector3{
		X: rl.Deg2rad * pitch,
		Y: rl.Deg2rad * yaw,
		Z: rl.Deg2rad * roll,
	}
	passengerBackWheelModel.Transform = rl.MatrixRotateXYZ(rotation)
}

func passengerFront() {
	// driver back
	var (
		pitch = float32(-wheelsRotation)
		roll  = float32(-180)
		yaw   = float32(-carYaw+180) + -steering // Max -28
	)

	rotation := rl.Vector3{
		X: rl.Deg2rad * pitch,
		Y: rl.Deg2rad * yaw,
		Z: rl.Deg2rad * roll,
	}
	passengerFrontWheelModel.Transform = rl.MatrixRotateXYZ(rotation)
}
