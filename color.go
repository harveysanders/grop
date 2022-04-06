package grop

import "strings"

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

// colorize determines if matches should be highlight based on "--color" option
// and returns a string with or without colors.
// If color option is "auto", matches are only highlighted if
// output it set to stdout.
func colorize(s string, matches []string, when string, isStdout bool) string {
	switch when {
	case "always":
		return highlightMatches(s, matches)
	case "auto":
		if !isStdout {
			return s
		}
		return highlightMatches(s, matches)
	}
	return s
}

// highlightMatches returns a new string with all matches wrapped in
// ANSI color characters
func highlightMatches(s string, matches []string) string {
	for _, m := range matches {
		s = strings.Replace(s, m, Red.String()+m+Reset.String(), 1)
	}
	return s
}
