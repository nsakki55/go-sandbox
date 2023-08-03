package main

import (
	"errors"
	"fmt"
	"github.com/cenkalti/backoff"
	"time"
)

func main() {
	if err := exampleRetry(); err != nil {
		fmt.Println(err)
	}
}

func exampleRetry() error {
	now := time.Now()

	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = time.Second * 5
	b.InitialInterval = time.Second

	operation := func() error {
		elapsedSeconds := time.Since(now).Seconds()
		fmt.Println(elapsedSeconds)

		return errors.New("retry err")
	}

	if err := backoff.Retry(operation, b); err != nil {
		return err
	}

	return nil
}
