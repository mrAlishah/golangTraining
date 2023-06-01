# Back off
Typically, once a connection has failed, you do not want to retry immediately to avoid flooding the network or the server with requests. To allow this, it's necessary to implement a back-off approach to your retry strategy. A back-off algorithm waits for a set period before retrying after the first failure, this then increments with subsequent failures up to a maximum duration.<br/>

Back off is implemented in package go-resiliency package and the retrier package. <br/> 
To download it to your working project
```go
    go get github.com/eapache/go-resiliency/retrier
```

Another package is [link](https://github.com/cenkalti/backoff/blob/v4/example_test.go) 
```go
    go get github.com/cenkalti/backoff
```

### Design for Back off Pattern
To create a Back off Patter for your microservice, we follow steps:<br/>
**Step 1:** To create a new retrier, we use the New function which has the signature:
```go
    r := retrier.New(retrier.ConstantBackoff(3, 1*time.Second), nil)
```
**Step 2:** Run our service with our backoff pattern.
```go
    n := 0
    err := r.Run(func() error {
        fmt.Println("Attempt: ", n)
        n++
        return fmt.Errorf("Failed")
    })
```
### Specification for retrier package.
To create a new retrier, we use the New function which has the signature:
```go
func New(backoff []time.Duration, class Classifier) *Retrier
```

The first parameter is an array of Duration. Rather than calculating this by hand, we can use the two built-in methods which will generate this for us:
```go
func ConstantBackoff(n int, amount time.Duration) []time.Duration
```

The ConstantBackoff function generates a simple back-off strategy of retrying n times and waiting for the given amount of time between each retry:
```go
func ExponentialBackoff(n int, initialAmount time.Duration) []time.Duration
```
The ExponentialBackoff function generates a simple back-off strategy of retrying n times doubling the time between each retry.<br/>

The second parameter is a Classifier. This allows us a nice amount of control over what error type is allowed to retry and what will fail immediately.
```go
type DefaultClassifier struct{}
```
The DefaultClassifier type is the simplest form: if there is no error returned then we succeed; if there is any error returned then the retrier enters the retry state.
```go
type BlacklistClassifier []error
```

The BlacklistClassifier type classifies errors based on a blacklist. If the error is in the given blacklist it immediately fails; otherwise, it will retry.
```go
type WhitelistClassifier []error
```
The WhitelistClassifier type is the opposite of the blacklist, and it will only retry when an error is in the given white list.Any other errors will fail.
