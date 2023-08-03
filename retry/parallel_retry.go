package main

import (
	"fmt"
	"github.com/cenkalti/backoff"
	"time"
)

func main() {
	go execute()

	fmt.Scanln()

}

func execute() {
	trials := []int{1, 2}
	for _, t := range trials {
		if err := exampleRetry(t); err != nil {
			fmt.Println(err)
		}
	}
}

func exampleRetry(trial int) error {
	now := time.Now()

	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = time.Second * 5
	b.InitialInterval = time.Second

	operation := func() error {
		if trial == 1 {
			fmt.Println("example %d", trial)
			return nil
		}
		elapsedSeconds := time.Since(now).Seconds()
		fmt.Println(elapsedSeconds)

		return fmt.Errorf("%s", "retry err")
	}

	if err := backoff.Retry(operation, b); err != nil {
		return err
	}

	return nil
}
