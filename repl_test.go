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
		{
			input:    "   ",
			expected: []string{},
		},
		{
			input:    "solo",
			expected: []string{"solo"},
		},
		{
			input:    "  spaces",
			expected: []string{"spaces"},
		},
		{
			input:    "CaPITaLizaTiON ISN'T necessary",
			expected: []string{"capitalization", "isn't", "necessary"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Expected length: %d, Got: %d", len(c.expected), len(actual))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Expected word: %s, Got: %s", expectedWord, word)
			}
		}
	}
}
