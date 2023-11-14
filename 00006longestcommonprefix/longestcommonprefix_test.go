package main

import "testing"

func TestLongestPrefix(t *testing.T) {
	t.Run("Longest Prefix", func(t *testing.T) {
		testString := stringEval{wordList: []string{"colorado", "color", "cold"}}
		expected := "col"
		if actual := testString.LongestPrefix(); actual != expected {
			t.Errorf("Expected %s, got %s", expected, actual)
		}
		testString = stringEval{wordList: []string{"1", "b", "c"}}
		expected = ""
		if actual := testString.LongestPrefix(); actual != expected {
			t.Errorf("Expected %s, got %s", expected, actual)
		}
		testString = stringEval{wordList: []string{"spot", "spotty", "spotted"}}
		expected = "spot"
		if actual := testString.LongestPrefix(); actual != expected {
			t.Errorf("Expected %s, got %s", expected, actual)
		}
	})
}

