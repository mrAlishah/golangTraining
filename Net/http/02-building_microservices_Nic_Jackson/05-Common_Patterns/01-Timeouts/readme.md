# Timeouts
A timeout is an incredibly useful pattern while communicating with other services or data stores. The idea is that you set a limit on the response of a server and, if you do not receive a response in the given time, then you write a business logic to deal with this failure, such as retrying or sending a failure message back to the upstream service.
* Connection Timeout - The time it takes to open a network connection to the server
* Request Timeout - The time it takes for a server to process a request

The Timeout pattern is a common pattern used in event-driven systems to handle situations where a response to an event or request has not been received within a specified timeframe.<br/>
In this pattern, a timer is set when an event or request is made, and if the expected response is not received before the timer expires, a timeout event is triggered. The system can then take appropriate action based on the timeout event, such as retrying the request or notifying the user that the operation has failed.<br/>

To create timeout for any service, we follow steps below:
**Step 1:** Use go-resiliency package to create a instance deadline to set timeout
```go
dl := deadline.New(1 * time.Second)
```
**Step 2:** Use instance deadline to run our service
```go
err := dl.Run(func(stopper <-chan struct{}) error {
		slowFunction() <=============== This is our service
		return nil
	})
```
**Step 3:** Write a business logic to deal with this failure, such as retrying or sending a failure message back to the upstream service
```go
	switch err {
	case deadline.ErrTimedOut:
		fmt.Println("Timeout")
	default:
		fmt.Println(err)
	}
```
----
some code:
```go
var waitTimeout = 4 * time.Millisecond
dl := deadline.New(waitTimeout)

if err := dl.Run(takesFiveMillis); err != nil {  //err == deadline.ErrTimedOut
	fmt.Println("takesFiveMillis: ", err.Error())
}
...
func takesFiveMillis(stopper <-chan struct{}) error {
    time.Sleep(5 * time.Millisecond)
    return nil
}

```

[about struct{}](https://stackoverflow.com/a/52035899)
```go
done := make(chan struct{}) // for waiting app time process to finish app
err := dl.Run(func(stopper <-chan struct{}) error {
    // when you read stopper , it means deadline.ErrTimedOut will return
    // without read stopper, process finish well
    <-stopper
    close(done)
    return nil
})

	if err == deadline.ErrTimedOut {
		fmt.Println("done: ", err.Error())
	}

```

### Context package
Package context defines the Context type, which carries deadlines, cancellation signals, and other request-scoped values across API boundaries and between processes.<br/>
More details: [here](https://pkg.go.dev/context)

### Note: Implement timeout using Context