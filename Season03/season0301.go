package main

import (
	"fmt"
	"time"
)

const (
	SENIN  = "Monday"
	SELASA = "Tuesday"
	RABU   = "Wednesday"
	KAMIS  = "Thursday"
	JUMAT  = "Friday"
	SABTU  = "Saturday"
	MINGGU = "Sunday"
)

const (
	DEFAULTDISC = 5.0
	SMALLDISC   = 10.25
	MEDIUMDISC  = 25.10
	HIGHDISC    = 45.25
)

func main() {
	//cetakHello()

	//fungsi1()
	//cetak("Budi", "Wati", "Iwan", "Joko", "Jhon", "Jono")
	resStr := stringBuilder("Budi", "Wati", "Iwan", "Joko", "Jhon", "Jono")
	fmt.Println(resStr)
}

func stringBuilder(names ...string) []map[string]string {
	var response []map[string]string
	for i, val := range names {
		key := fmt.Sprintf("Name_%02d", i)
		value := val
		tempValue := map[string]string{
			key: value,
		}
		response = append(response, tempValue)
	}

	return response
}

func cetak(names ...string) {
	for idx, value := range names {
		fmt.Printf("Idx : %d , value : %s\n", idx, value)
	}
}

func fungsi1() {
	var amt int32 = 10000
	nominalDisc, total := hitungTotal(amt)
	cetakHarga(amt, nominalDisc, total)
}
func hitungTotal(amt int32) (nominalDisc int32, total int32) {
	nominalDisc = getDiscNominal(amt)
	total = amt - nominalDisc
	return nominalDisc, total
}

func getDiscNominal(amt int32) int32 {
	return amt * int32(getDiscount()) / 100
}

func getDiscount() float64 {
	discount := DEFAULTDISC
	hari := time.Now().Weekday()

	switch hari.String() {
	case SENIN:
	case SELASA:
		fmt.Println("Dicount Hari Senin dan Selasa ")
		discount = HIGHDISC
	case RABU:
	case KAMIS:
	case JUMAT:
		fmt.Println("Dicount Hari Rabu, Kamis dan Jum'at ")
		discount = MEDIUMDISC
	case SABTU:
		fmt.Println("Dicount Hari Sabtu ")
		discount = SMALLDISC
	case MINGGU:
	default:
		fmt.Println("Default Discount")
	}
	return discount
}

/*
*
Fungsi dengan paramter
@harga
@nominalDisc
@totalAmt
*/
func cetakHarga(harga, nominalDisc, totalAmt int32) {
	fmt.Println("Amt : ", harga)
	fmt.Println("Disc : ", nominalDisc)
	fmt.Println("Total : ", totalAmt)
}

func cetakHello() {
	fmt.Println("Hello")
}
