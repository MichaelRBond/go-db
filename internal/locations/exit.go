package locations

type Direction string

const (
	North Direction = "north"
	South Direction = "south"
	East  Direction = "east"
	West  Direction = "west"
	Up    Direction = "up"
	Down  Direction = "down"
)

type ExitType string

const (
	ExitTypeNone ExitType = "NONE"
	ExitTypeDoor ExitType = "door"
)

type ExitState string

const (
	ExitStateNone   ExitState = "NONE"
	ExitStateOpen   ExitState = "open"
	ExitStateClosed ExitState = "closed"
)

type ExitDescription struct {
	Open   string `json:"open"`
	Closed string `json:"closed"`
}

type RoomExit struct {
	Direction   Direction       `json:"direction"`
	RoomID      string          `json:"roomId"`
	Type        ExitType        `json:"type"`
	State       ExitState       `json:"state"`
	Description ExitDescription `json:"description"`
}
