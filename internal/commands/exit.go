package commands

import (
	"github.com/MichaelRBond/go-db/internal/locations"
	"github.com/MichaelRBond/go-db/internal/player"
)

var Exit = Command{
	Name:        "/exit",
	Description: "Exits the application.",
	Help:        "Usage: /exit\nExits the application.",
	Execute:     exitFunction,
}

func exitFunction(player *player.Player, rooms *locations.RoomsById, cmd ParsedCommand) (CommandReturn, error) {
	return CommandReturn{Control: ControlExit}, nil
}
