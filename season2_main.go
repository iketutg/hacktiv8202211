package main

import (
	"fmt"
	"math"
	"strings"
)

const (
	MAHASISWA = "mahasiswa"
	KARYAWAN  = "karyawan"
	DOSEN     = "dosen"
)

func main() {
	temp := []float64{
		-29.0, -28.0, 32.0, -31.0, -23.0, -29.0, -33.0, -28.0,
	}
	gropuptemp := make(map[float64][]float64)
	for _, v := range temp {
		nilai := math.Trunc(v/10) * 10
		gropuptemp[nilai] = append(gropuptemp[nilai], v)
	}
	fmt.Println(gropuptemp)

}

func StringExplore() {
	string1 := "Programming"
	string2 := "Programming Go"
	string3 := "Programming"

	fmt.Println(strings.Compare(string1, string2)) //-1 lebih kecil
	fmt.Println(strings.Compare(string1, string3)) //0
	fmt.Println(strings.Compare(string2, string3)) // 1 lebih besar

	text := "Go Programming"
	substring1 := "Go"
	substring2 := "Golang"

	fmt.Println(strings.Contains(text, substring1))
	fmt.Println(strings.Contains(text, substring2))

	text2 := "car car "
	resultText := strings.Replace(text2, "r", "t", 3)
	fmt.Println(resultText)
}

func alias() {
	type SECOND = uint
	var detik SECOND = 50
	fmt.Println(detik)

	var x uint8 = 10
	var y byte

	y = x
	fmt.Println(x)
	fmt.Println(y)
}
func map2() {
	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar golang"
	book["author"] = "ketut"
	book["auth-email"] = "iketutg@ikg.com"
	fmt.Println("Before delete : ", book)
	fmt.Println("Before delete len : ", len(book))

	delete(book, "auth-email")
	fmt.Println("after delete : ", book)
	fmt.Println("after delete len : ", len(book))

	value, ok := book["auth-email"]
	fmt.Printf("Value : %v - %v\n", value, ok)

	value, ok = book["author"]
	fmt.Printf("Value : %v - %v\n", value, ok)
}
func map1() {
	var person map[string]string
	person = map[string]string{}
	person["name"] = "iketut"
	person["age"] = "45"
	person["address"] = "Jalan Pidana"
	fmt.Println(person)
	fmt.Println("Person name ", person["name"])
	fmt.Println("Person age ", person["age"])

	for k, v := range person {
		fmt.Printf("K : %s ,  v :%s\n", k, v)
	}

	student := map[string]string{
		"ipk":      "3,6",
		"fullname": "Iketut Gunawan",
		"email":    "iketutg@ikg.com",
	}

	for k, v := range student {
		fmt.Printf("K : %s ,  v :%s\n", k, v)
	}

}

func sliceString() {
	var buah = make([]string, 3)
	fmt.Println(buah)
	fmt.Printf("%v\n", buah)
	buah[0] = "durian"
	buah[1] = "mangis"
	buah[2] = "apple"
	fmt.Println(len(buah))
	fmt.Println(cap(buah))
	buah = append(buah, "nangka", "jambu")
	fmt.Println(buah)
	fmt.Println(len(buah))
	fmt.Println(cap(buah))

	buah2 := []string{"salak", "kedongdong", "kelapa"}
	newBuah := append(buah, buah2...)
	fmt.Println(newBuah)
	fmt.Println(len(newBuah))
	fmt.Println(cap(newBuah))

	newBuahs := copy(buah, buah2)
	fmt.Println(newBuahs)
	fmt.Println(len(buah))
	fmt.Println(cap(buah))
}
func sliceOk() {
	slice5 := []int{1, 2, 3, 4, 5}
	fmt.Println(slice5)
	fmt.Println(len(slice5))
	fmt.Println(cap(slice5))

	newSlice5 := slice5[1:3] // 2 ,3 value / length 2 , cap 4
	fmt.Println(newSlice5)
	fmt.Println(len(newSlice5))
	fmt.Println(cap(newSlice5))

	newSlice6 := append(newSlice5, 10)
	fmt.Println(newSlice6)
	fmt.Println(len(newSlice6))
	fmt.Println(cap(newSlice6))

	slice5x := []int{1, 2, 3, 4, 5}
	newSlice5a := append(slice5x, 10)
	fmt.Println(newSlice5a)
	fmt.Println(len(newSlice5a))
	fmt.Println(cap(newSlice5a))
}
func sliceNil() {
	var slice4 []int32
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
	if slice4 == nil {
		fmt.Println("Slice Nil")
	}

	slice5 := make([]int, 0)
	fmt.Println(len(slice5))
	fmt.Println(cap(slice5))
	if slice5 == nil {
		fmt.Println("Slice Nil")
	}
}

func slice1and3() {
	ints := make([]int, 5) //length 5 & capicity 5
	fmt.Println(len(ints))
	fmt.Println(cap(ints))
	slice2 := make([]int, 3, 5)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))
	//declare with index pos
	slice3 := []int{99: 88}
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))
	fmt.Println(slice3)
}

func arrayLogic() {
	fmt.Println("Array")

	var myArray1 [3]int
	fmt.Println(myArray1)
	var myArrStr [3]string
	fmt.Println(myArrStr)
	var myArr2 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(myArr2)
	var myArr3 = [...]int{1, 2, 3, 4, 5}
	fmt.Println(myArr3)
	var myArrStr2 = [3]string{
		"satu",
		"dua",
		"tiga",
	}
	fmt.Println(myArrStr2)
	var arr [3]int
	fmt.Println(arr)
	arr[0] = 1
	arr[1] = 10
	arr[2] = 20
	//arr[3] = 40
	fmt.Println(arr)

	//multidimensi
	angka := [3][4]int{
		{1, 2, 3, 4},
		{10, 20, 30, 40},
		{100, 200, 300, 400},
	}
	fmt.Println(angka)

	returnAngka := angka[2][3]
	fmt.Println(returnAngka)

	var arrx [3]int
	fmt.Println(arrx)
	fmt.Println("Len : ", len(arrx))
	fmt.Println("Cap : ", cap(arrx))

	var buah = []string{"Durian", "Semangka", "Nangka", "Pisang"}
	for idx, val := range buah {
		fmt.Printf("Indx : %d val %s", idx, val)
	}

	planets := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
	}
	for i := 0; i < len(planets); i++ {
		fmt.Println(" Name Planet : ", planets[i])
	}

	planetA := planets[0:4]
	fmt.Println(planetA)
	planetB := planets[4:]
	fmt.Println(planetB)

}

func loopingArrayString() {
	nameArr := []string{"Ali", "Topan", "Tono", "Budi"}

	for i := 0; i < len(nameArr); i++ {
		fmt.Println(" Name : ", nameArr[i])
	}
	for _, value := range nameArr {
		//fmt.Println("Index : ", idx)
		fmt.Println("Value : ", value)
	}
}

func loopingContinue() {
	for i := 1; i < 10; i++ {
		if i%2 == 1 {
			continue
		}
		if i > 8 {
			break
		}
		fmt.Println("Angka I : ", i)
	}
}
func loopingBreakexit() {

	sum := 0
exitloop:
	for {
		fmt.Printf("Looping %d\n", sum)
		if sum == 3 {
			fmt.Println("Value Sum 3 is break ")
			break
		}

		sumx := 0
		for {
			fmt.Println("Loop in loop Sum : %d  Sumx : %d ", sum, sumx)
			if sumx == 2 {
				fmt.Println("Sumx 2 is break exit 2nd loop")
				break exitloop
			}
			sumx++
		}

		sum++
	}
}

func loopingDoWhile() {
	value := 0
	for {
		fmt.Println("Value : ", value)
		if value == 5 {
			break
		}
		value++
	}
}

func loopingWhileCond() {
	x := 0
	for x < 3 {
		fmt.Println(" X : ", x)
		x++
	}
}

func looping1() {
	for i := 0; i < 10; i++ {
		fmt.Println("Looping ke : ", i)
	}
}

func condtionFallthroug() {
	var score = 70
	var optionValue = 100
	switch {
	case score == 100:
		fmt.Println("Nilai A")
	case score > 80 && score < 100:
		fmt.Println("Nilai B")
	case score > 70 && score < 80:
		fmt.Println("Nilai C")
		fallthrough
	case optionValue == 100:
		fmt.Println("Nilai B")
	default:
		fmt.Println("Nilai E")
	}
}
func condtionSwitchRelation() {
	var score = 100
	switch {
	case score == 100:
		fmt.Println("Nilai A")
	case score > 80 && score < 100:
		fmt.Println("Nilai B")
	default:
		fmt.Println("Nilai C")
	}
}

func condtionSwitchCase() {
	categori := "KARYAWAN"
	categori = strings.ToLower(categori)

	switch categori {
	case MAHASISWA:
		fmt.Println("Free Watter")
	case KARYAWAN:
		fmt.Println("Free Coffee")
	case DOSEN:
		fmt.Println("Free lunch")
	default:
		{
			fmt.Println("Anonymous")
			fmt.Println("Unknow Category")
		}

	}
}

func condtionIf() {
	nilaiuts := 70
	if nilaiuts > 60 {
		fmt.Println("Anda lulus")
		return
	}
	fmt.Println("Anda tidak lulus")
}

func condtion2() {
	var tahunsekarang = 2022

	if umur := tahunsekarang - 2010; umur > 17 {
		fmt.Println("Bisa mengajukan SIM")
		fmt.Println(umur)
	} else {
		fmt.Println("Belum bisa")
		fmt.Println(umur)
	}
}

func condition1() {
	ok := false
	if ok {
		fmt.Println("OK True ")
	} else {
		fmt.Println("Not OK false")
	}
}
