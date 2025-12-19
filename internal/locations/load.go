package locations

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var RoomsDirectory = "data/rooms"

type RoomsById map[string]*Room

func (rooms *RoomsById) GetRoomById(id string) (*Room, bool) {
	room, exists := (*rooms)[id]
	return room, exists
}

func LoadRooms() (*RoomsById, error) {
	rooms := make(RoomsById)

	files, err := os.ReadDir(RoomsDirectory)
	if err != nil {
		fmt.Printf("Error reading rooms directory: %v\n", err)
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if !strings.HasSuffix(file.Name(), ".json") {
			continue
		}

		path := filepath.Join(RoomsDirectory, file.Name())
		data, err := os.ReadFile(path)
		if err != nil {
			return nil, fmt.Errorf("reading room file %q: %w", path, err)
		}

		var room Room
		if err := json.Unmarshal(data, &room); err != nil {
			return nil, fmt.Errorf("parsing room JSON %q: %w", path, err)
		}

		if room.Id == "" {
			return nil, fmt.Errorf("room in %q has empty id", path)
		}

		if _, exists := rooms[room.Id]; exists {
			return nil, fmt.Errorf("duplicate room id %q (file %q)", room.Id, path)
		}

		rooms[room.Id] = &room

	}

	fmt.Printf("Loaded %d rooms\n", len(rooms))

	return &rooms, nil
}
