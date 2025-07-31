package notfound

import (
	"github.com/nuriofernandez/car-around-the-city/render/screens/game/chat"
)

func Execute(command string, args []string) {
	chat.Add("Command '" + command + "' not found!")
}
