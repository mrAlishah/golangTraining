package main

import "testing"

// go test -bench=. -benchmem -benchtime=10000x

func BenchmarkReturnSliceWithPointers(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ReturnSliceWithPointers(10000)
	}
}

func BenchmarkReturnSliceWithStructs(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ReturnSliceWithStructs(10000)
	}
}
