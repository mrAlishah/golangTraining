first also we know<br>
Go is a concurrent language and not a parallel one
# What is concurrency?
Concurrency is the capability to deal with lots of things at once
# What is parallelism and how is it different from concurrency?

Parallelism is doing lots of things at the same time.
# Goroutines
Goroutines are functions or methods that run concurrently with other functions or methods<br><br>
Goroutine vs Thread
Goroutine: A Goroutine is a function or method which executes independently and simultaneously in connection with any other Goroutines present in your program. Or in other words, every concurrently executing activity in Go language is known as a Goroutines. <br><br>
Thread: A process is a part of an operating system which is responsible for executing an application. Every program that executes on your system is a process and to run the code inside the application a process uses a term known as a thread. A thread is a lightweight process, or in other words, a thread is a unit which executes the code under the program. So every program has logic and a thread is responsible for executing this logic.<br><br>

1.Goroutines are extremely cheap
2.The Goroutines are multiplexed to a fewer number of OS threads
3.Goroutines communicate using channels.
**Channels by design prevent race conditions from happening when accessing shared memory using Goroutines.**<br>
Channels can be thought of as a pipe using which Goroutines communicate<br>
use "go" for create goroutine<br><br><br>
- When a new Goroutine is started, the goroutine call returns immediately. Unlike functions, the control does not wait for the Goroutine to finish executing. The control returns immediately to the next line of code after the Goroutine call and any return values from the Goroutine are ignored.<br><br><br>
- he main Goroutine should be running for any other Goroutines to run. If the main Goroutine terminates then the program will be terminated and no other Goroutine will run.