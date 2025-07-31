package chatrender

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/nuriofernandez/car-around-the-city/render/screens/game/chat"
	"github.com/nuriofernandez/car-around-the-city/render/screens/game/chat/chatinput"
	"github.com/nuriofernandez/car-around-the-city/settings"
)

func Render() {
	// Print last 10 chat messages
	messages := chat.LastMessages()
	for i := 0; i < 10; i++ {
		rl.DrawText(messages[i], 10, int32(10+(i*20)), 20, rl.White)
	}

	// If chat is open, print the input text box
	if chatinput.ChatOpen {
		chatinput.InputChat()
		RenderInputBox()
	}

	// If chat is not open, check if key 't' is pressed
	if !chatinput.ChatOpen && rl.IsKeyPressed(rl.KeyT) {
		chatinput.ChatOpen = true
	}
}

func RenderInputBox() {
	// Command line
	rec := rl.Rectangle{
		X:      10,
		Y:      210,
		Width:  float32(settings.ScreenWidth - 200),
		Height: 30,
	}
	rl.DrawRectangleRounded(rec, 0.3, 0, rl.Black)

	// Draw command/message
	rl.DrawText(chatinput.CurrentCommand, 15, 215, 20, rl.White)
}
