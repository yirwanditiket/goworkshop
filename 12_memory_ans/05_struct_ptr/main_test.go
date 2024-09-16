package main

import (
	"testing"
)

func BenchmarkStructCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewStructCopy()
		// s.AddTwo(i)
	}
}

func BenchmarkStructPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewStructPointer()
		// s.AddTwo(i)
	}
}
