package commands

import (
	"fmt"
	"strings"

	"github.com/MichaelRBond/go-db/internal/locations"
	"github.com/MichaelRBond/go-db/internal/player"
)

var Say = Command{
	Name:        "say",
	Description: "Outputs the provided message.",
	Help:        "Usage: say <message>\nOutputs the specified message to the current room.",
	Execute:     sayFunction,
}

func sayFunction(player *player.Player, rooms *locations.RoomsById, cmd ParsedCommand) (CommandReturn, error) {
	var ret = CommandReturn{
		Control: ControlNone,
	}

	if len(cmd.Args) == 0 {
		return ret, fmt.Errorf("no message provided")
	}

	ret.Message = strings.Join(cmd.Args, " ")
	return ret, nil
}
