package main

import (
	"fmt"
	"strings"
)

type isOddNum func(int) bool

func main() {
	//closure callback
	var num = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 13, 14, 15, 16}

	var resFind = findOddNumber(num, func(m int) bool {
		return m%2 != 0
	})
	fmt.Println("Hasil : ", resFind)

}

func findOddNumber(numbers []int, panggil isOddNum) int {
	var totalOdd int
	for _, value := range numbers {
		if panggil(value) {
			totalOdd++
		}
	}
	return totalOdd
}

// closure return value
func closure2() {
	var allStudent = []string{"Jono", "Joko", "Iwan", "Dino", "Budi"}
	findStudent := findStudent(allStudent)
	student := findStudent("Iwan")
	fmt.Println("ditemukan :", student)
}

func findStudent(allStudent []string) func(string) string {
	return func(findStr string) string {
		var student string
		var idx int
		for i, v := range allStudent {
			if strings.ToUpper(v) == strings.ToUpper(findStr) {
				student = v
				idx = i
				break
			}
		}
		if student == "" {
			return fmt.Sprintf("Nama tidak ditemukan : %s", findStr)
		}
		return fmt.Sprintf("Nama %s ditemukan di pos %d ", student, idx)
	}
}

func closure1() {
	var evenNumber = func(number ...int) []int {
		var res []int
		for _, val := range number {
			if val%2 == 0 {
				res = append(res, val)
			}
		}
		return res
	}

	var number = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}

	fmt.Println(evenNumber(number...))

	//apa   == polindrome true
	//mama  == false
	//abcddcba == true

	var isPolindrome = func(str string) bool {
		var temp string = ""
		for i := len(str) - 1; i > 0; i-- {
			temp += string(byte(str[i]))
		}
		return temp != str
	}("katak")

	fmt.Println(isPolindrome)

	var isPolindrome2 = func(str string) bool {
		for i := 0; i < (len(str) / 2); i++ {
			posakhir := len(str) - 1 - i
			if str[i] != str[posakhir] {
				return false
			}
		}
		return true
	}("apa")

	_ = isPolindrome2

}
