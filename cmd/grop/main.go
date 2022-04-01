package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/harveysanders/grop"
)

var ignoreCase bool

func init() {
	flag.BoolVar(&ignoreCase, "i", false, "case insensitive match")

	flag.Usage = usage
}

func main() {
	flag.Parse()
	var in io.Reader
	args := os.Args[1:]

	if len(args) == 0 {
		usage()
		os.Exit(1)
	}

	if err := grop.Run(args, os.Stdout, in, grop.Options{IgnoreCase: ignoreCase}); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func usage() {
	fmt.Fprintln(os.Stderr, "usage: grop [-i] [pattern] [file]")
}
