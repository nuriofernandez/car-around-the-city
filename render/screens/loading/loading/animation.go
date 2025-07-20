package loading

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"time"
)

type LoadingAnimation struct {
	PositionX int32
	PositionY int32
	Loading   bool
	start     time.Time
}

func (l *LoadingAnimation) Render() {
	// Not loading, should not render
	if l.Loading == false {
		return
	}

	// if start time is not defined, define it
	if l.start == (time.Time{}) {
		l.start = time.Now()
	}

	seconds := int(time.Now().Sub(l.start).Milliseconds() / 300)
	stage := seconds % 4
	if stage == 0 {
		rl.DrawRectangle(l.PositionX, l.PositionY, 16, 16, rl.Black)
		rl.DrawRectangle(l.PositionX+20, l.PositionY, 16, 16, rl.Gray)
		rl.DrawRectangle(l.PositionX+40, l.PositionY, 16, 16, rl.Gray)
	} else if stage == 1 {
		rl.DrawRectangle(l.PositionX, l.PositionY, 16, 16, rl.Gray)
		rl.DrawRectangle(l.PositionX+20, l.PositionY, 16, 16, rl.Black)
		rl.DrawRectangle(l.PositionX+40, l.PositionY, 16, 16, rl.Gray)
	} else if stage == 2 {
		rl.DrawRectangle(l.PositionX, l.PositionY, 16, 16, rl.Gray)
		rl.DrawRectangle(l.PositionX+20, l.PositionY, 16, 16, rl.Gray)
		rl.DrawRectangle(l.PositionX+40, l.PositionY, 16, 16, rl.Black)
	} else if stage == 3 {
		rl.DrawRectangle(l.PositionX, l.PositionY, 16, 16, rl.Gray)
		rl.DrawRectangle(l.PositionX+20, l.PositionY, 16, 16, rl.Gray)
		rl.DrawRectangle(l.PositionX+40, l.PositionY, 16, 16, rl.Gray)
	}
}
