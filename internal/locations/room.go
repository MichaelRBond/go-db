package locations

import (
	"fmt"
	"strings"
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
	output := fmt.Sprintf("\n%s\n%s\n", room.Name, room.Description)
	exits := room.GetExitSliceAsStrings()

	output += fmt.Sprintf("Exits: [%s]\n", strings.Join(exits, ", "))

	return output
}
