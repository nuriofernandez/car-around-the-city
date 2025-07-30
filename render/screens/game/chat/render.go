package chat

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Render() {
	messages := LastMessages()
	for i := 0; i < 10; i++ {
		rl.DrawText(messages[i], 10, int32(70+(i*15)), 17, rl.White)
	}
}
