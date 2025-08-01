package terrain

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

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
	for _, mesh := range TerrainModel.GetMeshes() {
		ray := rl.NewRay(rl.NewVector3(location.X, location.Y+0.65, location.Z), rl.NewVector3(0, -1, 0))
		hit := rl.GetRayCollisionMesh(ray, mesh, TerrainModel.Transform)
		if hit.Hit {

			if hit.Distance <= bestHit.Distance || bestHit.Distance == 0 {
				bestHit = hit
			}
		}
	}

	if !bestHit.Hit {
		return -99
	}

	return bestHit.Point.Y
}
