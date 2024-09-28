package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  GoLang is Awesome  ",
			expected: []string{"golang", "is", "awesome"},
		},
		{
			input:    "TESTING, one TWO three",
			expected: []string{"testing,", "one", "two", "three"},
		},
		{
			input:    "  Mixed CASE and  SPACES ",
			expected: []string{"mixed", "case", "and", "spaces"},
		},
		{
			input:    "123 456 789",
			expected: []string{"123", "456", "789"},
		},
	}

	for _, cs := range cases {
		fResult := cleanInput(cs.input) // fResult meaning function result
		if len(fResult) != len(cs.expected) {
			t.Errorf("This lengths are not equal: %v vs %v",
				len(fResult), len(cs.expected))
			continue
		}
		for i := range fResult {
			if fResult[i] != cs.expected[i] {
				t.Errorf("%v does not equal %v", fResult[i], cs.expected[i])
			}
		}
	}

}
