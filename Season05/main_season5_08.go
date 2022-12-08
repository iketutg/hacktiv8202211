package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Channel timeout")

	var msg = make(chan int)

	go sendData(msg)
	receiveData(msg)
}

func receiveData(msg <-chan int) {
loopingout:
	for {
		select {
		case data := <-msg:
			fmt.Println("Receive Data : ", data)
		case <-time.After(time.Second * 8):
			fmt.Println("Waktu Habis 8 detik")
			break loopingout
		}
	}
}

func sendData(msg chan<- int) {
	for i := 0; true; i++ {
		msg <- i
		timesleep := rand.Int()%10 + 1
		fmt.Println("Waktu sleep nya : ", timesleep)
		time.Sleep(time.Duration(timesleep) * time.Second)
	}
}
