package grop

type Color int

const (
	Reset Color = iota
	Red
	Green
	Yellow
	Blue
	Purple
	Cyan
	Gray
	White
)

func colorize(s string, c Color) string {
	return c.String() + s + Reset.String()
}

func (c Color) String() string {
	switch c {
	case Reset:
		return "\033[0m"
	case Red:
		return "\033[31m"
	case Green:
		return "\033[32m"
	case Yellow:
		return "\033[33m"
	case Blue:
		return "\033[34m"
	case Purple:
		return "\033[35m"
	case Cyan:
		return "\033[36m"
	case Gray:
		return "\033[37m"
	case White:
		return "\033[97m"
	}
	return ""
}
