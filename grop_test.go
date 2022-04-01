package grop

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestSearch(t *testing.T) {
	cases := []struct {
		label    string
		term     string
		filepath string
		results  string
	}{
		{label: "Handle empty term", term: "", filepath: "test_files/A_Mad_Tea_Party.txt", results: ""},
		{label: "Return lines matching term", term: "thing", filepath: "test_files/A_Mad_Tea_Party.txt", results: strings.Join([]string{`Alice looked all round the table, but there was nothing on it but tea. "I don't see any wine," she remarked.`,
			`"I do," Alice hastily replied; "at least -- at least I mean what I say -- that's the same thing, you know."`,
			`"Not the same thing a bit!" said the Hatter. "Why, you might just as well say that "I see what I eat" is the same thing as "I eat what I see!"`,
			`"You might just as well say," added the March Hare, "that "I like what I get" is the same thing as "I get what I like"!"`,
			`"You might just as well say," added the Dormouse, which seemed to be talking in its sleep, "that "I breathe when I sleep" is the same thing as "I sleep when I breathe"!"`,
			`"It is the same thing with you," said the Hatter, and here the conversation dropped, and the party sat silent for a minute, while Alice thought over all she could remember about ravens and writing-desks, which wasn't much.`,
		}, "\n") + "\n"},
	}

	for _, c := range cases {
		var buf bytes.Buffer
		want := c.results

		file, err := os.Open(c.filepath)
		if err != nil {
			t.Errorf("Unexpected error reading file: %q", err)
		}
		defer file.Close()

		if err = Search(&buf, file, c.term, Options{}); err != nil {
			t.Errorf("Unexpected Search() error: %q", err)
		}

		got := buf.String()
		if got != want {
			t.Errorf("Search Failed: %q.\nGot: %q\nwant %q", c.label, got, want)
		}
	}
}

func TestSearchWithOpts(t *testing.T) {
	var buf bytes.Buffer
	term := "hat"
	options := Options{
		IgnoreCase: true,
	}
	inputLines := []string{
		"This and that",
		"Look at my hat!",
		"this one, not THAT!",
		"end",
	}

	searchDoc := strings.Join(inputLines, "\n")

	r := strings.NewReader(searchDoc)

	err := Search(&buf, r, term, options)
	if err != nil {
		t.Errorf("Unexpected Error %v", err)
	}

	want := strings.Join(inputLines[:3], "\n") + "\n"
	got := buf.String()

	if got != want {
		t.Errorf("Case insensitive Search() failed.\nGot: %q\nWant: %q", got, want)
	}

}
