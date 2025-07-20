package game

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Render() {
	Movement()

	CameraController()

	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	rl.BeginMode3D(camera)

	rl.DrawModel(terrainModel, rl.NewVector3(0, 0, 0), 1.0, rl.White)
	rl.DrawModel(carModel, carPosition, 1.0, rl.White)

	rl.EndMode3D()

	cords := fmt.Sprintf("[%d,%d,%d]", int(camera.Position.X), int(camera.Position.Y), int(camera.Position.Z))
	rl.DrawText(cords, 10, 10, 20, rl.DarkGray)

	rl.EndDrawing()
}
