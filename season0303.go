package main

import "fmt"

/*
Pointer
*/
func main() {

}

func pointer3() {
	var number int = 1000
	fmt.Println("Before ", number)
	changeValue(&number)
	fmt.Println("after ", number)
}

func changeValue(angka *int) {
	*angka = *angka * 10
}

func pointer2() {
	var personA string = "iketut"
	var personB *string = &personA

	fmt.Println("PersonA  value : ", personA)
	fmt.Println("PersonA  mem address : ", &personA)
	fmt.Println("PersonB value : ", *personB)
	fmt.Println("PersonB mem address : ", personB)
	*personB = "Gunawan"

	fmt.Println("PersonA  value : ", personA)
	fmt.Println("PersonA  mem address : ", &personA)
	fmt.Println("PersonB value : ", *personB)
	fmt.Println("PersonB mem address : ", personB)
}

func pointer1() {
	var angka int = 10
	var angka2 *int = &angka

	fmt.Println("Angka Number value : ", angka)
	fmt.Println("Angka Number mem address : ", &angka)
	fmt.Println("angka2 Number : ", *angka2)
	fmt.Println("Angka2 Number mem address : ", angka2)

}
