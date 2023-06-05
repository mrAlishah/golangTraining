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
* **statsD:** that provides a simple interface for sending metrics to StatsD.The Statsd Go client allows developers to easily instrument their applications and send metrics to StatsD for monitoring and analysis. It supports various metric types such as counters, gauges, timers.
* **grafana:** is an open-source platform for visualizing and analyzing data, particularly time-series data, from various sources.
* **prometheus:** is a popular open-source monitoring system and time-series database.It is designed to collect metrics from various sources, store them in a time-series database, and provide a powerful query language for analyzing and alerting on those metrics.

Metrics Data -> statsD -> prometheus -> grafana

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

### Grafana
[Documentation](https://grafana.com/docs/grafana/latest/developers/http_api/data_source/)

## Logging
When working with highly distributed containers, you may have 100 instances of your application running rather than one or
two. This means that if you need to grep your log files, you will be doing this over hundreds of files instead of just a couple.<br/>
To save the trouble, the best way to solve this problem is not to write the logs to
disk in the first place. A distributed logging store, such as an ELK stack, or software as a service platform, such as Logmatic or
Loggly, solve this problem for us and give us a fantastic insight into the health and operating condition of our system.<br/>
My personal preference is to only store log data for short periods of time, such as 30 days; this allows you to maintain
diagnostic traces which could be useful for troubleshooting without the cost of maintaining historical data. For historical data, a
metrics platform is best, as you can cheaply store this data over a period of years, which can be useful to compare current
performance with that of a historic event.

### Distributed tracing with Correlation IDs
we used the **header X-Request-ID** which allows us to mark all the service calls for
an individual request with the same ID so that we can later query them.we looked at the header X-Request-ID which allows us to mark all the service calls for
an individual request with the same ID so that we can later query them.
```go
func (c *correlationHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
   if r.Header.Get("X-Request-ID") == "" {
   r.Header.Set("X-Request-ID", uuid.New().String())
   }
   c.next.ServeHTTP(rw, r)
}
```
The handler is implemented using the middleware pattern when we wish to use it all we need to do is wrap the actual handler
like so:
```go
http.Handle("/helloworld", handlers.NewCorrelationHandler(validation))
```
depending upon your requirements you may like to check out **Zipkin** is a distributed tracing system designed to trouble shoot
latency, which is becoming incredibly popular http://zipkin.io. There are also tools from **DataDog**, **NewRelic**, and **AWS X-
Ray**, it is too much to go into depth into these applications, however, please spend an hour and familiarize yourself with their
capabilities as you never know when you are going to need them.

### Elasticsearch, Logstash, and Kibana (ELK)
Elasticsearch, Logstash, and Kibana are pretty much the industry standard when it comes to logging verbose data. All of the
output which would traditionally be streamed to a log file is stored in a central location which you can query with a graphical
interface tool, Kibana.
* **Logstash:** is an open-source data processing pipeline that helps to collect, parse, and store logs for future use.The processed data can then be sent to various destinations, such as Elasticsearch, for indexing and analysis.
* **Elasticsearch:** is a popular open-source search and analytics engine that is often used in logging to help manage and analyze log data. It allows for fast, scalable, and flexible searching and querying of large volumes of data, making it well-suited for logging use cases where organizations need to store and retrieve logs from various sources.
* **Kibana:** is an open-source data visualization and exploration tool used in logging to help users understand and gain insights from log data.With Kibana, users can create custom dashboards, visualizations, and reports to analyze log data.

Logstash can be used to collect logs from various sources and parse them into structured data, which is then sent to Elasticsearch for indexing and storage. Kibana can be used to visualize and explore the log data, allowing users to search, filter, and drill down into specific events or patterns within their log data.

If we look at our Docker Compose file, you will see three entries for our ELK stack:
```dockerfile
elasticsearch:
image: elasticsearch:2.4.2
ports:
- 9200:9200
- 9300:9300
  environment:
  ES_JAVA_OPTS: "-Xms1g -Xmx1g"
  kibana:
  image: kibana:4.6.3
  ports:
- 5601:5601
  environment:
- ELASTICSEARCH_URL=http://elasticsearch:9200
  links:
- elasticsearch
  logstash:
  image: logstash
  command: -f /etc/logstash/conf.d/
  ports:
- 5000:5000
  volumes:
- ./logstash.conf:/etc/logstash/conf.d/logstash.conf
  links:
- elasticsearch
```
* Elasticsearch is our datastore for our logging data, 
* Kibana is the application we will use for querying this data, 
* and Logstash is used for reading the data from your application logs and storing it in Elasticsearch.

Log -> Logstash -> ElasticSearch -> Kibana

logstash config:[documents](https://www.elastic.co/guide/en/logstash/current/configuration.html)
```json
input {
   tcp {
   port => 5000
   codec => "json"
   type => "json"
   }
}
## Add your filters / logstash plugins configuration here
output {
elasticsearch {
   hosts => "elasticsearch:9200"
   }
}
```

### Logrus
we are using the Logstash plugin which allows you to send our logs direct to the Logstash endpoint rather than writing them to a file and then having Logstash pick them up:
```go
func createLogger(address string) (*logrus.Logger, error) {
   retryCount := 0
   
   l := logrus.New()
   hostname, _ := os.Hostname()
   var err error
   
   // Retry connection to logstash incase the server has not yet come up
   for ; retryCount < 10; retryCount++ {
        conn, err := net.Dial("tcp", address)
        if err == nil {
         
            hook := logrustash.New(
              conn,
              logrustash.DefaultFormatter(
                 logrus.Fields{"hostname": hostname},
                 ),
            )
      
            l.Hooks.Add(hook)
            return l, err
         }
   
         log.Println("Unable to connect to logstash, retrying")
         time.Sleep(1 * time.Second)
   }
   
   log.Fatal("Unable to connect to logstash")
   return nil, err
}

```
Adding plugins to Logrus is very simple. We define the hook which is in a separate package, specifying the connection protocol,
address, application name, and a fields collection which is always sent to the logger:
```go
func NewHookWithFields(protocol, address, appName string, alwaysSentFields logrus.Fields) (*Hook, error)
```
We then register the plugin with the logger using the hooks method:
```go
func AddHook(hook Hook)
```

## Exceptions
Golang has two great methods for handling unexpected errors:
* **Panic :** The built-in panic function stops the normal execution of the current goroutine. All the deferred functions are run in the normal
  way then the program is terminated:
* **Recover:** The recover function allows an application to manage the behavior of a panicking goroutine. When called inside a deferred function, recover stops the execution of the panic and returns the error passed to
  the call of panic:
[Look at link](https://medium.com/@jfeng45/go-microservice-with-clean-architecture-application-logging-b43dc5839bce)
```go
func catchPanic() {
	if err := recover(); err != nil {
		logger.Log.Errorf("%+v\n%s", err,debug.Stack())
	}
}
```

Using software as a service, such as **Datadog** or **Logmatic**, is an excellent
way to get up and running very quickly, and alerts integration with **OpsGenie** or **PagerDuty** will allow you to receive instant
alerts whenever a problem may occur.

## summary
* Prometheus: http://localhost:9090
* grafana: http://localhost:3000/ <br/>
user: admin <br/>
pass: admin <br/>

* kibana: http://localhost:5601

## troubleshooting
error docker-compose on macos:
>failed to solve: rpc error: code = Unknown desc = failed to solve with frontend dockerfile.v0: failed to read dockerfile: open /var/lib/docker/tmp/buildkit-mount4245262655/Dockerfile: no such file or directory

solution: before run docker-compose up, run below two commands.
```shell
export DOCKER_BUILDKIT=0
export COMPOSE_DOCKER_CLI_BUILD=0
```