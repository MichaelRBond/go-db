package commands

var Exit = Command{
	Name:        "/exit",
	Description: "Exits the application.",
	Help:        "Usage: /exit\nExits the application.",
	Execute:     exitFunction,
}

func exitFunction(cmd ParsedCommand) (CommandReturn, error) {
	return CommandReturn{Control: ControlExit}, nil
}
