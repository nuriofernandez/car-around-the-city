package game

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/nuriofernandez/car-around-the-city/render/screens/game/vehicle"
)

func Render() {
	Movement()
	vehicle.PreRender()

	CameraController()

	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	rl.BeginMode3D(camera)

	rl.DrawModel(terrainModel, rl.NewVector3(0, 0, 0), 1.0, rl.White)

	vehicle.Render()

	rl.EndMode3D()

	carPosition := vehicle.GetCarPos()
	cords := fmt.Sprintf("[%d,%d,%d]", int(carPosition.X), int(carPosition.Y), int(carPosition.Z))
	rl.DrawText(cords, 10, 10, 20, rl.Black)

	rl.EndDrawing()
}
