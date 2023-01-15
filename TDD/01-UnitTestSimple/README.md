
# TDD / Simple learning unit test
<br>

## Unit Test: What do you need to remember?
- Tests are located in the same directory as thecode under test.
- When you write a test file name should end with `_test.go`. 
- To run just one package of a project: `go test packageName` .
- To run *all the tests* of a package: `go test ./...` .
- `func TestFuncName(t *testing.T) {}`
    - should start with `Test` .
    - Generally one function/method = one test function.
- Do not hesitate to use subtests :
    - `t.Run("nominal case", func(t *testing.T) { ... }) `

## Cover Test
- To run and consider covering unit test code use `-cover` : `go test ./... -cover` .

