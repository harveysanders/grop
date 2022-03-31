package main

import (
	"bytes"
	"os"
	"testing"
)

func TestSearch(t *testing.T) {
	cases := []struct {
		term     string
		filepath string
		results  string
	}{
		{term: "power", filepath: "test_files/Kendrick-lamar-dna-lyrics.txt", results: `I got power, poison, pain, and joy inside my DNA
	The reason my power's here on earth
	Money and power, the mecca of marriages`},
	}

	for _, c := range cases {
		want := c.results
		// r, err := os.ReadFile(c.filepath)
		file, err := os.OpenFile(c.filepath, os.O_RDONLY, os.ModePerm)
		var buf bytes.Buffer

		if err != nil {
			t.Errorf("Unexpected error reading file: %q", err)
		}

		if err = Search(&buf, file, c.term); err != nil {
			t.Errorf("Unexpected Search() error: %q", err)
		}

		got := buf.String()
		if got != want {
			t.Errorf("Search Failed: %q.\nGot: %q\nwant %q", c.term, got, want)
		}
	}
}
