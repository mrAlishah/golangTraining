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
<img src="Img%2F03-Testing.jpg" width="600">
<img src="Img%2F05-BDD.png" width="600">
<img src="Img%2F06-BDD-TDD.png" width="600">
<img src="Img%2F07-Behaviour-Driven Development - Cucumber.png" width="600">
<img src="Img%2F08-BDD.jpeg" width="600">
<img src="Img%2F09-BDD.jpeg" width="600">
