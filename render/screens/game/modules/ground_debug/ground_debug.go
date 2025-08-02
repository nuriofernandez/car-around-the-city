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

var vehicleBoxModel rl.Model

func Load() {
	vehicleBoxModel = rl.LoadModelFromMesh(rl.GenMeshCube(2.5, 1.3, 6.3))
}

func PreRenderGroundCube() {
	driverFront = terrain.HitGround(vehicle.PedVehicle.DriverFrontWheel.Position.Location)
	driverBack = terrain.HitGround(vehicle.PedVehicle.DriverBackWheel.Position.Location)
	passengerFront = terrain.HitGround(vehicle.PedVehicle.PassengerFrontWheel.Position.Location)
	passengerBack = terrain.HitGround(vehicle.PedVehicle.PassengerBackWheel.Position.Location)
}

func RenderGroundCube() {
	RenderCarBox()
	rl.DrawCube(rl.Vector3Add(driverFront, rl.Vector3{Y: 0.1}), 0.2, 0.2, 0.2, rl.Yellow)
	rl.DrawCube(rl.Vector3Add(driverBack, rl.Vector3{Y: 0.1}), 0.2, 0.2, 0.2, rl.Yellow)
	rl.DrawCube(rl.Vector3Add(passengerFront, rl.Vector3{Y: 0.1}), 0.2, 0.2, 0.2, rl.Yellow)
	rl.DrawCube(rl.Vector3Add(passengerBack, rl.Vector3{Y: 0.1}), 0.2, 0.2, 0.2, rl.Yellow)
}

func RenderCarBox() {
	vehicleBoxModel.Transform = vehicle.PedVehicle.Body.Object.Transform
	rl.DrawModel(vehicleBoxModel,
		rl.Vector3Add(vehicle.PedVehicle.Body.Position.Location, rl.Vector3{Y: 0.55}),
		1,
		rl.NewColor(255, 255, 255, 30),
	)
}
