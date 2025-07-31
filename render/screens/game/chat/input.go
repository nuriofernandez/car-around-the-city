package chat

import rl "github.com/gen2brain/raylib-go/raylib"

var CurrentCommand = ""

func InputChat() {
	char := rl.GetCharPressed()
	if (char >= 32) && (char <= 125) {
		CurrentCommand = CurrentCommand + getString(char)
	}
	if rl.IsKeyDown(rl.KeyBackspace) && len(CurrentCommand) > 0 {
		CurrentCommand = CurrentCommand[:len(CurrentCommand)-1]
	}
	if rl.IsKeyPressed(rl.KeyEnter) {
		Add(CurrentCommand)
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
