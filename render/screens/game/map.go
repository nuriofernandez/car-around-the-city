package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var terrainModel rl.Model
var carModel rl.Model

func LoadMap() {
	terrainModel = rl.LoadModel("resources/city-model.glb")
	carModel = rl.LoadModel("resources/car.glb")
}
