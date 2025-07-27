package vehicle

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/nuriofernandez/car-around-the-city/render/screens/game/models/position"
	"math"
)

func (v *Vehicle) transformWheel(offset rl.Vector3) position.Position {
	return position.NewPositionVector(transformWheel(v.Body.Position, offset))
}

// Multiply 3x3 matrix by a vector
func multiplyMatrixVector(m [3][3]float32, v rl.Vector3) rl.Vector3 {
	return rl.Vector3{
		X: m[0][0]*v.X + m[0][1]*v.Y + m[0][2]*v.Z,
		Y: m[1][0]*v.X + m[1][1]*v.Y + m[1][2]*v.Z,
		Z: m[2][0]*v.X + m[2][1]*v.Y + m[2][2]*v.Z,
	}
}

func getRotationMatrix(pitch, roll, yaw float32) [3][3]float32 {
	p := pitch * math.Pi / 180 // around X (right)
	r := roll * math.Pi / 180  // around Z (forward)
	y := yaw * math.Pi / 180   // around Y (up)

	// Pitch: rotate around X
	Rx := [3][3]float32{
		{1, 0, 0},
		{0, float32(math.Cos(float64(p))), float32(math.Sin(float64(p)))},
		{0, float32(-math.Sin(float64(p))), float32(math.Cos(float64(p)))},
	}

	// Yaw: rotate around Y
	Ry := [3][3]float32{
		{float32(math.Cos(float64(y))), 0, float32(-math.Sin(float64(y)))},
		{0, 1, 0},
		{float32(math.Sin(float64(y))), 0, float32(math.Cos(float64(y)))},
	}

	// Roll: rotate around Z
	Rz := [3][3]float32{
		{float32(math.Cos(float64(r))), float32(math.Sin(float64(r))), 0},
		{float32(-math.Sin(float64(r))), float32(math.Cos(float64(r))), 0},
		{0, 0, 1},
	}

	// Final rotation = Yaw * Pitch * Roll (Ry * Rx * Rz)
	return multiplyMatrices(Ry, multiplyMatrices(Rx, Rz))
}

func multiplyMatrices(a, b [3][3]float32) [3][3]float32 {
	var result [3][3]float32
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			result[i][j] = 0
			for k := 0; k < 3; k++ {
				result[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return result
}

// Apply the body transform to a wheel offset
func transformWheel(body position.Position, offset rl.Vector3) rl.Vector3 {
	rot := getRotationMatrix(body.Rotation.Pitch, body.Rotation.Roll, body.Rotation.Yaw)
	rotated := multiplyMatrixVector(rot, offset)
	return rl.Vector3{
		X: rotated.X + body.Location.X,
		Y: rotated.Y + body.Location.Y,
		Z: rotated.Z + body.Location.Z,
	}
}
