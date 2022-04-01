package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

var caseInsensitive bool

func init() {
	flag.BoolVar(&caseInsensitive, "i", false, "case insensitive match")
}

func main() {
	flag.Parse()
	fmt.Println("i has value", caseInsensitive)
	var in io.Reader
	if err := run(os.Args[1:], os.Stdout, in); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(args []string, w io.Writer, r io.Reader) error {
	term := args[0]
	opts := Options{caseInsensitive}

	fmt.Println(opts)
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
