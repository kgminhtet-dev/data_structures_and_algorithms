package day3

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {

	testcases := []struct {
		input    []int
		expected []int
	}{
		{[]int{45, 50, 30, 15, 7, 11}, []int{7, 11, 15, 30, 45, 50}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 1, 1, 1, 1}, []int{1, 1, 1, 1, 1}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
	}

	for _, tc := range testcases {
		result := bubbleSort(tc.input)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Expected %v, but got %v\n", tc.expected, result)
		}
	}
}
