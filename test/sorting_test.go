package test

import (
	"github.com/kgminhtet-dev/data_structures_and_algorithms/day3"
	"reflect"
	"testing"
)

func TestSorting(t *testing.T) {
	testcases := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 1, 1, 1, 1}, []int{1, 1, 1, 1, 1}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
	}

	t.Run("Bubble Sort", func(t *testing.T) {
		for _, tc := range testcases {
			result := day3.BubbleSort(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v\n", tc.expected, result)
			}
		}
	})
}
