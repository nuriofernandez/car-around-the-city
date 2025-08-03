package teleport

import (
	"github.com/nuriofernandez/car-around-the-city/render/screens/game/chat"
	"github.com/nuriofernandez/car-around-the-city/render/screens/game/vehicle"
	"strconv"
)

func Teleport(command string, args []string) {
	if len(args) != 3 {
		chat.Add("Please specify an option /tp <x> <y> <z>")
		return
	}

	x, err := strconv.Atoi(args[0])
	if err != nil {
		chat.Add("Invalid argument <x>! Must be a number")
	}

	y, err := strconv.Atoi(args[1])
	if err != nil {
		chat.Add("Invalid argument <y>! Must be a number")
	}

	z, err := strconv.Atoi(args[2])
	if err != nil {
		chat.Add("Invalid argument <z>! Must be a number")
	}

	// Invalid option
	chat.Add("Teleporting to '" + args[0] + "," + args[1] + "," + args[2] + "'...")
	vehicle.PedVehicle.Body.Position.Location.X = float32(x)
	vehicle.PedVehicle.Body.Position.Location.Y = float32(y)
	vehicle.PedVehicle.Body.Position.Location.Z = float32(z)
}
