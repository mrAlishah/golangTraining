package main

import (
	"fmt"
	"log"
	"time"

	"github.com/cenkalti/backoff"
)

func main() {
	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = 10 * time.Second

	doSomething := func() error {
		//do something and return error
		fmt.Println("Do something done")

		return nil
	}

	err := backoff.Retry(doSomething, b)
	if err != nil {
		log.Fatalf("error after retrying: %v", err)
	}
}
