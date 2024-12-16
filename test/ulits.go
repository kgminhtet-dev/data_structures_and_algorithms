package test

func isArraySorted(arr []int) bool {
	if len(arr) == 0 || len(arr) == 1 {
		return true
	}

	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			return false
		}
	}
	return true
}
