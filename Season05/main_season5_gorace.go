package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var counter int
var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go increment("Joko")
	go increment("JoNo")
	wg.Wait()
	fmt.Println("Final Counter : ", counter)
}

func increment(data string) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 20; i++ {
		x := counter
		x++
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		counter = x
		fmt.Printf("Data : %s Index : %d  Counter : %d\n", data, i, counter)
	}
	wg.Done()
}
