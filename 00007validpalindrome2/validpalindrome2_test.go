package main

import "testing"

func TestPalindrome2(t *testing.T) {
	t.Run("Palidrome with removal", func(t *testing.T) {
		testString := String{word: "abcba"}
		expected := true
		if actual := testString.isPalindrome(); actual != expected {
			t.Errorf("Expected %t, got %t", expected, actual)
		}
		testString = String{word: "foobof"}
		expected = true
		if actual := testString.isPalindrome(); actual != expected {
			t.Errorf("Expected %t, got %t", expected, actual)
		}
		testString = String{word: "abccab"}
		expected = false
		if actual := testString.isPalindrome(); actual != expected {
			t.Errorf("Expected %t, got %t", expected, actual)
		}
		testString = String{word: "abbfa"}
		expected = true
		if actual := testString.isPalindrome(); actual != expected {
			t.Errorf("Expected %t, got %t", expected, actual)
		}
	})
}
