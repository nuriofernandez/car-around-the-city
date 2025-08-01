package terrain

import rl "github.com/gen2brain/raylib-go/raylib"

func Render() {
	rl.DrawModel(TerrainModel, rl.NewVector3(0, 0, 0), 1.0, rl.White)
}
