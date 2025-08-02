package game

import (
	"github.com/nuriofernandez/car-around-the-city/render/screens/game/modules/terrain"
	"github.com/nuriofernandez/car-around-the-city/render/screens/game/vehicle"
)

func Falling() {
	carPosition := vehicle.PedVehicle.Body.Position.Location

	var bestHitY = terrain.HitGround(carPosition).Y

	minPos := bestHitY + 0.333785 //  car space

	if minPos < vehicle.PedVehicle.Body.Position.Location.Y {
		vehicle.PedVehicle.Body.Position.Location.Y -= 0.18
		vehicle.PedVehicle.IsGrounded = false
	}

	if vehicle.PedVehicle.Body.Position.Location.Y < minPos {
		vehicle.PedVehicle.IsGrounded = true
		vehicle.PedVehicle.Body.Position.Location.Y = minPos
	}
}
