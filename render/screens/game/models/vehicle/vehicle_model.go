package vehicle

import rl "github.com/gen2brain/raylib-go/raylib"

var hardcodedVehicleModel = VehicleModel{
	BodyFile:                  "resources/car.glb",
	WheelFile:                 "resources/wheel.glb",
	PassengerFrontWheelOffset: rl.NewVector3(-0.961470, 0.042085, 1.61003),
	PassengerBackWheelOffset:  rl.NewVector3(-0.963425, 0.042085, -1.73766),
	DriverFrontWheelOffset:    rl.NewVector3(0.954046, 0.042085, 1.60227),
	DriverBackWheelOffset:     rl.NewVector3(0.963129, 0.042085, -1.73766),
}

type VehicleModel struct {
	BodyFile                  string
	WheelFile                 string
	PassengerFrontWheelOffset rl.Vector3
	PassengerBackWheelOffset  rl.Vector3
	DriverFrontWheelOffset    rl.Vector3
	DriverBackWheelOffset     rl.Vector3
}
