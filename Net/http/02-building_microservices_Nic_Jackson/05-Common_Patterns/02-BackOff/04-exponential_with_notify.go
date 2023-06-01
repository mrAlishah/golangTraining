package main

import (
	"fmt"
	"log"
	"time"

	"github.com/cenkalti/backoff"
)

func main() {
	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = 3 * time.Minute

	var (
		//val int64
		err error
	)

	retryable := func() error {
		_, err = doSomething()
		return err
	}

	notify := func(err error, t time.Duration) {
		log.Printf("error: %v happened at time: %v", err, t)
	}

	err = backoff.RetryNotify(retryable, b, notify)
	if err != nil {
		log.Fatalf("error after retrying: %v", err)
	}
}

func doSomething() (int64, error) {
	fmt.Println("Do something Done")
	return 6, nil
}
