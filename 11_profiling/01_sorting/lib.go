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

// heapSort sorts an array using heap sort algorithm
func HeapSort(arr []int) {
	n := len(arr)

	// Build a max-heap
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// Extract elements from the heap one by one
	for i := n - 1; i >= 0; i-- {
		// Move current root to the end
		arr[0], arr[i] = arr[i], arr[0]

		// Call heapify on the reduced heap
		heapify(arr, i, 0)
	}
}

// quickSort recursively sorts the slice using the Quicksort algorithm
func QuickSort(arr []int, low, high int) {
	if low < high {
		pi := partition(arr, low, high) // Partition the array and get the pivot index

		// Recursively sort elements before and after partition
		QuickSort(arr, low, pi-1)
		QuickSort(arr, pi+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high] // Choose the rightmost element as the pivot
	i := low - 1       // Pointer for the greater element

	for j := low; j < high; j++ {
		// If the current element is smaller than the pivot
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i] // Swap elements
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1] // Swap the pivot element with the element at i+1
	return i + 1                              // Return the partitioning index
}

func heapify(arr []int, n, i int) {
	largest := i     // Initialize largest as root
	left := 2*i + 1  // Left child
	right := 2*i + 2 // Right child

	// Check if left child exists and is greater than root
	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	// Check if right child exists and is greater than largest so far
	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	// If largest is not root
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i] // Swap

		// Recursively heapify the affected subtree
		heapify(arr, n, largest)
	}
}
