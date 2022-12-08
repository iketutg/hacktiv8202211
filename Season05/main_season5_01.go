package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Buffer Channel")

	channel := make(chan int, 2)

	fmt.Println(cap(channel))
	fmt.Println(len(channel))

	go func(ch chan int) {
		for i := 0; i < 5; i++ {
			fmt.Printf("Func go Routine %d start send data\n", i)
			ch <- i
			fmt.Printf("Func go Routine %d after send data \n", i)
		}
		close(ch)
	}(channel)

	fmt.Println("Main Go routine ")
	time.Sleep(time.Second * 3)

	for v := range channel {
		fmt.Println("Main Go Routine in range ", v)
	}

}
