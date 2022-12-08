package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Season 5 Channel")

	channel := make(chan string)

	channel <- "Hallo Channel"

	time.Sleep(time.Second * 2)
	results := <-channel
	fmt.Println("Result Channel : ", results)

	//close(channel)
}
