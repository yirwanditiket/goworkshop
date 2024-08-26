package main

import (
	"sort"
	"testing"
)

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cp := append([]int{}, toSort...)
		BubbleSort(cp)
	}
}

func BenchmarkGoSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cp := append([]int{}, toSort...)
		sort.Ints(cp)
	}
}

func BenchmarkHeapSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cp := append([]int{}, toSort...)
		HeapSort(cp)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cp := append([]int{}, toSort...)
		QuickSort(cp, 0, len(cp)-1)
	}
}
