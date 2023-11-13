package main

import "testing"

func TestBinarySum(t *testing.T) {
	t.Run("Test binary sum", func(t *testing.T) {
		testString := IncomingBinaries{"100", "1"}
		expected := "101"
		if actual := testString.BinarySum(); actual != expected {
			t.Errorf("Expected %s, got %s", expected, actual)
		}

		testString = IncomingBinaries{"11", "1"}
		expected = "100"
		if actual := testString.BinarySum(); actual != expected {
			t.Errorf("Expected %s, got %s", expected, actual)
		}

		testString = IncomingBinaries{"1", "0"}
		expected = "1"
		if actual := testString.BinarySum(); actual != expected {
			t.Errorf("Expected %s, got %s", expected, actual)
		}
		// Extra one for good measure
		testString = IncomingBinaries{"1101", "100"}
		expected = "10001"
		if actual := testString.BinarySum(); actual != expected {
			t.Errorf("Expected %s, got %s", expected, actual)
		}
	})
}

