package main

import "testing"

func TestValidPalindrome(t *testing.T) {
	t.Run("VacuumCleaningRoute", func(t *testing.T) {
		stringOne := "LR"
		stringOneValidPalindrome := StoppedAtAStartingPos(stringOne)
		if stringOneValidPalindrome != true {
			t.Errorf("Expected %t, got %t", true, stringOneValidPalindrome)
		}

		stringTwo := "URURD"
		stringTwoValidPalindrome := StoppedAtAStartingPos(stringTwo)
		if stringTwoValidPalindrome != false {
			t.Errorf("Expected %t, got %t", false, stringTwoValidPalindrome)
		}

		stringThree := "RUULLDRD"
		stringThreeValidPalindrome := StoppedAtAStartingPos(stringThree)
		if stringThreeValidPalindrome != true {
			t.Errorf("Expected %t, got %t", true, stringThreeValidPalindrome)
		}
	})
}
