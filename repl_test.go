package main

import (
	"testing"
)

func TestCleanInput(t *testing.T)  {
	cases := []struct {
		input  		string
		expected 	[]string
	}{
		{
			input: "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "One  TWO three",
			expected: []string{"one", "two", "three"},
		},
		{
			input: "",
			expected: []string{},
		},
		{
			input: "one,two,three",
			expected: []string{"one,two,three"},
		},
	}

	for _, c := range cases{
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Actual with len %d is not equeal to expected len %d", len(actual), len(c.expected))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Expected %s is not equeal %s", expectedWord, word)
			}
		}
	}
}
