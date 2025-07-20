package game

import rl "github.com/gen2brain/raylib-go/raylib"

func Load() {
	LoadMap()
	rl.DisableCursor()
}

func UnLoad() {
	rl.EnableCursor()
}
