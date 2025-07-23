package vehicle

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func PreRender() {
	EngineSim()
	calculateCarWheelsRotation()
}

func Render() {
	// Car
	rl.DrawModel(carModel, carPosition, 1.0, rl.White)

	// Car back wheels
	rl.DrawModel(driverBackWheelModel, rl.Vector3Add(carPosition, rotateOffsetY(driverBackOffset, carYaw)), 1, rl.White)
	rl.DrawModel(passengerBackWheelModel, rl.Vector3Add(carPosition, rotateOffsetY(passengerBackOffset, carYaw)), 1, rl.White)
	rl.DrawModel(driverFrontWheelModel, rl.Vector3Add(carPosition, rotateOffsetY(driverFrontOffset, carYaw)), 1, rl.White)
	rl.DrawModel(passengerFrontWheelModel, rl.Vector3Add(carPosition, rotateOffsetY(passengerFrontOffset, carYaw)), 1, rl.White)

	HandleDirection()
}
