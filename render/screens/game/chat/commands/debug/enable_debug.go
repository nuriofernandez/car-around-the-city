package debug

import (
	"github.com/nuriofernandez/car-around-the-city/render/screens/game/chat"
)

var GroundDebugEnabled = false

func EnableDebug(command string, args []string) {
	if len(args) == 0 {
		chat.Add("Please specify an option /debug <enable|disable>")
		return
	}

	enableDisable := args[0]
	if enableDisable == "enable" {
		GroundDebugEnabled = true
		chat.Add("Ground debug enabled")
		return
	}
	if enableDisable == "disable" {
		GroundDebugEnabled = false
		chat.Add("Ground debug disabled")
		return
	}

	// Invalid option
	chat.Add("Please specify a valid option /debug <enable|disable>")
}
