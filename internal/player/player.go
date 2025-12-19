package player

type Player struct {
	Location string
}

func InitPlayer() *Player {
	return &Player{
		Location: "grandmas_kitchen",
	}
}

func (p *Player) SetLocation(locationId string) error {
	p.Location = locationId
	return nil
}
