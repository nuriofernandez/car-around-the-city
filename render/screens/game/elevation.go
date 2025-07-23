package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/nuriofernandez/car-around-the-city/render/screens/game/vehicle"
)

func Elevation() {
	carPosition := vehicle.GetCarPos()

	currentPoint := float32(-1)
	for _, mesh := range terrainModel.GetMeshes() {
		ray := rl.NewRay(rl.NewVector3(float32(carPosition.X), carPosition.Y+2, float32(carPosition.Z)), rl.NewVector3(0, -1, 0))
		hit := rl.GetRayCollisionMesh(ray, mesh, terrainModel.Transform)
		if hit.Hit {
			hitPoint := hit.Point.Y

			if hitPoint > currentPoint {
				currentPoint = hitPoint
			}
			currentPoint = hitPoint
		}
	}

	carPosition.Y = currentPoint + 0.333785 //  car space
}
