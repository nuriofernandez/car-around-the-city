package game

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"time"
)

var rotation float32 = 1
var start time.Time = time.Now()

func Render() {
	Movement()

	CameraController()

	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	rl.BeginMode3D(camera)

	rl.DrawModel(terrainModel, rl.NewVector3(0, 0, 0), 1.0, rl.White)

	// Car
	rl.DrawModel(carModel, carPosition, 1.0, rl.White)

	// Car wheels
	rl.DrawModelEx(driverWheelModel, rl.NewVector3(carPosition.X+0.954046, carPosition.Y+0.042085, carPosition.Z+1.60227), rl.Vector3{1, 0, 0}, wheelsRotation, rl.Vector3{1, 1, 1}, rl.White) // driver front
	rl.DrawModelEx(driverWheelModel, rl.NewVector3(carPosition.X+0.963129, carPosition.Y+0.042085, carPosition.Z-1.73766), rl.Vector3{1, 0, 0}, wheelsRotation, rl.Vector3{1, 1, 1}, rl.White) // driver back

	rl.DrawModelEx(passengerWheelModel, rl.NewVector3(carPosition.X-0.96147, carPosition.Y+0.042085, carPosition.Z+1.61003), rl.Vector3{1, 0, 0}, wheelsRotation, rl.Vector3{1, 1, 1}, rl.White)  // passenger front
	rl.DrawModelEx(passengerWheelModel, rl.NewVector3(carPosition.X-0.963425, carPosition.Y+0.042085, carPosition.Z-1.73766), rl.Vector3{1, 0, 0}, wheelsRotation, rl.Vector3{1, 1, 1}, rl.White) // passenger back

	rl.EndMode3D()

	cords := fmt.Sprintf("[%d,%d,%d] R:%.2f / S:%.2f", int(carPosition.X), int(carPosition.Y), int(carPosition.Z), wheelsRotation, acceleration)
	rl.DrawText(cords, 10, 10, 20, rl.Black)

	rl.EndDrawing()
}
