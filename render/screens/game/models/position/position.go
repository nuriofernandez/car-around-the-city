package position

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

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

func Distance(a, b rl.Vector3) float32 {
	dx := b.X - a.X
	dy := b.Y - a.Y
	dz := b.Z - a.Z
	return float32(math.Sqrt(float64(dx*dx + dy*dy + dz*dz)))
}
