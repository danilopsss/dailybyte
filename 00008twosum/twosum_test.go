package main

import "testing"

func TestTwoSumNested(t *testing.T) {
	t.Run("K=10", func(t *testing.T) {
		testSlice := TwoSum{slice: []int{1, 3, 8, 2}, k: 10}
		expected := true
		if actual := testSlice.hasSum(); actual != expected {
			t.Errorf("Expected %t, got %t", expected, actual)
		}
	})
	t.Run("K=8", func(t *testing.T) {
		testSlice := TwoSum{slice: []int{3, 9, 13, 7}, k: 8}
		expected := false
		if actual := testSlice.hasSum(); actual != expected {
			t.Errorf("Expected %t, got %t", expected, actual)
		}
	})

	t.Run("K=4", func(t *testing.T) {
		testSlice := TwoSum{slice: []int{4, 2, 6, 5, 2}, k: 4}
		expected := true
		if actual := testSlice.hasSum(); actual != expected {
			t.Errorf("Expected %t, got %t", expected, actual)
		}
	})
}
func TestTwoSumNotNested(t *testing.T) {
	t.Run("K=10", func(t *testing.T) {
		testSlice := TwoSum{slice: []int{1, 3, 8, 2}, k: 10}
		expected := true
		if actual := testSlice.hasSumWithNoNestedLoop(); actual != expected {
			t.Errorf("Expected %t, got %t", expected, actual)
		}
	})
	t.Run("K=8", func(t *testing.T) {
		testSlice := TwoSum{slice: []int{3, 9, 13, 7}, k: 8}
		expected := false
		if actual := testSlice.hasSumWithNoNestedLoop(); actual != expected {
			t.Errorf("Expected %t, got %t", expected, actual)
		}
	})

	t.Run("K=4", func(t *testing.T) {
		testSlice := TwoSum{slice: []int{4, 2, 6, 5, 2}, k: 4}
		expected := true
		if actual := testSlice.hasSumWithNoNestedLoop(); actual != expected {
			t.Errorf("Expected %t, got %t", expected, actual)
		}
	})

}
