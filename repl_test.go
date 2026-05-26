package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {

	cases := map[string]struct {
		input    string
		expected []string
	}{
		"empty input": {
			input:    "",
			expected: []string{},
		},
		"only spaces": {
			input:    "    ",
			expected: []string{},
		},
		"spaces on beginning and end of the string": {
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		"common input": {
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			actual := cleanInput(c.input)

			if len(actual) != len(c.expected) {
				t.Fatalf("expected: %v, got: %v", actual, c.expected)
			}

			for i := range actual {
				word := actual[i]
				expectedWord := c.expected[i]
				fmt.Println(word, expectedWord)
				if word != expectedWord {
					t.Fatalf("expected: %s, got: %s", word, expectedWord)
				}
			}
		})
	}

}
