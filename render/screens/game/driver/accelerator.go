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
	if Acceleration < 0.1 && Acceleration > -0.1 {
		Acceleration = 0
	}
	if Acceleration < -0.1 { // REVERSE
		Acceleration = maxSpeed(Acceleration + 0.02)
	}
	if Acceleration > 0.1 { // FORWARD
		Acceleration = maxSpeed(Acceleration - 0.02)
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
