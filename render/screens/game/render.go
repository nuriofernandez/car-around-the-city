package game

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/nuriofernandez/car-around-the-city/render/screens/game/audio"
	"github.com/nuriofernandez/car-around-the-city/render/screens/game/chat"
	"github.com/nuriofernandez/car-around-the-city/render/screens/game/vehicle"
	"github.com/nuriofernandez/car-around-the-city/settings"
)

func Render() {
	Movement()
	Falling()
	vehicle.PreRender()
	CalculateGroundCube()
	audio.Loop()
	CalculatePitch()

	CameraController()

	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	rl.BeginMode3D(camera)

	rl.DrawModel(terrainModel, rl.NewVector3(0, 0, 0), 1.0, rl.White)

	vehicle.Render()
	RenderGroundCube()

	rl.EndMode3D()

	carPosition := vehicle.GetCarPos()
	cords := fmt.Sprintf("[%d,%d,%d]", int(carPosition.X), int(carPosition.Y), int(carPosition.Z))
	rl.DrawText(cords, settings.ScreenWidth-10-rl.MeasureText(cords, 20), 10, 20, rl.Black)

	chat.Render()

	rl.EndDrawing()
}
