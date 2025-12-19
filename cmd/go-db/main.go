package main

import (
	"fmt"
	"os"

	"github.com/MichaelRBond/go-db/internal/commands"
	"github.com/MichaelRBond/go-db/internal/db"
	"github.com/MichaelRBond/go-db/internal/locations"
	"github.com/MichaelRBond/go-db/internal/player"
	"github.com/MichaelRBond/go-db/internal/prompt"
)

func main() {
	fmt.Printf("Starting Grandma’s Old Dungeon Brawl\n\n")

	if error := db.InitializeDB(); error != nil {
		fmt.Println("Error initializing DB:", error)
	}

	rooms, error := locations.LoadRooms()
	if error != nil {
		fmt.Println("Error loading rooms:", error)
		os.Exit(1)
	}

	player := player.InitPlayer()

	initialLookResult, _ := commands.Look.Execute(player, rooms, commands.ParsedCommand{})
	fmt.Printf("\n%s\n\n", initialLookResult.Message)

	if error := prompt.RunPrompt(player, rooms); error != nil {
		fmt.Println("Error running UI:", error)
	}
}
