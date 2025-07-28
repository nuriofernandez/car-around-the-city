package game

import (
	"github.com/nuriofernandez/car-around-the-city/render/screens/game/vehicle"
	"math"
)

func CalculatePitch() {
	frontWheelPoint := Hit(vehicle.PedVehicle.DriverFrontWheel.Position.Location)
	backWheelPoint := Hit(vehicle.PedVehicle.DriverBackWheel.Position.Location)
	vehicleYaw := vehicle.PedVehicle.Body.Position.Rotation.Yaw

	// Calculate difference in world coordinates
	dx := float64(vehicle.PedVehicle.DriverFrontWheel.Position.Location.X - vehicle.PedVehicle.DriverBackWheel.Position.Location.X)
	dy := float64(frontWheelPoint - backWheelPoint)
	dz := float64(vehicle.PedVehicle.DriverFrontWheel.Position.Location.Z - vehicle.PedVehicle.DriverBackWheel.Position.Location.Z)

	forwardX := math.Cos(float64(-vehicleYaw))*dx - math.Sin(float64(-vehicleYaw))*dz
	forwardZ := math.Sin(float64(-vehicleYaw))*dx + math.Cos(float64(-vehicleYaw))*dz

	pitch := float32(math.Atan2(dy, math.Sqrt(forwardX*forwardX+forwardZ*forwardZ)) * 180 / math.Pi)
	vehicle.PedVehicle.Body.Position.Rotation.Pitch = pitch
}
