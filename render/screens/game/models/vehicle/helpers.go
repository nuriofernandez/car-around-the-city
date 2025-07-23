package vehicle

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/nuriofernandez/car-around-the-city/render/screens/game/models/entity"
	"math"
)

func rotateOffsetY(offset rl.Vector3, yaw float32) rl.Vector3 {
	cosYaw := float32(math.Cos(float64(rl.Deg2rad * yaw)))
	sinYaw := float32(math.Sin(float64(rl.Deg2rad * yaw)))

	rotatedX := offset.X*cosYaw - offset.Z*sinYaw
	rotatedZ := offset.X*sinYaw + offset.Z*cosYaw

	return rl.NewVector3(rotatedX, offset.Y, rotatedZ)
}

func transform(entity *entity.Entity) {
	rotation := rl.Vector3{
		X: rl.Deg2rad * entity.Position.Rotation.Pitch,
		Y: rl.Deg2rad * entity.Position.Rotation.Yaw,
		Z: rl.Deg2rad * entity.Position.Rotation.Roll,
	}
	entity.Object.Transform = rl.MatrixRotateXYZ(rotation)
}
