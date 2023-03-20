package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/pkg/errors"
)

func shortProcess(ctx context.Context) error {
	time.Sleep(2 * time.Second)
	return errors.New("failed")
}

func longProcess(ctx context.Context) {

	for {
		select {
		case <-time.After(5 * time.Second):
			fmt.Println("done")
		case <-ctx.Done():
			fmt.Println("stopped long process")
		default:
			fmt.Println("ha...")
		}
	}

}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			if err := shortProcess(ctx); err != nil {
				logrus.WithError(err).Error("Something goes terrible wrong")
				cancel()
			}
			wg.Done()
		}()
	}
	go longProcess(ctx)
	wg.Wait()
}
