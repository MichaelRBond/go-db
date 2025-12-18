package commands

import (
	"fmt"

	"github.com/MichaelRBond/go-db/internal/locations"
	"github.com/MichaelRBond/go-db/internal/player"
)

var Help = Command{
	Name:        "/help",
	Description: "Provides help information for commands.",
	Help:        "Usage: /help [command]\nIf a command is specified, detailed help for that command is provided.",
	Execute:     helpFunction,
}

func helpFunction(player *player.Player, rooms *locations.RoomsById, cmd ParsedCommand) (CommandReturn, error) {
	fmt.Println("Available Commands:")
	for name, command := range CommandList {
		fmt.Printf("%s: %s\n", name, command.Description)
	}

	return CommandReturn{}, nil
}
