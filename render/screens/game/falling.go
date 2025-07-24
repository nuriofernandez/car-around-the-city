package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/nuriofernandez/car-around-the-city/render/screens/game/vehicle"
)

func Falling() {
	carPosition := vehicle.PedVehicle.Body.Position.Location

	currentPoint := float32(-99)
	for _, mesh := range terrainModel.GetMeshes() {
		ray := rl.NewRay(rl.NewVector3(float32(carPosition.X), carPosition.Y+0.65, float32(carPosition.Z)), rl.NewVector3(0, -1, 0))
		hit := rl.GetRayCollisionMesh(ray, mesh, terrainModel.Transform)
		if hit.Hit {
			hitPoint := hit.Point.Y

			if hit.Distance <= 1.2 {
				currentPoint = hitPoint
			}
		}
	}

	minPos := currentPoint + 0.333785 //  car space

	if minPos <= vehicle.PedVehicle.Body.Position.Location.Y {
		vehicle.PedVehicle.Body.Position.Location.Y -= 0.18
	}

	if vehicle.PedVehicle.Body.Position.Location.Y < minPos {
		vehicle.PedVehicle.Body.Position.Location.Y = minPos
	}
}
