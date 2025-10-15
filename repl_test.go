package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		//make a list of cases
		//add test cases for them one by one
		//write functionale to satisfy tests one by one
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		
		if len(actual) != len(c.input) {
			t.Errorf("resulting slice length != expected slice length")
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("expected %s, got %s", expectedWord, word)
			}
		}
	}
}
