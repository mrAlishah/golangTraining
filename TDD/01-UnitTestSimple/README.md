
# TDD / Simple learning unit test
<br>

## Unit Test: What do you need to remember?
- Tests are located in the same directory as the code under test.
- When you write a test file name should end with `_test.go`. 
- To run just one package of a project: `go test packageName` .
- To run *all the tests* of a package: `go test ./...` .
    - Adding the `-v` flag increases verbosity.
    the test will print out the names of all the executed test functions and the time spent for their execution. Additionally, the test displays the output of printing to the error log, for example, when you use t.Log() or t.Logf()

- `func TestFuncName(t *testing.T) {}`
    - should start with `Test` .
    - Generally one function/method = one test function.
- Do not hesitate to use subtests :
    - `t.Run("nominal case", func(t *testing.T) { ... }) `
-  [A deep dive into unit test](https://blog.logrocket.com/a-deep-dive-into-unit-testing-in-go/)   

## Coverage Test
When writing tests, it is often important to know how much of your actual code the tests cover. This is generally referred to as coverage.

- To run and consider covering unit test code use `-cover` : `go test ./... -cover` .
- Go saved this coverage data in the file coverage.out .Now you can present the results in a web browser or file.
    `go test ./PackageName -coverprofile=coverage.out -covermode=count` </br>
    `go tool cover -func=coverage.out` </br>
    `go tool cover -html=coverage.out`

## Benchmark Test
A benchmark is a type of function that executes a code segment multiple times and compares each output against a standard, assessing the code’s overall performance level.
- `func BenchmarkFuncName(b *testing.B) {}`
    - should start with `Benchmark` .
- To run all benchmarks, use `go test -bench=.` or `go test ./packageName -bench=.`
  The argument to -bench is a regular expression that specifies which benchmarks should be run 
- To verify that the benchmark produces a consistent result, you can run it multiple times by passing a number to the `-count` flag = -count 5
- To avoid executing any test functions in the test files, pass a regular expression to the `-run` flag = `go test -bench=. -count 5 -run=^#`
-To include memory allocation statistics in the benchmark output, add the `-benchmem` flag = `go test -bench=. -count 5 -run=^# -benchmem`.  
--  [A deep dive into benchmark test](https://blog.logrocket.com/benchmarking-golang-improve-function-performance/)  

## Example Test
this is really regular test. 
- `func ExampleFuncName(...) {}`
    - should start with `Example` .
- you have to check output by comment: `//Output:`
- the expected result test is below by comment 
- `go test  -run ^ExampleFuncName$ ./addpack`
