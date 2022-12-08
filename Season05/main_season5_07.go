package main

import "fmt"

func getMax(number []int, maxchan chan int) {
	var max = number[0]
	for _, v := range number {
		if max < v {
			max = v
		}
	}
	maxchan <- max
}

func main() {

	var number = []int{1, 3, 4, 5, 26, 6, 4, 3, 4, 6, 6, 9, 10, 30, 50, 1, 2, 3, 4, 5, 6, 8}
	maxchan := make(chan int)
	go getMax(number, maxchan)

	avgchan := make(chan float64)
	go getAverage(number, avgchan)
	var max int
	var avg float64
	for i := 0; i < 2; i++ {
		select {
		case max = <-maxchan:
			fmt.Println("Max : ", max)
		case avg = <-avgchan:
			fmt.Println("Avg : ", avg)
		}
	}
}

func getAverage(number []int, avgchan chan float64) {
	sum := 0
	for _, v := range number {
		sum += v
	}
	avgchan <- float64(sum) / float64(len(number))
}
