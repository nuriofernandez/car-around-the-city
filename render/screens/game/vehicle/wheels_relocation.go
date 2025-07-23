package vehicle

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

var driverBackOffset = rl.NewVector3(0.963129, 0.042085, -1.73766)
var passengerBackOffset = rl.NewVector3(-0.963425, 0.042085, -1.73766)
var driverFrontOffset = rl.NewVector3(0.954046, 0.042085, 1.60227)
var passengerFrontOffset = rl.NewVector3(-0.961470, 0.042085, 1.61003)

func rotateOffsetY(offset rl.Vector3, yaw float32) rl.Vector3 {
	cosYaw := float32(math.Cos(float64(rl.Deg2rad * yaw)))
	sinYaw := float32(math.Sin(float64(rl.Deg2rad * yaw)))

	rotatedX := offset.X*cosYaw - offset.Z*sinYaw
	rotatedZ := offset.X*sinYaw + offset.Z*cosYaw

	return rl.NewVector3(rotatedX, offset.Y, rotatedZ)
}
