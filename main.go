package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var in io.Reader
	term := os.Args[1]

	// Use file as first input
	fpath := os.Args[len(os.Args)-1]
	if fpath == "" {
		// Then fallback to stdin
		in = os.Stdin
	}

	if err := Search(os.Stdout, in, term); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
