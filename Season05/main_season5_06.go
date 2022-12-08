package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Hello !"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "Goodbye"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch1:
			fmt.Println("Receive Message Channel 1 ", msg)
		case msg2 := <-ch2:
			fmt.Println("Receive Message Channel 2 ", msg2)
		}
	}

}
