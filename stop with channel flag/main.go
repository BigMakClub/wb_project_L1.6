package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	done := make(chan struct{})

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-done:
				fmt.Println("Горутина завершила свою работу")
				return
			default:
				fmt.Println("Что-то делаем")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()
	fmt.Println("Ждем 10 итераций и закрваем канал")
	for i := 0; i < 10; i++ {
		fmt.Printf("Итерация %d\n", i)
		time.Sleep(300 * time.Millisecond)
	}
	close(done)
	wg.Wait()
}
