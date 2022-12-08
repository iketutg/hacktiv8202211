package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type Mahasiswa struct {
	Nim  string `required:"true"`
	Nama string
	Age  int
}

type Dosen struct {
	Nid  string
	Name string `required:"true" max:"10"`
	Age  int
}

func main() {
	mhs := Mahasiswa{Nim: "1234567890"}
	var refType reflect.Type = reflect.TypeOf(mhs)
	fmt.Println(" Jumlah Field ", refType.NumField())
	fmt.Println("Field Name : ", refType.Field(0).Name)
	fmt.Println("Field Type : ", refType.Field(0).Type)
	fmt.Println("Field Tag : ", refType.Field(0).Tag)
	fmt.Println("Value from Field Get Key Requeired : ", refType.Field(0).Tag.Get("required"))
	fmt.Println("========")
	fmt.Println("Mahasiswa ", mhs)
	invalid := isInvalid(mhs)
	fmt.Println("Return 1: ", invalid)

	mhs1 := Mahasiswa{Nama: "iketutg", Age: 40}
	fmt.Println("Mahasiswa ", mhs1)
	invalid = isInvalid(mhs1)
	fmt.Println("Return 2: ", invalid)
	fmt.Println("##################")
	dosen := Dosen{Name: "Gunawan ..............."}
	var refD reflect.Type = reflect.TypeOf(dosen)
	fmt.Println(" Jumlah Field ", refD.NumField())
	fmt.Println("Field Name : ", refD.Field(1).Name)
	fmt.Println("Field Type : ", refD.Field(1).Type)
	fmt.Println("Field Tag : ", refD.Field(1).Tag)
	fmt.Println("Value from Field Get Key Requeired : ", refD.Field(1).Tag.Get("required"))
	fmt.Println("Value Tag Max: ", refD.Field(1).Tag.Get("max"))
	fmt.Println("========")

	fmt.Println("Hasil validasi : ", isInvalidV2(dosen))

}

func isInvalidV2(data interface{}) bool {
	of := reflect.TypeOf(data)
	for i := 0; i < of.NumField(); i++ {
		field := of.Field(i)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == "" {
				return false
			}
			vstring, ok := value.(string)
			if !ok {
				return false
			}
			max, _ := strconv.Atoi(field.Tag.Get("max"))

			if len(vstring) > max {
				return false
			}

		}
	}

	return true
}

func isInvalid(data interface{}) bool {
	of := reflect.TypeOf(data)
	for i := 0; i < of.NumField(); i++ {
		field := of.Field(i)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == "" {
				return false
			}
		}
	}

	return true
}
