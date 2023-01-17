
# TDD / Simple learning unit test
<br>

## Unit Test: What do you need to remember?
- Tests are located in the same directory as the code under test.
- When you write a test file name should end with `_test.go`. 
- To run just one package of a project: `go test packageName` .
- To run *all the tests* of a package: `go test ./...` .
    - Adding the `-v` flag increases verbosity
- `func TestFuncName(t *testing.T) {}`
    - should start with `Test` .
    - Generally one function/method = one test function.
- Do not hesitate to use subtests :
    - `t.Run("nominal case", func(t *testing.T) { ... }) `

## Coverage Test
When writing tests, it is often important to know how much of your actual code the tests cover. This is generally referred to as coverage.

- To run and consider covering unit test code use `-cover` : `go test ./... -cover` .
- Go saved this coverage data in the file coverage.out .Now you can present the results in a web browser.
    `go test ./PackageName -coverprofile=coverage.out`
    `go tool cover -html=coverage.out`



