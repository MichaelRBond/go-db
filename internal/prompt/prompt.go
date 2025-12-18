package prompt

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/MichaelRBond/go-db/internal/commands"
	"github.com/MichaelRBond/go-db/internal/locations"
	"github.com/MichaelRBond/go-db/internal/player"
)

func RunPrompt(player *player.Player, rooms *locations.RoomsById) error {

	reader := bufio.NewReader(os.Stdin)

	for true {
		fmt.Printf("command /help for help: ")
		rawInput, _ := reader.ReadString('\n')
		requestedCommand, args, err := parseInput(rawInput)
		if err != nil {
			fmt.Println("Error parsing input:", err)
			continue
		}

		command, err := retrieveCommand(requestedCommand)
		if err != nil {
			fmt.Println("Error retrieving command:", err)
			continue
		}

		commandOutput, err := command.Execute(player, rooms, commands.ParsedCommand{
			Arg:  strings.Join(args, " "),
			Args: args,
		})

		if err != nil {
			fmt.Println("Error executing command:", err)
			continue
		}

		if commandOutput.Control == commands.ControlExit {
			fmt.Printf("\nGrandma says goodbye!\n")
			return nil
		}

		fmt.Printf("%s\n", commandOutput.Message)
	}

	return nil
}

func parseInput(input string) (string, []string, error) {
	trimmedInput := strings.TrimSpace(input)
	splitInput := strings.Fields(trimmedInput)

	if len(splitInput) == 0 {
		return "", nil, fmt.Errorf("no input provided")
	}

	return splitInput[0], splitInput[1:], nil
}

func retrieveCommand(input string) (commands.Command, error) {
	cmd, exists := commands.CommandList[input]
	if !exists {
		return commands.Command{}, fmt.Errorf("unknown command: %s", input)
	}
	return cmd, nil
}
