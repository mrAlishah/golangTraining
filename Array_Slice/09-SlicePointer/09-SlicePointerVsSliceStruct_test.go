package main

import "testing"

// go test -bench=. -benchmem -benchtime=10000x

/*************************************************************************************************/
/* Benchmark return Slice of Struct
/*************************************************************************************************/
func BenchmarkReturnSliceWithStructs(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ReturnSliceWithStructs(10000)
	}
}

func BenchmarkReturnSliceWithPointers(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ReturnSliceWithPointers(10000)
	}
}

func BenchmarkReturnSliceWithParamPointerSliceStruct(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var perPPS []Person
		ReturnSliceWithParamPointerSliceStruct(&perPPS, 10000)
	}
}

func BenchmarkReturnSliceWithParamPointersSlicePointer(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var perPPP []*Person
		ReturnSliceWithParamPointersSlicePointer(&perPPP, 10000)
	}
}

func BenchmarkReturnSliceWithParamPointerSliceStructMaked(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var perPPS []Person
		perPPS = make([]Person, 10000)
		ReturnSliceWithParamPointerSliceStructMaked(&perPPS, 10000)
	}
}

func BenchmarkReturnSliceWithParamPointersSlicePointerMaked(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var perPPP []*Person
		perPPP = make([]*Person, 10000)
		ReturnSliceWithParamPointersSlicePointerMaked(&perPPP, 10000)
	}
}

/*************************************************************************************************/
/* Benchmark return just one Struct
/*************************************************************************************************/

func BenchmarkReturnOneStructWithStructs(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ReturnOneStructWithStructs(b.N)
	}
}

func BenchmarkReturnOneStructWithPointers(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ReturnOneStructWithPointers(b.N)
	}
}

func BenchmarkReturnOneStructAsParamPointers(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var p Person
		ReturnOneStructAsParamPointers(&p)
	}
}

/*
goos: darwin
goarch: amd64
pkg: golangTraining/Array_Slice/09-SlicePointer
cpu: Intel(R) Core(TM) i7-8850H CPU @ 2.60GHz
BenchmarkReturnSliceWithStructs-12                                 10000              9570 ns/op           81920 B/op          1 allocs/op
BenchmarkReturnSliceWithPointers-12                                10000            146575 ns/op          161921 B/op      10001 allocs/op
BenchmarkReturnSliceWithParamPointerSliceStruct-12                 10000             14219 ns/op           81920 B/op          1 allocs/op
BenchmarkReturnSliceWithParamPointersSlicePointer-12               10000            152470 ns/op          161921 B/op      10001 allocs/op
BenchmarkReturnSliceWithParamPointerSliceStructMaked-12            10000             12144 ns/op           81920 B/op          1 allocs/op
BenchmarkReturnSliceWithParamPointersSlicePointerMaked-12          10000            140166 ns/op          161921 B/op      10001 allocs/op
BenchmarkReturnOneStructWithStructs-12                             10000                 0.2697 ns/op          0 B/op          0 allocs/op
BenchmarkReturnOneStructWithPointers-12                            10000                 0.2842 ns/op          0 B/op          0 allocs/op
BenchmarkReturnOneStructAsParamPointers-12                         10000                 0.2678 ns/op          0 B/op          0 allocs/op
PASS
ok      golangTraining/Array_Slice/09-SlicePointer      5.706s

*/
