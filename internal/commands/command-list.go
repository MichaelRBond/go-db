package commands

var CommandList map[string]Command

func init() {
	CommandList = map[string]Command{
		"look":  Look,
		"move":  Move,
		"north": MoveNorth,
		"n":     MoveNorth,
		"south": MoveSouth,
		"east":  MoveEast,
		"west":  MoveWest,
		"up":    MoveUp,
		"say":   Say,
		"/exit": Exit,
		"/help": Help,
	}
}
