package commands

import (
	"github.com/MichaelRBond/go-db/internal/locations"
	"github.com/MichaelRBond/go-db/internal/player"
)

var Look = Command{
	Name:        "look",
	Description: "Provides a description of the current location.",
	Help:        "Usage: look\nProvides a description of the current location.",
	Execute:     lookFunction,
}

func lookFunction(player *player.Player, rooms *locations.RoomsById, cmd ParsedCommand) (CommandReturn, error) {
	look := locations.DisplayRoom(rooms, player.Location)
	return CommandReturn{Message: look}, nil
}
