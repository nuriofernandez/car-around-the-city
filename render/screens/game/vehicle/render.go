package vehicle

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

func PreRender() {
	HandleDirection()
	EngineSim()
	CarWheels()
}

func Render() {
	// Car
	rl.DrawModel(carModel, carPosition, 1.0, rl.White)

	driverBackOffset := rl.NewVector3(0.963129, 0.042085, -1.73766)
	passengerBackOffset := rl.NewVector3(-0.963425, 0.042085, -1.73766)
	driverFrontOffset := rl.NewVector3(0.954046, 0.042085, 1.60227)
	passengerFrontOffset := rl.NewVector3(-0.961470, 0.042085, 1.61003)

	// Car back wheels
	rl.DrawModel(driverBackWheelModel, rl.Vector3Add(carPosition, rotateOffsetY(driverBackOffset, carYaw)), 1, rl.White)
	rl.DrawModel(passengerBackWheelModel, rl.Vector3Add(carPosition, rotateOffsetY(passengerBackOffset, carYaw)), 1, rl.White)
	rl.DrawModel(driverFrontWheelModel, rl.Vector3Add(carPosition, rotateOffsetY(driverFrontOffset, carYaw)), 1, rl.White)
	rl.DrawModel(passengerFrontWheelModel, rl.Vector3Add(carPosition, rotateOffsetY(passengerFrontOffset, carYaw)), 1, rl.White)
}

func rotateOffsetY(offset rl.Vector3, yaw float32) rl.Vector3 {
	cosYaw := float32(math.Cos(float64(rl.Deg2rad * yaw)))
	sinYaw := float32(math.Sin(float64(rl.Deg2rad * yaw)))

	rotatedX := offset.X*cosYaw - offset.Z*sinYaw
	rotatedZ := offset.X*sinYaw + offset.Z*cosYaw

	return rl.NewVector3(rotatedX, offset.Y, rotatedZ)
}
