package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		start := time.Now()
		for {
			if time.Since(start) > 5*time.Second {
				fmt.Println("время вышло - выходим")
				return
			}
		}
	}()
	wg.Wait()
}
