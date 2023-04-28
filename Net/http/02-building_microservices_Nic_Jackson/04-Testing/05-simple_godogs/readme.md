# Godog: Cucumber for Golang
look at:
[link1](https://github.com/cucumber/godog)
[sample](https://github.com/cucumber/godog/tree/main/_examples/godogs)
[youtube time 18](https://www.youtube.com/watch?v=ucLN1T0H5-A)
[link2](https://medium.com/propertyfinder-engineering/golang-api-testing-with-godog-2de8944d2511)

## Install
1.
```go
go install github.com/cucumber/godog/cmd/godog@v0.12.0
```
Adding `@v0.12.0` will install `v0.12.0` specifically instead of master.<br/>

With go version prior to 1.17, use go get github.com/cucumber/godog/cmd/godog@v0.12.0. Running within the $GOPATH, you would also need to set GO111MODULE=on, like this:
```go
GO111MODULE=on go get github.com/cucumber/godog/cmd/godog@v0.12.0
```
2. set $PATH=$GOPATH/bin
3. `godog version `

## Implement Godogs
1. **Step 1** - Setup a go module
Initiate the go module inside the godogs directory by running `go mod init godogs`
2. **Step 2** - Create gherkin feature 
First of all, we describe our feature in plain text:
```
Feature: eat godogs
    In order to be happy
    As a hungry gopher
    I need to be able to eat godogs

    Scenario: Eat 5 out of 12
        Given there are 12 godogs
        When I eat 5
        Then there should be 7 remaining
```
add the text above into `features/godogs.feature`

3. **Step 3** - Create godog step definitions
    - Create and copy the step definitions "godogsPure_test.go.pure" into a new file by running `godogs_test.go`
    - Run `go test` in the godogs directory
4. **Step 4** - Create the main program to test
 create your logic code.like `godogs.go`
5. **Step 5** - Add some test logic to the step definitions
 Now lets implement our step definitions to test our feature requirements.
6. **Step 6** - run test
```go
go test -v godogs_test.go
```