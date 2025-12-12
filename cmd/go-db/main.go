package main

import (
	"fmt"

	"github.com/MichaelRBond/go-db/internal/db"
	"github.com/MichaelRBond/go-db/internal/prompt"
)

func main() {
	fmt.Printf("Starting Grandma’s Old Dungeon Brawl\n\n")
	if error := db.InitializeDB(); error != nil {
		fmt.Println("Error initializing DB:", error)
	}
	if error := prompt.RunPrompt(); error != nil {
		fmt.Println("Error running UI:", error)
	}
}
