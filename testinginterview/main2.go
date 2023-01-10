package main

import "fmt"

func main() {
	msgChan := make(chan string)

	go func() {
		for i := 0; i < 5; i++ {
			msgChan <- fmt.Sprintf("Data ke %d", i)
		}
	}()

	for msg := range msgChan {
		fmt.Println("Pesan :", msg)
	}

}
