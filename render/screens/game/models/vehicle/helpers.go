package vehicle

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/nuriofernandez/car-around-the-city/render/screens/game/models/entity"
)

func transform(entity *entity.Entity) {
	rotation := rl.Vector3{
		X: rl.Deg2rad * entity.Position.Rotation.Pitch,
		Y: rl.Deg2rad * entity.Position.Rotation.Yaw,
		Z: rl.Deg2rad * entity.Position.Rotation.Roll,
	}
	entity.Object.Transform = rl.MatrixRotateXYZ(rotation)
}
