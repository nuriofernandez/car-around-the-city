package game

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/nuriofernandez/car-around-the-city/render/screens/game/audio"
	"github.com/nuriofernandez/car-around-the-city/render/screens/game/chat/chatrender"
	"github.com/nuriofernandez/car-around-the-city/render/screens/game/chat/commands/debug"
	"github.com/nuriofernandez/car-around-the-city/render/screens/game/modules/camera"
	"github.com/nuriofernandez/car-around-the-city/render/screens/game/modules/ground_debug"
	"github.com/nuriofernandez/car-around-the-city/render/screens/game/modules/terrain"
	"github.com/nuriofernandez/car-around-the-city/render/screens/game/vehicle"
	"github.com/nuriofernandez/car-around-the-city/settings"
)

func Render() {
	// Pre-render
	Movement()
	Falling()
	vehicle.PreRender()
	if debug.GroundDebugEnabled {
		ground_debug.PreRenderGroundCube()
	}
	audio.Loop()
	CalculatePitch()
	camera.PreRender()

	// 3D Render
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)
	camera.BeginMode3D()

	terrain.Render()
	vehicle.Render()
	if debug.GroundDebugEnabled {
		ground_debug.RenderGroundCube()
	}

	rl.EndMode3D()

	// 2D Render
	carPosition := vehicle.GetCarPos()
	cords := fmt.Sprintf("[%d,%d,%d]", int(carPosition.X), int(carPosition.Y), int(carPosition.Z))
	rl.DrawText(cords, settings.ScreenWidth-10-rl.MeasureText(cords, 20), 10, 20, rl.Black)

	chatrender.Render()

	rl.EndDrawing()
}
