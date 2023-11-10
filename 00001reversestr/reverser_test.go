package reversestr

import "testing"

func TestStringReverser(t *testing.T) {
	t.Run("StringReverser", func(t *testing.T) {
		stringOne := "Cat"
		stringOneReversed := StringReverser(stringOne)
		if stringOneReversed != "taC" {
			t.Errorf("Expected %s, got %s", "taC", stringOneReversed)
		}

		stringTwo := "The Daily Byte"
		stringTwoReversed := StringReverser(stringTwo)
		if stringTwoReversed != "etyB yliaD ehT" {
			t.Errorf("Expected %s, got %s", "etyB yliaD ehT", stringTwoReversed)
		}

		stringThree := "civic"
		stringThreeReversed := StringReverser(stringThree)
		if stringThreeReversed != "civic" {
			t.Errorf("Expected %s, got %s", "civic", stringThreeReversed)
		}

	})
}