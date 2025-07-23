package position

import rl "github.com/gen2brain/raylib-go/raylib"

type Rotation struct {
	Pitch float32
	Roll  float32
	Yaw   float32
}
type Position struct {
	Location rl.Vector3
	Rotation Rotation
}

func NewPosition(x, y, z, pitch, roll, yaw float32) Position {
	return Position{
		Location: rl.Vector3{X: x, Y: y, Z: z},
		Rotation: Rotation{
			Pitch: pitch,
			Roll:  roll,
			Yaw:   yaw,
		},
	}
}

func NewPositionVector(vector rl.Vector3) Position {
	return Position{
		Location: vector,
	}
}
