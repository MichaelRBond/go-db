package locations

import (
	"fmt"
	"strings"

	"github.com/MichaelRBond/go-db/internal/color"
)

type Room struct {
	Id          string     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Exits       []RoomExit `json:"exits"`
}

func (room *Room) GetExit(direction string) (*RoomExit, bool) {
	for _, exit := range room.Exits {
		if string(exit.Direction) == direction {
			return &exit, true
		}
	}
	return nil, false
}

func (room *Room) GetExitSliceAsStrings() []string {
	exits := []string{}
	for _, exit := range room.Exits {
		exits = append(exits, string(exit.Direction))
	}
	return exits
}

func (room *Room) DisplayRoom() string {
	output := fmt.Sprintf("\n%s\n\n%s\n",
		color.Wrap(room.Name, color.Yellow),
		room.Description,
	)
	exits := room.GetExitSliceAsStrings()

	output += fmt.Sprintf("\n%s [%s]\n", color.Wrap("Exits:", color.Blue), strings.Join(exits, ", "))

	return output
}
