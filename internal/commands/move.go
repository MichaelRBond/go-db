package commands

import (
	"errors"
	"fmt"

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
	room, roomExists := rooms.GetRoomById(player.Location)
	if !roomExists {
		return CommandReturn{}, errors.New("current location not found")
	}

	exit, exitExits := room.GetExit(cmd.Args[0])
	if !exitExits {
		return CommandReturn{}, errors.New("invalid direction")
	}

	newRoom, newRoomExists := rooms.GetRoomById(exit.RoomID)
	if !newRoomExists {
		return CommandReturn{}, errors.New("destination room not found")
	}
	player.SetLocation(exit.RoomID)

	ret := CommandReturn{
		Message: fmt.Sprintf("You move %s.\n%s", cmd.Args[0], newRoom.DisplayRoom()),
	}
	return ret, nil
}

func moveNorthFunction(player *player.Player, rooms *locations.RoomsById, cmd ParsedCommand) (CommandReturn, error) {
	return moveFunction(player, rooms, ParsedCommand{Args: []string{"north"}})
}

func moveSouthFunction(player *player.Player, rooms *locations.RoomsById, cmd ParsedCommand) (CommandReturn, error) {
	return moveFunction(player, rooms, ParsedCommand{Args: []string{"south"}})
}

func moveEastFunction(player *player.Player, rooms *locations.RoomsById, cmd ParsedCommand) (CommandReturn, error) {
	return moveFunction(player, rooms, ParsedCommand{Args: []string{"east"}})
}

func moveWestFunction(player *player.Player, rooms *locations.RoomsById, cmd ParsedCommand) (CommandReturn, error) {
	return moveFunction(player, rooms, ParsedCommand{Args: []string{"west"}})
}

func moveUpFunction(player *player.Player, rooms *locations.RoomsById, cmd ParsedCommand) (CommandReturn, error) {
	return moveFunction(player, rooms, ParsedCommand{Args: []string{"up"}})
}

func moveDownFunction(player *player.Player, rooms *locations.RoomsById, cmd ParsedCommand) (CommandReturn, error) {
	return moveFunction(player, rooms, ParsedCommand{Args: []string{"down"}})
}
