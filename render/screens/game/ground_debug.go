package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/nuriofernandez/car-around-the-city/render/screens/game/vehicle"
)

var (
	driverFront    rl.Vector3
	driverBack     rl.Vector3
	passengerFront rl.Vector3
	passengerBack  rl.Vector3
)

func CalculateGroundCube() {
	driverFront = HitGround(vehicle.PedVehicle.DriverFrontWheel.Position.Location)
	driverBack = HitGround(vehicle.PedVehicle.DriverBackWheel.Position.Location)
	passengerFront = HitGround(vehicle.PedVehicle.PassengerFrontWheel.Position.Location)
	passengerBack = HitGround(vehicle.PedVehicle.PassengerBackWheel.Position.Location)
}

func RenderGroundCube() {
	rl.DrawCube(driverFront, 0.2, 0.2, 0.2, rl.Yellow)
	rl.DrawCube(driverBack, 0.2, 0.2, 0.2, rl.Yellow)
	rl.DrawCube(passengerFront, 0.2, 0.2, 0.2, rl.Yellow)
	rl.DrawCube(passengerBack, 0.2, 0.2, 0.2, rl.Yellow)
}

func HitGround(location rl.Vector3) rl.Vector3 {
	hit := Hit(location)
	return rl.Vector3{
		X: location.X,
		Y: hit,
		Z: location.Z,
	}
}

func Hit(location rl.Vector3) float32 {
	var bestHit rl.RayCollision
	for _, mesh := range terrainModel.GetMeshes() {
		ray := rl.NewRay(rl.NewVector3(float32(location.X), location.Y+0.65, float32(location.Z)), rl.NewVector3(0, -1, 0))
		hit := rl.GetRayCollisionMesh(ray, mesh, terrainModel.Transform)
		if hit.Hit {

			if hit.Distance <= bestHit.Distance || bestHit.Distance == 0 {
				bestHit = hit
			}
		}
	}

	if !bestHit.Hit {
		return 0
	}

	return bestHit.Point.Y
}
