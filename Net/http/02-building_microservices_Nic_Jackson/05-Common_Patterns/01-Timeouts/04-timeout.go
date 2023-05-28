package main

import (
	"errors"
	"fmt"
	"github.com/eapache/go-resiliency/deadline"
	"time"
)

func main() {
	var waitTimeout = 4 * time.Millisecond
	//var waitTimeout = 10  * time.Millisecond

	dl := deadline.New(waitTimeout)

	if err := dl.Run(takesFiveMillis); err != nil {
		fmt.Println("takesFiveMillis: ", err.Error())
	}

	if err := dl.Run(takesTwentyMillis); err == deadline.ErrTimedOut {
		fmt.Println("takesTwentyMillis: ", err.Error())
	}

	if err := dl.Run(returnsError); err.Error() == "foo" {
		fmt.Println("returnsError: ", err.Error())
	}

	// https://stackoverflow.com/questions/52035390/why-using-chan-struct-when-wait-something-done-not-chan-interface
	// https://stackoverflow.com/a/52035899
	done := make(chan struct{}) // for waiting app time process to finish app
	err := dl.Run(func(stopper <-chan struct{}) error {
		// when you read stopper , it means deadline.ErrTimedOut will return
		// without read process finish completely
		<-stopper
		close(done)
		return nil
	})

	if err == deadline.ErrTimedOut {
		fmt.Println("done: ", err.Error())
	}
	<-done

	//time.Sleep(2 * time.Second)
}

func takesFiveMillis(stopper <-chan struct{}) error {
	time.Sleep(5 * time.Millisecond)
	return nil
}

func takesTwentyMillis(stopper <-chan struct{}) error {
	time.Sleep(20 * time.Millisecond)
	return nil
}

func returnsError(stopper <-chan struct{}) error {
	// Note:when you read stopper means the deadline was occurred and return Error not return
	// return deadline.ErrTimedOut Not errors.New("foo")
	//fmt.Println(<-stopper)

	return errors.New("foo")
}
