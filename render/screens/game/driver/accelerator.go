package driver

const (
	MAX_SPEED_FORWARD = 0.65
	MAX_SPEED_REVERSE = -0.3
)

var Acceleration = float32(0)

func Accelerate(speed float32) {
	Acceleration = maxSpeed(Acceleration + speed)
}

func Break() {
	if Acceleration < 0.05 && Acceleration > -0.05 {
		Acceleration = 0
	}
	if Acceleration < -0.05 { // REVERSE
		Acceleration = maxSpeed(Acceleration + 0.01)
	}
	if Acceleration > 0.05 { // FORWARD
		Acceleration = maxSpeed(Acceleration - 0.01)
	}
}

func maxSpeed(speed float32) float32 {
	if speed > MAX_SPEED_FORWARD {
		return MAX_SPEED_FORWARD
	}
	if speed < MAX_SPEED_REVERSE {
		return MAX_SPEED_REVERSE
	}
	return speed
}
