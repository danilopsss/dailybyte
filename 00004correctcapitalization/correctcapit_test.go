package main

import "testing"

func TestCorrectCapitalization(t *testing.T) {
	t.Run("All uppercases must be true", func(t *testing.T) {
		stringText := "USA"
		if !HasCorrectCapitalization(stringText) {
			t.Errorf("got %t want %t", false, true)
		}
		stringText = "ABCDEFG"
		if !HasCorrectCapitalization(stringText) {
			t.Errorf("got %t want %t", false, true)
		}
		stringText = "ANOTHERRANDOMSTRING"
		if !HasCorrectCapitalization(stringText) {
			t.Errorf("got %t want %t", false, true)
		}
	})

	t.Run("First uppercase letter must be true", func(t *testing.T) {
		stringText := "Calvin"
		if !HasCorrectCapitalization(stringText) {
			t.Errorf("got %t want %t", false, true)
		}
		stringText = "Anothertest"
		if !HasCorrectCapitalization(stringText) {
			t.Errorf("got %t want %t", false, true)
		}
		stringText = "Thelasoneipromise"
		if !HasCorrectCapitalization(stringText) {
			t.Errorf("got %t want %t", false, true)
		}
	})

	t.Run("Some lost uppercase in the middle false", func(t *testing.T) {
		stringText := "compUter"
		if !HasCorrectCapitalization(stringText) {
			t.Errorf("got %t want %t", false, true)
		}
		stringText = "noNsensestring"
		if !HasCorrectCapitalization(stringText) {
			t.Errorf("got %t want %t", false, true)
		}
		stringText = "anothernonSenseString"
		if !HasCorrectCapitalization(stringText) {
			t.Errorf("got %t want %t", false, true)
		}
	})

	t.Run("All lowercases are good", func(t *testing.T) {
		stringText := "coding"
		if !HasCorrectCapitalization(stringText) {
			t.Errorf("got %t want %t", false, true)
		}
		stringText = "anotherlowercase"
		if !HasCorrectCapitalization(stringText) {
			t.Errorf("got %t want %t", false, true)
		}
		stringText = "ikindalikingthischallenge"
		if !HasCorrectCapitalization(stringText) {
			t.Errorf("got %t want %t", false, true)
		}
	})
	
}

