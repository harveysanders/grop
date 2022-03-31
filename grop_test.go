package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestSearch(t *testing.T) {
	cases := []struct {
		term     string
		filepath string
		results  string
	}{
		{term: "", filepath: "test_files/A_Mad_Tea_Party.txt", results: ""},
		{term: "thing", filepath: "test_files/A_Mad_Tea_Party.txt", results: strings.Join([]string{`Alice looked all round the table, but there was nothing on it but tea. "I don't see any wine," she remarked.`,
			`"I do," Alice hastily replied; "at least -- at least I mean what I say -- that's the same thing, you know."`,
			`"Not the same thing a bit!" said the Hatter. "Why, you might just as well say that "I see what I eat" is the same thing as "I eat what I see!"`,
			`"You might just as well say," added the March Hare, "that "I like what I get" is the same thing as "I get what I like"!"`,
			`"You might just as well say," added the Dormouse, which seemed to be talking in its sleep, "that "I breathe when I sleep" is the same thing as "I sleep when I breathe"!"`,
			`"It is the same thing with you," said the Hatter, and here the conversation dropped, and the party sat silent for a minute, while Alice thought over all she could remember about ravens and writing-desks, which wasn't much.`,
		}, "\n") + "\n"},
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
