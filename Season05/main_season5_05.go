package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)
	var channel = make(chan string)
	go sendMessage(channel)
	printMessage(channel)
}

func printMessage(message <-chan string) {
	for value := range message {
		fmt.Println(value)
	}
}

func sendMessage(channel chan<- string) {
	for i := 0; i < 10; i++ {
		channel <- fmt.Sprintf("Message Id : %d", i)
	}
	close(channel)
}
