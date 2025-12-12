package commands

var CommandList map[string]Command

func init() {
	CommandList = map[string]Command{
		"say":   Say,
		"/exit": Exit,
		"/help": Help,
	}
}
