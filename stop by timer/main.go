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
		timer := time.NewTimer(5 * time.Second)
		for {
			select {
			case <-timer.C:
				fmt.Println("Завершили горутину через 5 секунд")
				return
			default:
				fmt.Println("Возьмите меня на стажу в WB")
				time.Sleep(300 * time.Millisecond)
			}
		}
	}()
	fmt.Println("Завершили main")
	wg.Wait()
}
