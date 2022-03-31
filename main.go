package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var in io.Reader

	if err := run(os.Args[1:], os.Stdout, in); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
