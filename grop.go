package main

import (
	"io"
	"os"
	"strings"
)

func Search(w io.Writer, r io.Reader, term string) error {
	if term == "" {
		return nil
	}

	res := []string{}
	b, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	// split into lines
	lines := strings.Split(string(b), "\n")
	for _, l := range lines {
		// Collect all matching strings
		if strings.Contains(l, term) {
			res = append(res, l)
		}
	}

	// Write results back to Writer
	if _, err := io.WriteString(w, strings.Join(res, "\n")+"\n"); err != nil {
		return err
	}
	return nil
}

func run(args []string, w io.Writer, r io.Reader) error {
	term := args[0]

	if len(args) == 1 {
		// Use stdin
		err := Search(os.Stdout, os.Stdin, term)
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

	err = Search(os.Stdout, file, term)
	if err != nil {
		return err
	}
	return nil
}
