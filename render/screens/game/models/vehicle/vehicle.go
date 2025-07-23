package vehicle

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/nuriofernandez/car-around-the-city/render/screens/game/models/entity"
	"github.com/nuriofernandez/car-around-the-city/render/screens/game/models/position"
	"math"
)

type Vehicle struct {
	Model               VehicleModel
	Body                entity.Entity
	passengerFrontWheel entity.Entity
	passengerBackWheel  entity.Entity
	driverFrontWheel    entity.Entity
	driverBackWheel     entity.Entity
	Steering            float32
	Acceleration        float32
	wheelRotation       float32
}

func (v *Vehicle) Update() {
	// Update wheels position
	v.passengerFrontWheel.Position = position.NewPositionVector(rl.Vector3Add(v.Body.Position.Location, rotateOffsetY(hardcodedVehicleModel.PassengerFrontWheelOffset, v.Body.Position.Rotation.Yaw)))
	v.passengerBackWheel.Position = position.NewPositionVector(rl.Vector3Add(v.Body.Position.Location, rotateOffsetY(hardcodedVehicleModel.PassengerBackWheelOffset, v.Body.Position.Rotation.Yaw)))
	v.driverFrontWheel.Position = position.NewPositionVector(rl.Vector3Add(v.Body.Position.Location, rotateOffsetY(hardcodedVehicleModel.DriverFrontWheelOffset, v.Body.Position.Rotation.Yaw)))
	v.driverBackWheel.Position = position.NewPositionVector(rl.Vector3Add(v.Body.Position.Location, rotateOffsetY(hardcodedVehicleModel.DriverBackWheelOffset, v.Body.Position.Rotation.Yaw)))

	// Update back wheels rotation
	v.driverBackWheel.Position.Rotation.Yaw = v.Body.Position.Rotation.Yaw
	v.passengerBackWheel.Position.Rotation.Yaw = v.Body.Position.Rotation.Yaw - 180

	// Update front wheels rotation
	v.driverFrontWheel.Position.Rotation.Yaw = v.Body.Position.Rotation.Yaw + v.Steering
	v.passengerFrontWheel.Position.Rotation.Yaw = v.Body.Position.Rotation.Yaw - 180 + v.Steering

	// Update wheels speed (pitch)
	v.wheelRotation = float32(math.Mod(float64(v.wheelRotation+(v.Acceleration*40)), 360.0))
	v.driverBackWheel.Position.Rotation.Pitch = -v.wheelRotation
	v.passengerBackWheel.Position.Rotation.Pitch = v.wheelRotation
	v.driverFrontWheel.Position.Rotation.Pitch = -v.wheelRotation
	v.passengerFrontWheel.Position.Rotation.Pitch = v.wheelRotation

	// Apply transformations
	transform(&v.Body)
	transform(&v.driverBackWheel)
	transform(&v.passengerBackWheel)
	transform(&v.driverFrontWheel)
	transform(&v.passengerFrontWheel)
}

func (v *Vehicle) Render() {
	v.Update()

	rl.DrawModel(v.Body.Object, v.Body.Position.Location, 1.0, rl.White)
	rl.DrawModel(v.passengerFrontWheel.Object, v.passengerFrontWheel.Position.Location, 1.0, rl.White)
	rl.DrawModel(v.passengerBackWheel.Object, v.passengerBackWheel.Position.Location, 1.0, rl.White)
	rl.DrawModel(v.driverFrontWheel.Object, v.driverFrontWheel.Position.Location, 1.0, rl.White)
	rl.DrawModel(v.driverBackWheel.Object, v.driverBackWheel.Position.Location, 1.0, rl.White)
}
