# Circuite Breaking
### Why we need Circuite Breaking Pattern
In a simpel word, we can say that we use this pattern, to open connection again to avoid failing error related to timeout with a service which have a large time processing.
Michael Nygard in his book "Release It" says:
>Circuit breakers are a way to automatically degrade functionality when the system is under stress.

### Work flow Circuite Pattern
* **1.** Under normal operations, like a circuit breaker in your electricity switch box, the breaker is closed and **1.1** traffic flows normally.
* **1.2** However, once the pre-determined error threshold has been exceeded, the breaker enters the open state, and all requests **(include 2. -> 2.1, 3. ->3.2)** immediately fail without even being attempted.
* After a period, a further request would be allowed and the circuit enters a halfopen state, in this state a failure immediately returns to the open state regardless of the errorThreshold.
* Once some requests have been processed without any error, then the circuit again returns to the closed state, and only if the number of failures
  exceeded the error threshold would the circuit open again.
  That gives us a little more context to why we need circuit breakers, but how can we implement them in Go?

![image](../Img/01-circuuit_breaking)
### Getting Started with Circuit Pattern
Circuite Pattern can be found at **go-resilience package** , it's called **breaker**. More details : [circuit-breaker](https://github.com/eapache/go-resiliency/blob/master/breaker/README.md)


To create a circuit breaker
```
func New(errorThreshold, successThreshold int, timeout time.Duration) *Breaker
```
We construct our circuit breaker with three parameters:
* The first errorThreshold, is the number of times a request can fail before the circuit opens
* The successThreshold, is the number of times that we need a successful request in the half-open state before we move back to open
* The timeout, is the time that the circuit will stay in the open state before changing to half-open

Design circuit pattern
```
	b := breaker.New(3, 1, 5*time.Second)

	for {
		result := b.Run(func() error {
			// Call some service
			time.Sleep(2 * time.Second)
			return fmt.Errorf("Timeout")
		})

		switch result {
		case nil:
			// success!
		case breaker.ErrBreakerOpen:
			// our function wasn't run because the breaker was open
			fmt.Println("Breaker open")
		default:
			fmt.Println(result)
		}

		time.Sleep(500 * time.Millisecond)
	}
```

Could you explain this output:
```
Timeout
Timeout
Timeout
Breaker open
Breaker open
Breaker open
...
Breaker open
Breaker open
Timeout
Breaker open
Breaker open
```
## Hystix library from Netflix
* One of the more modern implementations of circuit breaking and timeouts is the Hystix library from Netflix; Netflix is certainly renowned for producing some quality microservice architecture and the Hystrix client is something that has also been copied time and time again.
* Hystrix is described as "a latency and fault tolerance library designed to isolate points of access to remote systems, services, and third-party libraries, stop cascading failure, and enable resilience in complex distributed systems where failure is inevitable."

### Hystrix from Github
More details explanation: <br/>
[Netflix main source](https://github.com/Netflix/Hystrix)
[Golang Netflix Package source](https://github.com/afex/hystrix-go)
