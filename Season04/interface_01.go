package main

import "fmt"

func main() {
	var name interface{}
	name = "Joni"

	v, ok := name.(string)
	fmt.Println("Interface value : ", v)
	fmt.Println("Interface type : ", ok)

	var saldo interface{}
	saldo = 100000

	v1, ok1 := saldo.(int)
	fmt.Println("Interface value : ", v1)
	fmt.Println("Interface true or false  : ", ok1)

}
