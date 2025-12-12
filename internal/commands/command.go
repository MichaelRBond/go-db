package commands

type Command struct {
	Name        string
	Description string
	Help        string
	Execute     func(cmd ParsedCommand) (CommandReturn, error)
}
