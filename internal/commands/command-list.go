package commands

var CommandList map[string]Command

func init() {
	CommandList = map[string]Command{
		"look":  Look,
		"move":  Move,
		"say":   Say,
		"/exit": Exit,
		"/help": Help,
	}
}
