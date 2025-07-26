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
		PassengerFrontWheel: entity.Entity{
			Object: rl.LoadModel(model.WheelFile),
		},
		PassengerBackWheel: entity.Entity{
			Object: rl.LoadModel(model.WheelFile),
		},
		DriverFrontWheel: entity.Entity{
			Object: rl.LoadModel(model.WheelFile),
		},
		DriverBackWheel: entity.Entity{
			Object: rl.LoadModel(model.WheelFile),
		},
	}
	v.Update()

	return v
}
