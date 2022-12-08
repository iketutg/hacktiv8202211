package main

import (
	"fmt"
	"reflect"
)

func main() {
	var number = 32
	var rofV = reflect.ValueOf(number)

	fmt.Println("Tipe variable : ", rofV.Type())

	if rofV.Kind() == reflect.Int {
		fmt.Println("Nilai Variable :", rofV.Int())
	}

	fmt.Println(" Value : ", rofV.Interface())
	fmt.Printf(" Value : %T\n", rofV.Interface())
	var i = rofV.Interface().(int)
	fmt.Println(i)

	var mhs = &student{
		Name:  "Jono",
		Grade: 2,
	}
	fmt.Println(" Name : ", mhs.Name)
	fmt.Println("Grade : ", mhs.Grade)

	var reflect1 = reflect.ValueOf(mhs)

	var method = reflect1.MethodByName("SetName")
	fmt.Println("Method : ", &method)

	method.Call([]reflect.Value{
		reflect.ValueOf("Ketut"),
	})

	fmt.Println("Nama : ", mhs.Name)
	fmt.Println("Nama : ", mhs.Grade)
}

func (mhs *student) SetName(name string) {
	mhs.Name = name
}

type student struct {
	Name  string
	Grade int
}
