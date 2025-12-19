package commands

import (
	"errors"

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
	room, exists := rooms.GetRoomById(player.Location)
	if !exists {
		return CommandReturn{}, errors.New("current room does not exist")
	}

	look := room.DisplayRoom()
	return CommandReturn{Message: look}, nil
}
