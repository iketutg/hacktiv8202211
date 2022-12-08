package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println(" GoodByte from call defer")
	calldeferFunc()
	fmt.Println("Hallo")
	fmt.Println("Golang")
	os.Exit(1)
}

func calldeferFunc() {
	defer funcCetak()

	///.....
}

func funcCetak() {
	fmt.Println("Funct Cetak")
}
