package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "    hello world     ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "bulbasaur pikachu     ",
			expected: []string{"bulbasaur", "pikachu"},
		},
		{
			input:    "      I can't think of anymore test cases",
			expected: []string{"I", "can't", "think", "of", "anymore", "test", "cases"},
		},
	}

	for _, c := range cases {
		actual := CleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("The number of words doesn't match")
		}

		for i, word := range actual {
			expectedWord := c.expected[i]
			if expectedWord != word {
				t.Errorf("Words don't match expectations: %s != %s", expectedWord, word)
			}
		}
	}
}
