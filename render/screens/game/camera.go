package game

import rl "github.com/gen2brain/raylib-go/raylib"

var camera = rl.Camera{
	Position:   rl.NewVector3(carPosition.X+5, carPosition.Y+10.0, carPosition.Z+5.0),
	Target:     carPosition,
	Up:         rl.NewVector3(0.0, 1.0, 0.0),
	Fovy:       45.0,
	Projection: rl.CameraPerspective,
}

func CameraController() {
	camera.Position = rl.NewVector3(carPosition.X+0, carPosition.Y+7.0, carPosition.Z-10)
	camera.Target = carPosition
}
