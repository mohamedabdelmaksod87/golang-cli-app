package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	type TestData struct {
		input    string
		expected []string
	}

	cases := []TestData{
		{
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "HELLO world",
			expected: []string{"hello", "world"},
		},
	}

	for _, cs := range cases {
		actual := cleanInput(cs.input)
		if len(actual) != len(cs.expected) {
			t.Errorf("The lengths are not equal: %v vs %v", len(actual), len(cs.expected))
			continue
		}
		for i := range actual {
			actualWord := actual[i]
			expectedWord := cs.expected[i]
			if actualWord != expectedWord {
				t.Errorf("%v doesn't equal %v", actualWord, expectedWord)
			}
		}
	}
}
