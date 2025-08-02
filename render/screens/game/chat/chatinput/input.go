package chatinput

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var ChatOpen = false
var CurrentCommand = ""

func InputChat() {
	// Handle typing
	char := rl.GetCharPressed()
	if (char >= 32) && (char <= 125) {
		CurrentCommand = CurrentCommand + getString(char)
	}

	// Handle delete
	var backspacePressed = rl.IsKeyPressedRepeat(rl.KeyBackspace) || rl.IsKeyPressed(rl.KeyBackspace)
	if backspacePressed && len(CurrentCommand) > 0 {
		CurrentCommand = CurrentCommand[:len(CurrentCommand)-1]
	}

	// Handle send
	if rl.IsKeyPressed(rl.KeyEnter) {
		Process(CurrentCommand)
		ChatOpen = false
		CurrentCommand = ""
	}
}

func getString(r rune) string {
	var s string
	if r == 0 {
		return s
	}
	s += string(r)
	return s
}
