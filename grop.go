package grop

import (
	"io"
	"os"
	"regexp"
	"strings"
)

type Options struct {
	IgnoreCase bool
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

	res := []string{}

	// TODO: Optimize by reading one line at a time
	b, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	// split into lines
	lines := strings.Split(string(b), "\n")
	for _, l := range lines {
		// Collect all matching strings
		loc := reg.FindStringIndex(l)

		if loc != nil && loc[0] > -1 {
			res = append(res, colorizeMatch(l, loc[0], loc[1], Red))
		}
	}

	// Write results back to Writer
	if _, err := io.WriteString(w, strings.Join(res, "\n")+"\n"); err != nil {
		return err
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

func colorizeMatch(s string, start, end int, color Color) string {
	return s[:start] + color.String() + s[start:end] + Reset.String() + s[end:]
}
