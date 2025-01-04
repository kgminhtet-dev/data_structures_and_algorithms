package test

import (
	"testing"

	"github.com/kgminhtet-dev/data_structures_and_algorithms/day3"
	"github.com/kgminhtet-dev/data_structures_and_algorithms/day4"
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

	t.Run("Selection Sort", func(t *testing.T) {
		for _, tc := range testcases {
			result := day3.SelectionSort(tc)
			if !isArraySorted(result) {
				t.Errorf("Result should be sorted %v", result)
			}
		}
	})

	t.Run("Insertion Sort", func(t *testing.T) {
		for _, tc := range testcases {
			result := day4.InsertionSort(tc)
			if !isArraySorted(result) {
				t.Errorf("Result should be sorted %v", result)
			}
		}
	})
}
