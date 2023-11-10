package main

import "testing"

func TestValidPalindrome(t *testing.T) {
	t.Run("ValidPalindrome", func(t *testing.T) {
		stringOne := "level"
		stringOneValidPalindrome := ValidPalindrome(stringOne)
		if stringOneValidPalindrome != true {
			t.Errorf("Expected %t, got %t", true, stringOneValidPalindrome)
		}

		stringTwo := "algorithm"
		stringTwoValidPalindrome := ValidPalindrome(stringTwo)
		if stringTwoValidPalindrome != false {
			t.Errorf("Expected %t, got %t", false, stringTwoValidPalindrome)
		}

		stringThree := "A man, a plan, a canal: Panama."
		stringThreeValidPalindrome := ValidPalindrome(stringThree)
		if stringThreeValidPalindrome != true {
			t.Errorf("Expected %t, got %t", true, stringThreeValidPalindrome)
		}
	})
}
