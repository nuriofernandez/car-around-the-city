package loading

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/nuriofernandez/car-around-the-city/render/screens"
	"github.com/nuriofernandez/car-around-the-city/render/screens/game"
	"github.com/nuriofernandez/car-around-the-city/render/screens/loading/loading"
	"github.com/nuriofernandez/car-around-the-city/settings"
	"time"
)

var state = 0 // Tracking animation states (State Machine)
var start = time.Now()

var loadingAnim = loading.LoadingAnimation{
	Loading:   true,
	PositionX: settings.ScreenWidth - ((56) + 8),
	PositionY: settings.ScreenHeight - (16 + 8),
}

func Render() {
	// Control state based on time
	if state == 0 { // Preload
		if time.Now().After(start.Add(time.Second * 2)) {
			start = time.Now()
			state = 1
		}
	} else if state == 2 { // Actual load
		if time.Now().After(start.Add(time.Second * 2)) {
			start = time.Now()
			loadingAnim.Loading = false
			state = 3
		}
	} else if state == 3 { // Post load
		if time.Now().After(start.Add(time.Second * 2)) {
			start = time.Now()
			state = 4
		}
	}

	// Control loading state (1)
	if state == 1 {
		game.Load()
		start = time.Now()
		state = 2
	}

	// Control end state (start game)
	if state == 4 {
		defer screens.SetScreen(screens.GAME)
	}

	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)
	rl.DrawTexture(backgroundTexture, 0, 0, rl.White)

	// Loading animation
	loadingAnim.Render()

	if state == 3 {
		milliseconds := time.Now().Sub(start).Milliseconds()
		alpha := float32(milliseconds) / 2000

		text := "Car around the city"
		rl.DrawText(text, 20, 20, 50, rl.Fade(rl.Black, alpha))
	}

	rl.EndDrawing()
}
