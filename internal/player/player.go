package player

type Player struct {
	Location string
}

func InitPlayer() *Player {
	return &Player{
		Location: "grandmas_kitchen",
	}
}
