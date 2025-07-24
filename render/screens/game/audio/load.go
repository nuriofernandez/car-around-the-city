package audio

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/nuriofernandez/car-around-the-city/render/screens/game/driver"
)

var engineLoop rl.Music
var accelerateLoop rl.Music

var accelerate rl.Sound
var decelerate rl.Sound

func Load() {
	rl.InitAudioDevice()

	engineLoop = rl.LoadMusicStream("resources/engine-loop.wav")
	engineLoop.Looping = true
	rl.PlayMusicStream(engineLoop)
	accelerateLoop = rl.LoadMusicStream("resources/accelerate-loop.wav")
	accelerateLoop.Looping = true
	rl.PlayMusicStream(accelerateLoop)

	accelerate = rl.LoadSound("resources/accelerate.mp3")
	decelerate = rl.LoadSound("resources/decelerate.mp3")
}

var prevSpeed = float32(0)
var isAccelerating = false
var isDecelerating = false

func Loop() {
	rl.UpdateMusicStream(engineLoop)

	if driver.Acceleration >= driver.MAX_SPEED_FORWARD-0.03 {
		isAccelerating = true
		isDecelerating = false
		rl.UpdateMusicStream(accelerateLoop)
		prevSpeed = driver.Acceleration
		return
	}

	// Now is faster, accelerating
	if prevSpeed < driver.Acceleration {
		if !isAccelerating {
			isAccelerating = true
			rl.StopSound(decelerate)
			rl.PlaySound(accelerate)
			fmt.Println("AC!")
		} else {
			if !rl.IsSoundPlaying(accelerate) {
				rl.UpdateMusicStream(accelerateLoop)
			}
		}
		isDecelerating = false
	}

	// Now is slower, decelerating
	if prevSpeed > driver.Acceleration {
		if !isDecelerating {
			isDecelerating = true
			rl.StopSound(accelerate)
			rl.PlaySound(decelerate)
			fmt.Println("DC!")
		}
		isAccelerating = false
	}

	prevSpeed = driver.Acceleration

}
