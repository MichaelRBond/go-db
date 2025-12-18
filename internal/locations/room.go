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

func GetRoomById(rooms *RoomsById, id string) (*Room, bool) {
	room, exists := (*rooms)[id]
	return room, exists
}

func GetExit(room *Room, direction string) (*RoomExit, bool) {
	for _, exit := range room.Exits {
		if string(exit.Direction) == direction {
			return &exit, true
		}
	}
	return nil, false
}

func GetExitSliceAsStrings(room *Room) []string {
	exits := []string{}
	for _, exit := range room.Exits {
		exits = append(exits, string(exit.Direction))
	}
	return exits
}

func DisplayRoom(rooms *RoomsById, roomId string) string {
	room, exists := GetRoomById(rooms, roomId)
	if !exists {
		return fmt.Sprintf("Unknown location, id: %s", roomId)
	}

	output := fmt.Sprintf("\n%s\n%s\n", room.Name, room.Description)
	exits := GetExitSliceAsStrings(room)

	output += fmt.Sprintf("Exits: [%s]\n", strings.Join(exits, ", "))

	return output
}
