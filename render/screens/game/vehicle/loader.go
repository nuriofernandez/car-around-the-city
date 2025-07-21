package vehicle

import rl "github.com/gen2brain/raylib-go/raylib"

var carModel rl.Model
var driverFrontWheelModel rl.Model
var driverBackWheelModel rl.Model
var passengerFrontWheelModel rl.Model
var passengerBackWheelModel rl.Model

func Load() {
	carModel = rl.LoadModel("resources/car.glb")
	driverFrontWheelModel = rl.LoadModel("resources/wheel.glb")
	driverBackWheelModel = rl.LoadModel("resources/wheel.glb")
	passengerFrontWheelModel = rl.LoadModel("resources/passenger-wheel.glb")
	passengerBackWheelModel = rl.LoadModel("resources/passenger-wheel.glb")
}
