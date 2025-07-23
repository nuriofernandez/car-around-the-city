package vehicle

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/nuriofernandez/car-around-the-city/render/screens/game/models/entity"
	"github.com/nuriofernandez/car-around-the-city/render/screens/game/models/position"
)

func NewVehicle(x, y, z float32) Vehicle {
	var model = hardcodedVehicleModel
	var pos = position.NewPosition(x, y, z, 0, 0, 0)

	v := Vehicle{
		Model: model,
		Body: entity.Entity{
			Object:   rl.LoadModel(model.BodyFile),
			Position: pos,
		},
		passengerFrontWheel: entity.Entity{
			Object: rl.LoadModel(model.WheelFile),
		},
		passengerBackWheel: entity.Entity{
			Object: rl.LoadModel(model.WheelFile),
		},
		driverFrontWheel: entity.Entity{
			Object: rl.LoadModel(model.WheelFile),
		},
		driverBackWheel: entity.Entity{
			Object: rl.LoadModel(model.WheelFile),
		},
	}
	v.Update()

	return v
}
