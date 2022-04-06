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
	isStdout      bool   // Needed to determine if colors should print
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

	sc := bufio.NewScanner(r)
	for sc.Scan() {
		line := sc.Text()
		matches := reg.FindAllString(line, -1)
		if matches == nil {
			continue
		}

		// Handle color option
		res := colorize(line, matches, o.WhenHighlight, o.isStdout)
		_, err := io.WriteString(w, res+"\n")
		if err != nil {
			return err
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
