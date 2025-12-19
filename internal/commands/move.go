package commands

import (
	"errors"

	"github.com/MichaelRBond/go-db/internal/locations"
	"github.com/MichaelRBond/go-db/internal/player"
)

var Move = Command{
	Name:        "move",
	Description: "Move to a different location.",
	Help:        "Usage: move <direction>\nMove to a different location in the dungeon.",
	Execute:     moveFunction,
}

func moveFunction(player *player.Player, rooms *locations.RoomsById, cmd ParsedCommand) (CommandReturn, error) {
	room, roomExists := locations.GetRoomById(rooms, player.Location)
	if !roomExists {
		return CommandReturn{}, errors.New("current location not found")
	}

	exit, exitExits := locations.GetExit(room, cmd.Args[0])
	if !exitExits {
		return CommandReturn{}, errors.New("invalid direction")
	}

	player.SetLocation(exit.RoomID)

	ret := CommandReturn{
		Message: "You move " + cmd.Args[0] + "." + "\n" + locations.DisplayRoom(rooms, player.Location),
	}
	return ret, nil
}
