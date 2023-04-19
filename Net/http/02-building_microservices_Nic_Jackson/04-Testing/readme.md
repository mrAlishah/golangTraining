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
<img src="Img%2F02-Testing.jpg" width="600"><br/>

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
<img src="Img%2F09-Integerety.png" width="600">
<img src="Img%2F07-Behaviour-Driven Development - Cucumber.png" width="600">
<img src="Img%2F08-BDD.jpeg" width="600">
<img src="Img%2F09-BDD.jpeg" width="600"> 

### Unit-test
Our unit tests go right down to the bottom of the pyramid.

#### Law
* **First law:** You may not write production code until you have written a failing unit test
* **Second law:**  You may not write more of a unit test than is sufficient to fail, and not compiling is failing
* **Third law:**  You may not write more production code than is sufficient to pass the currently failing test

**Note:**
- Production code: Production means anything that you need to work reliably, and consistently.  
  Refer: [here](https://stackoverflow.com/questions/490289/what-exactly-defines-production) <br/>

  <img src="Img%2F10-AAA.jpeg" width="600">
  <img src="Img%2F11-AAA.jpeg" width="600">

