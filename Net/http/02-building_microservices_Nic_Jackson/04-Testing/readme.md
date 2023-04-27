# Introduction of testing
When you try to define what testing is, you will come up with a multitude of answers, and many of us will not understand the full benefits of testing until we've been burnt by buggy software or we have tried to change a complex code base which has no tests.
>"The art of a good night's sleep is knowing you will not get woken by a support call and the piece of mind from being able to
confidently change your software in an always moving market."

OK, so I am trying to be funny, but the concept is correct:
* Nobody enjoys debugging poorly written code, and indeed,
* Nobody enjoys the stress caused when a system fails.
* Starting out with a mantra of quality first can alleviate many of these problems.

## The testing pyramid
<img src="Img%2F01-Testing.png" width="600">
<img src="Img%2F02-Testing.jpg" width="600">
<img src="Img%2F09-Integerety.png" width="600"><br/>

* Unit test: Implement to detect error on your design. your cheapest (fastest) tests to run, which will be your unit tests, go at the bottom of the pyramid;
* Service: where you define a flow service in your application.
* UI: where user will use your applicaiton, and detect error, which are the costliest element.

### Automated testing
In the early days of automated testing, all the testing was completed at the top of the pyramid. While this did work from a quality perspective, it meant the process of debugging the area at fault would be incredibly complicated and time-consuming.

#### Problem
* If you were lucky, there might be a complete failure which could be tracked down to a stack trace.
* If you were unlucky, then the problem would be behavioral; and even if you knew the system inside out, it would involve plowing through thousands of lines of code and manually repeating the action to reproduce the failure.

## Outside-in development
When writing tests, I like to follow a process called outside-in development.<br/>
With outside-in development, you start by writing your tests almost at the top of the pyramid, determine what the functionality is going to be for the story you are working on, and then write some failing test for this story.<br/>
Then you work on implementing the unit tests and code which starts to get the various steps in the behavioral tests (BDD) to pass.<br/>
>Feature: As a user when I call the search endpoint, I would like to receive a list of kittens

The feature is the story which, in an agile environment is owned by the product owner. The feature is then broken down into scenarios which explain in greater detail the qualities that the code must have to be acceptable.
> Scenario: Invalid query
Given I have no search criteria
When I call the search endpoint
Then I should receive a bad request message

<img src="Img%2F03-Testing.jpg" width="600">
<img src="Img%2F05-BDD.png" width="600">
<img src="Img%2F06-BDD-TDD.png" width="600">
<img src="Img%2F07-Behaviour-Driven Development - Cucumber.png" width="600">
<img src="Img%2F08-BDD.jpeg" width="600">
<img src="Img%2F09-BDD.jpeg" width="600"> 

### Unit-test
Our unit tests go right down to the bottom of the pyramid.

#### Law
* **First law:** You may not write production code until you have written a failing unit test
* **Second law:**  You may not write more of a unit test than is sufficient to fail, and not compiling is failing
* **Third law:**  You may not write more production code than is sufficient to pass the currently failing test

### Note
One of the most effective ways to test a microservice in Go is not to fall into the trap of trying to execute all the tests through the HTTP interface.  
**Follow steps:**
1. **Step 1:** Create a pattern for test program  
   - We need develop **a pattern that avoids** creating a physical web server for testing our handlers, the code to create this kind of test is slow to run and incredibly tedious to write.
2. **Step 2:** Implement Unit test
   - What need to be doing is to test our handlers and the code within them as **Unit test**.
   - These tests will run far quicker than testing through the web server.
3. **Step 3:** Get coverage.
   - And if we think about coverage, we will be able to test the writing of the handlers in the **Cucumber** tests that execute a request to the running server which overall gives us 100% coverage of our code.

Production code: Production means anything that you need to work reliably, and consistently.  
Refer: [here](https://stackoverflow.com/questions/490289/what-exactly-defines-production) <br/>

### Naming for test case
The name of the test must have a particular name beginning with Test and then immediately following this an uppercase
character or number.  
For a example:
- Do not: TestmyHandler
- Should: Test1Handler
- Should: TestMyHandler
- Recommend: Test1MyHandler
- Recommend: TestSearchHandlerReturnsBadRequestWhenNoSearchCriteriaIsSent

  <img src="Img%2F10-AAA.jpeg" width="600">
  <img src="Img%2F11-AAA.jpeg" width="600">

**Mock**
idiomatically: Mock tests are exams that are prepared on guidelines same as real exams. These follow the same pattern, set of questions, difficulty level, and also time limits. And just because these mimic the real exam, these are called “Mock” tests!<br/><br/> 
Developing Mock testing: Mock Testing provides you the ability to isolate and test your code without any interference of the dependencies and other variables like network issues and traffic fluctuations. In simple words, in mock testing, we replace the dependent objects with mock objects.<br/>
* **stub** is replacement for some dependency in your code that will be used during test execution. It is typically built for one particular test and unlikely can be reused for another because it has hardcoded expectations and assumptions.
* **mock** takes stubs to next level. It adds means for configuration, so you can set up different expectations for different tests. That makes mocks more complicated, but reusable for different tests.
What's the difference between stub and mock in Go unit testing? [here](https://stackoverflow.com/questions/53360256/whats-the-difference-between-stub-and-mock-in-go-unit-testing)

### httptest.NewRequest
`http.Request` that can be created using `httptest.NewRequest` exported function.NewRequest returns a new incoming server Request, suitable for passing to an http.Handler for testing.
```go
func NewRequest(method, target string, body io.Reader) *http.Request
```
We can pass parameters to the `method` and the `target`, which is either the path or an absolute URL.<br/>
Finally, we can give it an io.Reader file which will correspond to the body of the request; if we do not pass a nil value then Request.ContentLength is set.
```go
	body, _ := json.Marshal(data)
	req := httptest.NewRequest("POST", "/search", bytes.NewReader(body))
```
### httptest.NewRecorder
`http.ResponseWriter` that can be created by using `httptest.NewRecorder` type which returns a `httptest.ResponseRecorder`.ResponseRecorder is an implementation of `http.ResponseWriter` that records its mutations for later inspection in tests.<br/>
The httptest.ResponseRecorder is an implementation of http.ResponseWriter and can be used to be passed into our server handler, record all the data that the handler will write to the response and return the data written afterwards.<br/>
Refer to Example: [httptest](https://speedscale.com/blog/testing-golang-with-httptest/)
```go
response := httptest.NewRecorder()
```
**httptest.ResponseRecorder:**
```go
type ResponseRecorder struct {
    Code int // the HTTP response code from WriteHeader
    HeaderMap http.Header // the HTTP response headers
    Body *bytes.Buffer // if non-nil, the bytes.Buffer to append written data to
    Flushed bool
    // contains filtered or unexported fields
}
```

### Run Unit Test
```go
 go test -v -race ./...
```
* The `-v` flag will print the output in a verbose style, and it will also print all the text written to the output by the application, even if the test succeeds.
* The `-race` flag enables Go's race detector which holds discover bugs with concurrency problems. A data race occurs when two Go routines access the same variable concurrently, and at least one of the accesses is a write. The race flag adds a small overhead to your test run, so I recommend you add it to all executions.
* Using `./...` as our final parameter allows us to run all our tests in the current folder as well as the child folders, it saves us from manually having to construct a list of packages or files to test.

## Behavirol Driven Development
### What is BDD?
* Behavioral Driven Development (BDD) and is a technique often executed by an application framework called Cucumber.
* It was developed by Dan North and was designed to create a common ground between developers and product owners.
