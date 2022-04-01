package main

import (
	"io"
	"regexp"
	"strings"
)

type Options struct {
	caseInsensitive bool
}

func Search(w io.Writer, r io.Reader, term string, o Options) error {
	if term == "" {
		return nil
	}

	m := "(?)" + term
	if o.caseInsensitive {
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
			res = append(res, l)
		}
	}

	// Write results back to Writer
	if _, err := io.WriteString(w, strings.Join(res, "\n")+"\n"); err != nil {
		return err
	}
	return nil
}
