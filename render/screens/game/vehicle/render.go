package vehicle

import rl "github.com/gen2brain/raylib-go/raylib"

func Render() {
	// Car
	rl.DrawModel(carModel, carPosition, 1.0, rl.White)

	// Car back wheels
	rl.DrawModel(driverBackWheelModel, rl.NewVector3(carPosition.X+0.963129, carPosition.Y+0.042085, carPosition.Z-1.73766), 1, rl.White)    // driver back
	rl.DrawModel(passengerBackWheelModel, rl.NewVector3(carPosition.X-0.963425, carPosition.Y+0.042085, carPosition.Z-1.73766), 1, rl.White) // passenger back
	rl.DrawModel(driverFrontWheelModel, rl.NewVector3(carPosition.X+0.954046, carPosition.Y+0.042085, carPosition.Z+1.60227), 1, rl.White)   // driver front
	rl.DrawModel(passengerFrontWheelModel, rl.NewVector3(carPosition.X-0.96147, carPosition.Y+0.042085, carPosition.Z+1.61003), 1, rl.White) // passenger front
}
