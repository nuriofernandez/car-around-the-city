package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var terrainModel rl.Model
var carModel rl.Model
var driverWheelModel rl.Model
var passengerWheelModel rl.Model

func LoadMap() {
	terrainModel = rl.LoadModel("resources/city-model.glb")
	carModel = rl.LoadModel("resources/car.glb")
	driverWheelModel = rl.LoadModel("resources/wheel.glb")
	passengerWheelModel = rl.LoadModel("resources/passenger-wheel.glb")
}
