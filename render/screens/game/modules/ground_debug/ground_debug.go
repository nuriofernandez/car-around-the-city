package ground_debug

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/nuriofernandez/car-around-the-city/render/screens/game/modules/terrain"
	"github.com/nuriofernandez/car-around-the-city/render/screens/game/vehicle"
)

var (
	driverFront    rl.Vector3
	driverBack     rl.Vector3
	passengerFront rl.Vector3
	passengerBack  rl.Vector3
)

func PreRenderGroundCube() {
	driverFront = terrain.HitGround(vehicle.PedVehicle.DriverFrontWheel.Position.Location)
	driverBack = terrain.HitGround(vehicle.PedVehicle.DriverBackWheel.Position.Location)
	passengerFront = terrain.HitGround(vehicle.PedVehicle.PassengerFrontWheel.Position.Location)
	passengerBack = terrain.HitGround(vehicle.PedVehicle.PassengerBackWheel.Position.Location)
}

func RenderGroundCube() {
	rl.DrawCube(rl.Vector3Add(driverFront, rl.Vector3{Y: 0.1}), 0.2, 0.2, 0.2, rl.Yellow)
	rl.DrawCube(rl.Vector3Add(driverBack, rl.Vector3{Y: 0.1}), 0.2, 0.2, 0.2, rl.Yellow)
	rl.DrawCube(rl.Vector3Add(passengerFront, rl.Vector3{Y: 0.1}), 0.2, 0.2, 0.2, rl.Yellow)
	rl.DrawCube(rl.Vector3Add(passengerBack, rl.Vector3{Y: 0.1}), 0.2, 0.2, 0.2, rl.Yellow)
}
