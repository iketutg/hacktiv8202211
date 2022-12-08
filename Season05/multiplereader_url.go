package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	//IP Address Class
	urls := []string{"http://www.google.com", "http://www.yahoo.com", "http://www.detik.com", "http://www.kompas.com"}

	chanurl := make(chan string, len(urls))
	message := make(chan string, len(urls))

	for i := 0; i < len(urls); i++ {
		go worker(chanurl, message, i)
	}

	for _, value := range urls {
		chanurl <- value
	}

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-message)
	}

}

func worker(chanurl chan string, message chan string, i int) {
	for {
		url := <-chanurl
		length, err := getPageUrl(url)
		if err == nil {
			message <- fmt.Sprintf("ID : %d URL : %s  Length Data : %d", i, url, length)
		} else {
			message <- fmt.Sprintf("Error : %s ID : %d URL : %s  Length Data : %d", err.Error(), i, url, length)
		}
	}
}

func getPageUrl(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	all, err1 := io.ReadAll(resp.Body)
	if err1 != nil {
		return 0, err1
	}
	return len(all), nil
}
