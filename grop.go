package main

import (
	"fmt"
	"io"
	"os"
)

func Search(w io.Writer, r io.Reader, term string) error {

	return nil
}

func run(args []string, w io.Writer, r io.Reader) error {
	term := args[0]
	fmt.Println(term)

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
