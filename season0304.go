package main

import "fmt"

type Employee struct {
	name   string
	age    int
	divisi string
}

type Student struct {
	name    string
	age     int
	ipk     float32
	balance int
}

type Person struct {
	name string
	age  int
}

/*
Structure
*/
func main() {
	//structure01()
	var peoplee = []Person{
		{name: "Jono", age: 30},
		{name: "Joko", age: 40},
		{name: "Budi", age: 50},
	}

	for _, val := range peoplee {
		fmt.Println("Name : ", val)
	}
}
func structure01() bool {
	student := getDataStudent()
	fmt.Println(student)
	ok := getBalance(&student)
	if ok {
		fmt.Println("Mendapatkan tambahan balance : ", student)
		return ok
	}
	fmt.Println("Tidak mendapakan tambahan : ", student)
	return ok
}

func getBalance(st *Student) bool {
	if st.ipk > 38.0 {
		st.balance += 100
		return true
	}
	return false
}
func getDataStudent() Student {
	return Student{
		age:     20,
		name:    "Budi",
		ipk:     39.9,
		balance: 100,
	}
}

func latihanStruct1() {
	var emp Employee
	emp.divisi = "Developer"
	emp.age = 45
	emp.name = "Gunawan"
	fmt.Println("Structure Employee ", emp)
	fmt.Println("Name Employee ", emp.name)
	var emp2 = Employee{
		age:    10,
		name:   "Budi",
		divisi: "Sekolah",
	}
	fmt.Println("Structure Employee ", emp2)
	fmt.Println("Name Employee ", emp2.name)
}
