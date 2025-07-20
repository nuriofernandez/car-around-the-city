package render

import (
	"github.com/nuriofernandez/car-around-the-city/render/screens"
	"github.com/nuriofernandez/car-around-the-city/render/screens/game"
	"github.com/nuriofernandez/car-around-the-city/render/screens/loading"
)

func HandleScreen() {
	currentScreen := screens.CurrentScreen()
	if currentScreen == screens.LOADING {
		loading.Render()
		return
	}

	if currentScreen == screens.GAME {
		game.Render()
		return
	}
}
