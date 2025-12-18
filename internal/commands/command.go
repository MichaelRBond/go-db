package commands

import (
	"github.com/MichaelRBond/go-db/internal/locations"
	"github.com/MichaelRBond/go-db/internal/player"
)

type Command struct {
	Name        string
	Description string
	Help        string
	Execute     func(player *player.Player, rooms *locations.RoomsById, cmd ParsedCommand) (CommandReturn, error)
}
