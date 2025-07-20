package screens

const (
	LOADING = 1
	GAME    = 2
)

var currentScreen = LOADING

func SetScreen(newScreen int) {
	currentScreen = newScreen
}

func CurrentScreen() int {
	return currentScreen
}
