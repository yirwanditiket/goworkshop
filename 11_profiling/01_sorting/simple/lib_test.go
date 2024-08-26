package main

import (
	"testing"
)

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cp := append([]int{}, toSort...)
		BubbleSort(cp)
	}
}
