package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	var num int
	var err error

	num, err = strconv.Atoi("123444B")
	if err == nil {
		fmt.Println("Number : ", num)

	} else {
		fmt.Println("Error ", err)
		fmt.Println(" Error : ", errors.New("Angka Anda Salah bukan Numeric"))
	}

}
