package commands

import (
	"github.com/nuriofernandez/car-around-the-city/render/screens/game/chat/commands/debug"
	"github.com/nuriofernandez/car-around-the-city/render/screens/game/chat/commands/notfound"
	teleport "github.com/nuriofernandez/car-around-the-city/render/screens/game/chat/commands/tp"
	"strings"
)

func IsCommand(message string) bool {
	return strings.HasPrefix(message, "/")
}

func GetExecutor(command string) func(command string, args []string) {
	if command == "debug" {
		return debug.EnableDebug
	}
	if command == "tp" {
		return teleport.Teleport
	}

	return notfound.Execute
}

func ExecuteCommand(message string) {
	messageSplit := strings.Split(message, " ")
	command := messageSplit[0][1:]
	arguments := messageSplit[1:]

	executor := GetExecutor(command)
	executor(command, arguments)
}
