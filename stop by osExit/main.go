package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	defer fmt.Println("main завершился по-нормальному")
	go func() {
		fmt.Println("doing something")
		time.Sleep(time.Second)
		fmt.Println("вызываю os.Exit(0)")
		os.Exit(0)
		fmt.Println("Это никто не уивидит ")
	}()

	for i := 0; i < 5; i++ {
		fmt.Printf("итерация в main %d\n", i)
		time.Sleep(300 * time.Millisecond)
	}
	fmt.Println("Main дошел до конца")
}
