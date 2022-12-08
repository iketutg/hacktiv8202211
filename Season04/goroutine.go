package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//go Hello()
	runtime.GOMAXPROCS(4)
	fmt.Println("Go rouinte ")

	go Process1(10)
	Process2(10)

	fmt.Println("No of Gorutine ", runtime.NumGoroutine())
	time.Sleep(time.Second * 5)

}

func Process2(max int) {
	fmt.Println("Process 2 called")
	for i := 0; i < max; i++ {
		fmt.Println("Process 2 Index : ", i)
		if i == 5 {
			time.Sleep(time.Second * 1)
		}
	}
	fmt.Println("Process 2 end")
}

func Process1(max int) {
	fmt.Println("Process 1 called")
	for i := 0; i < max; i++ {
		fmt.Println("Process 1 Index : ", i)

	}
	fmt.Println("Process 1 end")
}

func Hello() {
	fmt.Println("Go routine HEllo")
}
