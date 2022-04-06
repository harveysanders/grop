package grop

import (
	"bufio"
	"io"
	"os"
	"regexp"
)

type Options struct {
	IgnoreCase    bool
	WhenHighlight string // TODO: Use enum
}

func Search(w io.Writer, r io.Reader, term string, o Options) error {
	if term == "" {
		return nil
	}

	m := "(?)" + term
	if o.IgnoreCase {
		m = "(?i)" + term
	}
	reg, err := regexp.Compile(m)
	if err != nil {
		return err
	}

	s := bufio.NewScanner(r)
	for s.Scan() {
		l := s.Text()
		loc := reg.FindStringIndex(l)

		if loc != nil && loc[0] > -1 {
			// Write highlighted match back to Writer
			match := highlightMatch(l, loc[0], loc[1], Red, o.WhenHighlight)
			_, err := io.WriteString(w, match+"\n")
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func Run(args []string, w io.Writer, r io.Reader, opts Options) error {
	term := args[0]
	if len(args) == 1 {
		// Use stdin
		err := Search(os.Stdout, os.Stdin, term, opts)
		if err != nil {
			return err
		}
		return nil
	}

	// More than one arg, assume last arg is path to file
	fp := args[len(args)-1]
	file, err := os.Open(fp)
	if err != nil {
		return err
	}
	defer file.Close()

	err = Search(os.Stdout, file, term, opts)
	if err != nil {
		return err
	}
	return nil
}

func highlightMatch(s string, start, end int, color Color, when string) string {
	switch when {
	case "always":
		return s[:start] + color.String() + s[start:end] + Reset.String() + s[end:]
	case "auto":
		// TODO: Probably need to do something different here
		return s[:start] + color.String() + s[start:end] + Reset.String() + s[end:]
	case "never":
		return s
	}
	return s
}
