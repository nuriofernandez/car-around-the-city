package loading

import rl "github.com/gen2brain/raylib-go/raylib"

var backgroundTexture rl.Texture2D

func Load() {
	backgroundImage := rl.LoadImage("resources/loading-background.png")
	backgroundTexture = rl.LoadTextureFromImage(backgroundImage)
}
