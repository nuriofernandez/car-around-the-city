package chat

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/nuriofernandez/car-around-the-city/settings"
)

var ChatOpen = false

func Render() {
	// Print last 10 chat messages
	messages := LastMessages()
	for i := 0; i < 10; i++ {
		rl.DrawText(messages[i], 10, int32(70+(i*20)), 20, rl.White)
	}

	// If chat is open, print the input text box
	if ChatOpen {
		InputChat()
		RenderInputBox()
	}

	// If chat is not open, check if key 't' is pressed
	if !ChatOpen && rl.IsKeyPressed(rl.KeyT) {
		ChatOpen = true
	}
}

func RenderInputBox() {
	// Command line
	rec := rl.Rectangle{
		X:      10,
		Y:      270,
		Width:  float32(settings.ScreenWidth - 200),
		Height: 30,
	}
	rl.DrawRectangleRounded(rec, 0.3, 0, rl.Black)
	// Draw command/message
	rl.DrawText(CurrentCommand, 15, 275, 20, rl.White)
}
