package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/nuriofernandez/car-around-the-city/render/screens/game/vehicle"
	"math"
)

func CalculatePitch() {
	frontWheelPoint := Point(vehicle.PedVehicle.DriverFrontWheel.Position.Location)
	backWheelPoint := Point(vehicle.PedVehicle.DriverBackWheel.Position.Location)
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

func Point(vector3 rl.Vector3) float32 {
	currentPoint := float32(-99)
	for _, mesh := range terrainModel.GetMeshes() {
		ray := rl.NewRay(rl.NewVector3(vector3.X, vector3.Y+2, vector3.Z), rl.NewVector3(0, -1, 0))
		hit := rl.GetRayCollisionMesh(ray, mesh, terrainModel.Transform)
		if hit.Hit {
			hitPoint := hit.Point.Y

			if hit.Distance <= 5 {
				currentPoint = hitPoint
			}
		}
	}
	return currentPoint
}
