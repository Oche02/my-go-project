package piscine

func Abort(a, b, c, d, e int) int {
	arr := []int{a, b, c, d, e}

	// Simple bubble sort for 5 elements
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr[2] // The median (middle) element after sorting
}
