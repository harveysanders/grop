package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/harveysanders/grop"
)

var color string
var ignoreCase bool

func init() {
	flag.StringVar(&color, "color", "never", "when to highlight matching patterns ( always | auto | never )")
	flag.BoolVar(&ignoreCase, "i", false, "case insensitive match")

	flag.Usage = usage
}

func main() {
	flag.Parse()
	var in io.Reader
	args := flag.Args()

	if len(args) == 0 {
		usage()
		os.Exit(1)
	}

	opts := grop.Options{IgnoreCase: ignoreCase, WhenHighlight: color}

	if err := grop.Run(args, os.Stdout, in, opts); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func usage() {
	fmt.Fprintln(os.Stderr, "usage: grop [-i] [--color=when] [pattern] [file]")
}
