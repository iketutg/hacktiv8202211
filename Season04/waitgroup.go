package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	fmt.Println("Wait Group")

	buah := []string{
		"Mangga",
		"Pisang",
		"Jambu",
		"Nangka",
		"Durian",
	}
	var wg sync.WaitGroup

	for i, v := range buah {
		fmt.Println("Index ke ", i)
		wg.Add(1)
		go CetakBuah(i, v, &wg)
	}
	fmt.Println("Done")
	fmt.Println(" No Goroutine : ", runtime.NumGoroutine())

	wg.Wait()

}

func CetakBuah(idx int, value string, wg *sync.WaitGroup) {
	if idx == 3 {
		time.Sleep(time.Second * 10)
	}
	fmt.Printf("Index : %d  - Buah : %s \n", idx, value)
	wg.Done()
}

func getBuah(name string) {
	fmt.Sprintf("Hello %s", name)
}
