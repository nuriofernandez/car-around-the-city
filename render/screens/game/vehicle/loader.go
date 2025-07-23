package vehicle

import (
	"github.com/nuriofernandez/car-around-the-city/render/screens/game/models/vehicle"
)

var PedVehicle vehicle.Vehicle

func Load() {
	PedVehicle = vehicle.NewVehicle(-9.45, -1, 2)
}
