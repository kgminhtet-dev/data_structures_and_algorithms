package test

import (
	"github.com/kgminhtet-dev/data_structures_and_algorithms/day3"
	"testing"
)

func TestSorting(t *testing.T) {
	size := 20
	testcases := make([][]int, 5)

	for i := range testcases {
		testcases[i] = generateRandomArray(size)
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
