package commands

var CommandList map[string]Command

func init() {
	CommandList = map[string]Command{
		"look":  Look,
		"move":  Move,
		"north": MoveNorth,
		"south": MoveSouth,
		"east":  MoveEast,
		"west":  MoveWest,
		"down":  MoveDown,
		"up":    MoveUp,
		"say":   Say,
		"/exit": Exit,
		"/help": Help,
	}
}
