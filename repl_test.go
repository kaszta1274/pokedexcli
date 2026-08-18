package main

import "testing"

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
			input:    "a a",
			expected: []string{"a", "a"},
		},
		{
			input:    " Hello     WORLD",
			expected: []string{"hello", "world"},
		},
		{
			input:    "         ",
			expected: []string{},
		},
		// add more cases here
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("len(actual) = %v, want %v", len(actual), len(c.expected))
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("word = %v, expected %v", word, expectedWord)
				t.Fail()
			}
		}
	}
}
