# Pointers vs. values

## What is a pointer
 look at these link:
 - [Link 1](https://golangbot.com/pointers/)
 - [Link 2](https://dave.cheney.net/2017/04/26/understand-go-pointers-in-less-than-800-words-or-your-money-back)
 - [Link 3:What is a reference variable?](https://dave.cheney.net/2017/04/29/there-is-no-pass-by-reference-in-go)

## Pointers vs. values in parameters and return values
first please watch this link:
[Understanding Allocations: the Stack and the Heap](https://www.youtube.com/watch?v=ZMZpH4yT7M0)

### Return just one Struct
```go
type myStruct struct {
	Age int
}
```
1. To Return a copy of the struct <br/>
**RESULT:** It could be ok for small structs (because the overhead is minimal)
```go
func myfunc(age int) myStruct {
	res := myStruct{Age: age}
	return res
}
```

2. A pointer to the struct value created within the function
**RESULT:** It could be ok for bigger struct used

```go
func myfunc(age int) *myStruct {
    res := &myStruct{Age: age}
	return res
}
```

3. it expects an existing struct to be passed in and overrides the value.<br/>
**RESULT:** if you want to be extremely memory efficient because you can easily reuse a single struct instance between calls.
```go
func myfunc(res *myStruct) {
	res.Age = 1
}
```

if you see io.Reader can get why Go have been used 
```go
type Read interface {
	Read (p []byte) (n int, err error)
}
```
instead of
```go
type Read interface {
	Read (n int) (p []byte, err error)
}
```

BenchMark:
```
BenchmarkReturnOneStructWithStructs-12                             10000                 0.2697 ns/op          0 B/op          0 allocs/op
BenchmarkReturnOneStructWithPointers-12                            10000                 0.2842 ns/op          0 B/op          0 allocs/op
BenchmarkReturnOneStructAsParamPointers-12                         10000                 0.2678 ns/op          0 B/op          0 allocs/op
```

### Return Slice of Struct
slices are always pointers, so returning a pointer to a slice isn't useful. However, should I return a slice of struct values, a slice of pointers to structs.
- [Link 1](https://stackoverflow.com/questions/23542989/pointers-vs-values-in-parameters-and-return-values)
- [Link 2](https://talk.gocasts.ir/t/struct-pointer/260)
```go
func myfunc() []myStruct {
    return []myStruct{ MyStruct{age: 1} }
}

func myfunc() []*myStruct {
    return []myStruct{ &myStruct{age: 1} }
}

// not good
func myfunc(s *[]myStruct) {
    *s = []myStruct{ myStruct{age: 1} }
}

// not good
func myfunc(s *[]*myStruct) {
    *s = []myStruct{ &myStruct{age: 1} }
}
```

```
BenchmarkReturnSliceWithStructs-12                                 10000              9570 ns/op           81920 B/op          1 allocs/op
BenchmarkReturnSliceWithPointers-12                                10000            146575 ns/op          161921 B/op      10001 allocs/op
BenchmarkReturnSliceWithParamPointerSliceStruct-12                 10000             14219 ns/op           81920 B/op          1 allocs/op
BenchmarkReturnSliceWithParamPointersSlicePointer-12               10000            152470 ns/op          161921 B/op      10001 allocs/op
```