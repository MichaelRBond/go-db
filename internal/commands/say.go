package commands

import (
	"fmt"
	"strings"
)

var Say = Command{
	Name:        "say",
	Description: "Outputs the provided message.",
	Help:        "Usage: say <message>\nOutputs the specified message to the current room.",
	Execute:     sayFunction,
}

func sayFunction(cmd ParsedCommand) (CommandReturn, error) {
	var ret = CommandReturn{
		Control: ControlNone,
	}

	if len(cmd.Args) == 0 {
		return ret, fmt.Errorf("no message provided")
	}

	ret.Message = strings.Join(cmd.Args, " ")
	return ret, nil
}
