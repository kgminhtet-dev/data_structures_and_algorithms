package test

import (
	"github.com/kgminhtet-dev/data_structures_and_algorithms/day3"
	"testing"
)

func TestSorting(t *testing.T) {
	testcases := [][]int{
		[]int{1, 2, 3, 4, 5},
		[]int{5, 4, 3, 2, 1},
		[]int{1, 1, 1, 1, 1},
		[]int{1},
		[]int{},
	}

	t.Run("Bubble Sort", func(t *testing.T) {
		for _, tc := range testcases {
			result := day3.BubbleSort(tc)
			if !isArraySorted(result) {
				t.Errorf("Result should be sorted %v", result)
			}
		}
	})
}
