package main

import (
	"errors"
	"fmt"
	"github.com/cenkalti/backoff"
	"time"
)

func main() {
	listCustomers := []string{"test1", "test2", "test3"}
	for _, c := range listCustomers {
		if err := exampleRetry(c); err != nil {
			fmt.Println(err)
		}
	}
}

func exampleRetry(c string) error {
	if c == "test1" {
		fmt.Println(c)
		return nil
	}
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
