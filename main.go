package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/nuriofernandez/car-around-the-city/render"
	"github.com/nuriofernandez/car-around-the-city/render/screens/loading"
	"github.com/nuriofernandez/car-around-the-city/settings"
)

func main() {

	rl.InitWindow(settings.ScreenWidth, settings.ScreenHeight, "Car around the city")
	defer rl.CloseWindow()
	rl.SetTraceLogLevel(rl.LogAll)

	icon := rl.LoadImage("resources/icon.png")
	rl.SetWindowIcon(*icon)

	rl.SetTargetFPS(60)

	loading.Load()
	for !rl.WindowShouldClose() {
		render.HandleScreen()
	}
}
