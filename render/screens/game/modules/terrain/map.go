package terrain

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var TerrainModel rl.Model

func LoadMap() {
	TerrainModel = rl.LoadModel("resources/city-model.glb")
}
