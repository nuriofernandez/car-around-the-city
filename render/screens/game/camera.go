package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/nuriofernandez/car-around-the-city/render/screens/game/vehicle"
	"math"
)

var camera = rl.Camera{
	Position:   rl.NewVector3(vehicle.GetCarPos().X+5, vehicle.GetCarPos().Y+10.0, vehicle.GetCarPos().Z+5.0),
	Target:     *vehicle.GetCarPos(),
	Up:         rl.NewVector3(0.0, 1.0, 0.0),
	Fovy:       45.0,
	Projection: rl.CameraPerspective,
}

var (
	// Camera angle around the car
	cameraAngleX float32 = 0.0  // horizontal angle (yaw)
	cameraAngleY float32 = 20.0 // vertical angle (pitch)

	// Distance from the car
	cameraDistance float32 = 10.0
)

func CameraController() {
	HandleCameraRotation()
	camera.Target = *vehicle.GetCarPos()
}

func HandleCameraRotation() {
	// Mouse control
	deltaX := rl.GetMouseDelta().X
	deltaY := rl.GetMouseDelta().Y

	cameraAngleX -= deltaX * 0.3
	cameraAngleY += deltaY * 0.3

	// Clamp vertical angle to avoid flipping
	if cameraAngleY > 89.0 {
		cameraAngleY = 89.0
	}
	if cameraAngleY < -89.0 {
		cameraAngleY = -89.0
	}

	// Convert spherical coordinates to Cartesian
	radius := cameraDistance
	yaw := rl.Deg2rad * cameraAngleX
	pitch := rl.Deg2rad * cameraAngleY

	camera.Position.X = vehicle.GetCarPos().X + radius*float32(math.Cos(float64(pitch)))*float32(math.Sin(float64(yaw)))
	camera.Position.Y = vehicle.GetCarPos().Y + radius*float32(math.Sin(float64(pitch)))
	camera.Position.Z = vehicle.GetCarPos().Z + radius*float32(math.Cos(float64(pitch)))*float32(math.Cos(float64(yaw)))
}
