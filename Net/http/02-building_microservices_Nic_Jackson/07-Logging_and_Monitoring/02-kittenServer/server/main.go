package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/alexcesaro/statsd"
	"github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/building-microservices-with-go/chapter7/server/handlers"
	"github.com/sirupsen/logrus"
)

const port = 8091

// curl $(docker-machine ip):8091/helloworld -d '{"name": "Nic"}'
func main() {
	// create stateD instance
	statsd, err := createStatsDClient(os.Getenv("STATSD"))
	if err != nil {
		log.Fatal("Unable to create statsD client")
	}

	// create logrus logger instance
	logger, err := createLogger(os.Getenv("LOGSTASH"))
	if err != nil {
		log.Fatal("Unable to create logstash client")
	}

	setupHandlers(statsd, logger)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func setupHandlers(statsd *statsd.Client, logger *logrus.Logger) {
	// set middleware for helloworld handler
	validation := handlers.NewValidationHandler(
		statsd,
		logger,
		handlers.NewHelloWorldHandler(statsd, logger),
	)

	// set middleware for panic handler as a top lever for managing panic throw from bang handler
	bangHandler := handlers.NewPanicHandler(
		statsd,
		logger,
		handlers.NewBangHandler(),
	)

	// set http route
	// NewCorrelationHandler just set request ID as a middle ware for any request
	// CorrelationHandler -> ValidationHandler -> HelloWorldHandler
	http.Handle("/helloworld", handlers.NewCorrelationHandler(validation))
	// CorrelationHandler -> PanicHandler -> BangHandler
	http.Handle("/bang", handlers.NewCorrelationHandler(bangHandler))
}

func createStatsDClient(address string) (*statsd.Client, error) {
	//    environment:
	//      - "STATSD=statsd:9125"
	//      - "LOGSTASH=logstash:5000"
	return statsd.New(statsd.Address(address))
}

func createLogger(address string) (*logrus.Logger, error) {
	retryCount := 0

	l := logrus.New()
	hostname, _ := os.Hostname()
	var err error

	// Retry connection to logstash incase the server has not yet come up
	for ; retryCount < 10; retryCount++ {
		conn, err := net.Dial("tcp", address)
		if err == nil {

			// add new logstash plugin
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
