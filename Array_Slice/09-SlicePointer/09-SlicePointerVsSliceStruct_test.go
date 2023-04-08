package main

import "testing"

// go test -bench=. -benchmem -benchtime=10000x

/*************************************************************************************************/
/* Benchmark return Slice of Struct
/*************************************************************************************************/

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

/*************************************************************************************************/
/* Benchmark return just one Struct
/*************************************************************************************************/

func BenchmarkReturnOneStructWithPointers(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ReturnOneStructWithPointers(b.N)
	}
}

func BenchmarkReturnOneStructWithStructs(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ReturnOneStructWithStructs(b.N)
	}
}

/*
goos: darwin
goarch: amd64
pkg: golangTraining/Array_Slice/09-SlicePointer
cpu: Intel(R) Core(TM) i7-8850H CPU @ 2.60GHz
BenchmarkReturnSliceWithPointers-12                10000            174658 ns/op          161922 B/op      10001 allocs/op
BenchmarkReturnSliceWithStructs-12                 10000             12370 ns/op           81920 B/op          1 allocs/op
BenchmarkReturnOneStructeWithPointers-12           10000                 0.3716 ns/op          0 B/op          0 allocs/op
BenchmarkReturnOneStructWithStructs-12             10000                 0.3726 ns/op          0 B/op          0 allocs/op
PASS
ok      golangTraining/Array_Slice/09-SlicePointer      2.762s
*/
