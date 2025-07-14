package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Hello wildberries, i want to work here")
		time.Sleep(1 * time.Second)
		runtime.Goexit()
		fmt.Println("I love movie Clerks by Kevin Smith, but nobody will know it, becouse of runtime.Goexit()")
	}()

	wg.Wait()
}
