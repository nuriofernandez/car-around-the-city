package vehicle

import "github.com/nuriofernandez/car-around-the-city/render/screens/game/driver"

func PreRender() {
	EngineSim()

	PedVehicle.Acceleration = driver.Acceleration
	PedVehicle.Steering = driver.Steering
}

func Render() {
	PedVehicle.Render()

	HandleDirection()
}
