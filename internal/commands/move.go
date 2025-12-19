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

var MoveNorth = Command{
	Name:        "north",
	Description: "Move to the north.",
	Help:        "Usage: north <direction>\nMove through the north exit",
	Execute:     moveNorthFunction,
}

var MoveSouth = Command{
	Name:        "south",
	Description: "Move to the south.",
	Help:        "Usage: south <direction>\nMove through the south exit",
	Execute:     moveSouthFunction,
}

var MoveEast = Command{
	Name:        "east",
	Description: "Move to the east.",
	Help:        "Usage: east <direction>\nMove through the east exit",
	Execute:     moveEastFunction,
}

var MoveWest = Command{
	Name:        "west",
	Description: "Move to the west.",
	Help:        "Usage: west <direction>\nMove through the west exit",
	Execute:     moveWestFunction,
}

var MoveUp = Command{
	Name:        "up",
	Description: "Move up.",
	Help:        "Usage: up <direction>\nMove through the up exit",
	Execute:     moveUpFunction,
}

var MoveDown = Command{
	Name:        "down",
	Description: "Move down.",
	Help:        "Usage: down <direction>\nMove through the down exit",
	Execute:     moveDownFunction,
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

func moveNorthFunction(player *player.Player, rooms *locations.RoomsById, cmd ParsedCommand) (CommandReturn, error) {
	cmd.Args = []string{"north"}
	return moveFunction(player, rooms, cmd)
}

func moveSouthFunction(player *player.Player, rooms *locations.RoomsById, cmd ParsedCommand) (CommandReturn, error) {
	cmd.Args = []string{"south"}
	return moveFunction(player, rooms, cmd)
}

func moveEastFunction(player *player.Player, rooms *locations.RoomsById, cmd ParsedCommand) (CommandReturn, error) {
	cmd.Args = []string{"east"}
	return moveFunction(player, rooms, cmd)
}

func moveWestFunction(player *player.Player, rooms *locations.RoomsById, cmd ParsedCommand) (CommandReturn, error) {
	cmd.Args = []string{"west"}
	return moveFunction(player, rooms, cmd)
}

func moveUpFunction(player *player.Player, rooms *locations.RoomsById, cmd ParsedCommand) (CommandReturn, error) {
	cmd.Args = []string{"up"}
	return moveFunction(player, rooms, cmd)
}

func moveDownFunction(player *player.Player, rooms *locations.RoomsById, cmd ParsedCommand) (CommandReturn, error) {
	cmd.Args = []string{"down"}
	return moveFunction(player, rooms, cmd)
}
