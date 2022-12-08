package main

import "fmt"

func main() {
	message := make(chan string)

	students := []string{"Jono", "Joni", "Joko", "Joe", "Petter"}

	for _, v := range students {

		go func(v string) {
			fmt.Println(" Student : ", v)
			say := fmt.Sprintf("Hello My Name : %s", v)
			message <- say
		}(v)
	}

	for i := 0; i < 5; i++ {
		cetak(message)
	}

	chanint := make(chan int)

	go func(value chan int) {
		fmt.Println("Before send")
		value <- 100
	}(chanint)

	angka := <-chanint
	fmt.Println("Angka : ", angka)
	close(chanint)
}

func cetak(message chan string) {
	fmt.Println(<-message)
}
