package main

import (
	"fmt"
	"net/http"
)

func midleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Eksekusi midleware sebelum bisnis logic")
		handler.ServeHTTP(writer, request)
		fmt.Println("Eksekusi midleware setelah response bisnis logic")
	})
}

func bisnisLogic(writer http.ResponseWriter, request *http.Request) {
	//disini proses bisnis logic
	fmt.Println("Eksekusi bisnis logic handler")
	writer.Write([]byte("OK"))
}

func main() {
	bisnisLogicHandler := http.HandlerFunc(bisnisLogic)
	http.Handle("/", midleware(bisnisLogicHandler))
	http.ListenAndServe(":8000", nil)
}
