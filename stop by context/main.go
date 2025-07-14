package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Stopped by context")
				return
			default:
				fmt.Println("do something")
				time.Sleep(1 * time.Second)
			}
		}
	}()
	time.Sleep(10 * time.Second)
	cancel()
	wg.Wait()
}
