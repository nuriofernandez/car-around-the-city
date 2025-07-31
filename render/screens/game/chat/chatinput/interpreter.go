package chatinput

import (
	"github.com/nuriofernandez/car-around-the-city/render/screens/game/chat"
	"github.com/nuriofernandez/car-around-the-city/render/screens/game/chat/commands"
)

func Process(message string) {
	if commands.IsCommand(message) {
		commands.ExecuteCommand(message)
		return
	}
	chat.Add(message)
}
