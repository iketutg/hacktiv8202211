package main

import "fmt"

func PrintAll(values []interface{}) {
	for _, v := range values {
		fmt.Println(v)
	}
}

func Cetak(data interface{}) {
	fmt.Println("Cetak ", data)
}

func main() {
	str := []string{"halo", "hi", "Loha"}

	Cetak(str)

	strInterface := make([]interface{}, len(str))
	for i, v := range str {
		strInterface[i] = v
	}
	PrintAll(strInterface)
}
