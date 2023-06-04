# Logging and Monitoring
* **Metrics:** These are things such as time series data (for example, transaction or individual component timings).
* **Text-based logs:** Text-based records are your real old-fashioned logs which are spat out by things such as Nginx or a text log from your application software.
* **Exceptions:** Exceptions potentially could fall into the two previous categories; however, I like to break these out into a separate category since exceptions should be, well, exceptional.

## Logging best practices
In the free e-book, The pragmatic logging handbook, by Jon Gifford of Loggly (www.loggly.com), Jon proposes the following
eight best practices to apply when determining your logging strategy:
* Treat application logging as an ongoing iterative process. Log at a high level and then add deeper instrumentation.
* Always instrument anything that goes out of the process because distributed system problems are not well behaved.
* Always log unacceptable performance. Log anything outside the range in which you expect your system to perform.
* If possible, always log enough context for a complete picture of what happened from a single log event.
* View machines as your end consumer, not humans. Create records that your log management solution can interpret.
* Trends tell the story better than data points.
* Instrumentation is NOT a substitute for profiling and vice versa.
* Flying more slowly is better than flying blind. So the debate is not whether to instrument, just how much.

## Metrics
In my opinion, metrics are the most useful form of logging for day-to-day operations. Metrics are useful because we have
simple numeric data diagnosing performance problems with the service. We can plot this onto a time series dashboard and quite quickly set up alerting from the output as the data
is incredibly cheap to process and collect.<br/>
No matter what you are storing, the superior efficiency of metrics is that you are storing numeric data in a time-series database
using a unique key as an identifier.<br/>
for example:
* request timings
* request counts
* request success, and failure counts for handlers
* Exhausted CPU on host server
* Exhausted memory
* Network latency
* Slow data store queries
* Latency with downstream service caused by any of the preceding factors

### Naming conventions
I recommend you break up the name of your service using dot notation such as the following:
> environment.host.service.group.segment.outcome

* **environment:** This is the working environment; for example: production, staging
* **host:** This is the hostname of the server running the application
* **service:** The name of your service
* **group:** This is the top level grouping; for an API, this might be handlers
* **segment:** The child level information for the group; this will typically be the name of the handler in the instance of an API
* **outcome:** This is something which denotes the result of the operation, in an API you may have called, success, or you may choose to use HTTP status codes

Here is an example of how to use the following dot notation:
>prod.server1.kittenserver.handlers.list.ok
>prod.server1.kittenserver.mysql.select_kittens.timing

If your monitoring solution supports tags in addition to the event name, then I recommend you use tags for the environment and
host, this will make querying the data store a little easier. For example, if I have a handler which lists kittens which are running
on my production server then I may choose to add the following events to be emitted when the handler is called:
```go
func (h *list) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
    
	event := startTimingEvent("kittens.handlers.list.timing", ["production", "192.168.2.2"])
    defer event.Complete()
    
	dispatchIncrementEvent("kittens.handlers.list.called", ["production", "192.168.2.2"])
	
...
   
    if err != nil {
        dispatchIncrementEvent("kittens.handlers.list.failed", ["production", 192.168.2.2"])
    return
    }
	
    dispatchIncrementEvent("kittens.handlers.list.success", ["production", 192.168.2.2"])
}
```
This is a pseudo code, but you can see that we are dispatching three events from this handler:
1. The first event is that we are going so send some timing information.
2. In the next, we are simply going to send an increment count which is simply going to state that the handler has been
   called.
3. Finally, we are going to check if the operation has been successful. If not, we increment our handler-failed metric; if
   successful, we increment our success metric.

Naming our metrics in this way allows us to graph errors either on a granular level or makes it possible to write a query which
is at a higher level. For example, we may be interested in the total number of failed requests for the entire service, not just this
endpoint. Using this naming convention, we can query using wildcards; so to query all failures for this service, we could write a
metric like the following code:
> kittens.handlers.*.failed

If we were interested in all failed requests to handlers for all services, we could write the following query:
> `*.handlers.*.failed`

### stateD
We are using an open source package by Alex Cesaro (https://github.com/alexcesaro/statsd). This has a very simple
interface; to create our client, we call the new function and pass it a list of options. In this instance, we are only passing through
the address of the statsD server, which has been set by an environment variable:
```go
func New(opts ...Option) (*Client, error)
```
**Note:**The statsD client does not work synchronously, sending each metric when you make a call to the client; instead, it buffers all
the calls, and there is an internal goroutine which sends the data at a predetermined interval. This makes the operation highly
efficient, and you should not have to worry about any application slowdown.

```go
func createStatsDClient(address string) (*statsd.Client, error) {
   // it needs Storage and querying as a service for stateD metrics data
   //There are multiple options for storing and querying metric data; you have the possibility for self-hosting, or you can utilize a
   //software as a service.we will implement it in docker-compose by image: prom/statsd-exporter
   return statsd.New(statsd.Address(address))
   //return statsd.New()
}
```
In handlers.helloworld:

```go
const (
	helloworldSuccess string = "kittenserver.helloworld.success"
	helloworldFailed  string = "kittenserver.helloworld.failed"
	helloworldTiming  string = "kittenserver.helloworld.timing"
)

func (h *helloWorldHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
   timing := h.statsd.NewTiming()
   ...
   
   err := decoder.Decode(&request)
      if err != nil {
      // metric data
      h.statsd.Increment(helloworldFailed)
      ...
   }
   
   ...
   // metric data
   h.statsd.Increment(helloworldSuccess)
   // just simulate for delay
   time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
   
   //metric data
   //h.statsd.NewTiming().Send(helloworldTiming)
   timing.Send(helloworldTiming)
}
```

### Storage and querying
There are multiple options for storing and querying metric data; you have the possibility for self-hosting, or you can utilize a software as a service.

1. **Software as a service**
For software as a service (SaaS), I recommend looking at **Datadog**. To send metrics to Datadog, you have two options: one is to
   communicate with the API directly; the other is to run the Datadog collector as a container inside your cluster. The Datadog
   collector allows you to use StatsD as your data format and it supports a couple of nice extensions which standard StatsD does
   not, such as the ability to add additional tags or metadata to your metrics. Tagging allows you to categorize your data by user-
   defined tags, this allows you to keep your metric names specific to what they are monitoring without having to add
   environmental information.
2. **Self-hosted**
While it may be desirable to use a SaaS service for your production data, it is always useful to be able to run a server locally for
local development. There are many options for backend data stores such as **Graphite**, **Prometheus**, **InfluxDB**, and **ElasticSearch**;
however, when it comes to graphing, Grafana leads the way.<br/>
Let's spin up a Docker Compose stack for our list, kittenservice, so we can run through the simple steps of setting up
Prometheus with Grafana with Docker Compose.
If we look at the Docker compose file, we can see that we have three entries:
* statsD
* grafana
* prometheus

```dockerfile
prometheus:
image: prom/prometheus
links:
- statsd
volumes:
- ./prometheus.yml:/etc/prometheus/prometheus.yml
ports:
- 9090:9090
```