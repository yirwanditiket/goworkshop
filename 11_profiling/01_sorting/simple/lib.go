package main

func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		// Track if any swap is made
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// Swap the elements
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		// If no two elements were swapped, the list is sorted
		if !swapped {
			break
		}
	}
}
