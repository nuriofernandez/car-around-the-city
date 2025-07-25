package vehicle

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/nuriofernandez/car-around-the-city/render/screens/game/driver"
)

func GetCarPos() *rl.Vector3 {
	return &PedVehicle.Body.Position.Location
}

func GetDebug() string {
	return fmt.Sprintf("A: %.1f | S: %.2f", driver.Acceleration, driver.Steering)
}
