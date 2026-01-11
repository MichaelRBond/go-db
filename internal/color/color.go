package color

const (
	Reset = "\033[0m"

	Yellow = "\033[33m"
	White  = "\033[37m"
	Red    = "\033[31m"
	Blue   = "\033[34m"

	BoldYellow = "\033[1;33m"
	BoldWhite  = "\033[1;37m"
	BoldRed    = "\033[1;31m"
	BoldBlue   = "\033[1;34m"
)

func Wrap(text, colorCode string) string {
	return colorCode + text + Reset
}
