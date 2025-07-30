package vehicle

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/nuriofernandez/car-around-the-city/render/screens/game/models/entity"
	"math"
)

type Vehicle struct {
	Model               VehicleModel
	Body                entity.Entity
	PassengerFrontWheel entity.Entity
	PassengerBackWheel  entity.Entity
	DriverFrontWheel    entity.Entity
	DriverBackWheel     entity.Entity
	Steering            float32
	Acceleration        float32
	wheelRotation       float32
}

func (v *Vehicle) Update() {
	// Update wheels position
	v.PassengerFrontWheel.Position = v.transformWheel(hardcodedVehicleModel.PassengerFrontWheelOffset)
	v.PassengerBackWheel.Position = v.transformWheel(hardcodedVehicleModel.PassengerBackWheelOffset)
	v.DriverFrontWheel.Position = v.transformWheel(hardcodedVehicleModel.DriverFrontWheelOffset)
	v.DriverBackWheel.Position = v.transformWheel(hardcodedVehicleModel.DriverBackWheelOffset)

	// Update back wheels rotation
	v.DriverBackWheel.Position.Rotation.Yaw = v.Body.Position.Rotation.Yaw
	v.PassengerBackWheel.Position.Rotation.Yaw = v.Body.Position.Rotation.Yaw - 180

	// Update front wheels rotation
	v.DriverFrontWheel.Position.Rotation.Yaw = v.Body.Position.Rotation.Yaw + v.Steering
	v.PassengerFrontWheel.Position.Rotation.Yaw = v.Body.Position.Rotation.Yaw - 180 + v.Steering

	// Update wheels speed (pitch)
	v.wheelRotation = float32(math.Mod(float64(v.wheelRotation+(v.Acceleration*40)), 360.0))
	v.DriverBackWheel.Position.Rotation.Pitch = -v.wheelRotation
	v.PassengerBackWheel.Position.Rotation.Pitch = v.wheelRotation
	v.DriverFrontWheel.Position.Rotation.Pitch = -v.wheelRotation
	v.PassengerFrontWheel.Position.Rotation.Pitch = v.wheelRotation

	// Apply transformations
	transform(&v.Body)
	transform(&v.DriverBackWheel)
	transform(&v.PassengerBackWheel)
	transform(&v.DriverFrontWheel)
	transform(&v.PassengerFrontWheel)
}

func (v *Vehicle) Render() {
	rl.DrawModel(v.Body.Object, v.Body.Position.Location, 1.0, rl.White)
	rl.DrawModel(v.PassengerFrontWheel.Object, v.PassengerFrontWheel.Position.Location, 1.0, rl.White)
	rl.DrawModel(v.PassengerBackWheel.Object, v.PassengerBackWheel.Position.Location, 1.0, rl.White)
	rl.DrawModel(v.DriverFrontWheel.Object, v.DriverFrontWheel.Position.Location, 1.0, rl.White)
	rl.DrawModel(v.DriverBackWheel.Object, v.DriverBackWheel.Position.Location, 1.0, rl.White)
}
