package test

import "testing"

func TestIsArraySorted(t *testing.T) {
	testcases := []struct {
		input    []int
		expected bool
	}{
		{[]int{1, 2, 3, 4, 5}, true},
		{[]int{5, 4, 3, 2, 1}, false},
		{[]int{1, 1, 1, 1, 1}, true},
		{[]int{1}, true},
		{[]int{}, true},
	}

	for _, tc := range testcases {
		result := isArraySorted(tc.input)
		if result != tc.expected {
			t.Errorf("Expected %v, but got %v\n", tc.expected, result)
		}
	}
}
