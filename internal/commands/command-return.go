package commands

type CommandReturnControl string

const (
	ControlExit CommandReturnControl = "EXIT"
	ControlNone CommandReturnControl = "NONE"
)

type CommandReturn struct {
	Message string
	Control CommandReturnControl
}
