package camera

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/nuriofernandez/car-around-the-city/render/screens/game/vehicle"
)

func PreRender() {
	HandleCameraRotation()
	Camera.Target = *vehicle.GetCarPos()
}

func BeginMode3D() {
	rl.BeginMode3D(Camera)
}
