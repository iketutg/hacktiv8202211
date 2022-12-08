package main

import "fmt"

func main() {
	channel := make(chan string)

	go Cetak("Hallo ", channel)
	go Cetak("Hallo 2", channel)
	result := <-channel
	fmt.Println("Result : ", result)
	result2, ok := <-channel
	if ok {
		fmt.Println("Result2 : ", result2)
	} else {
		fmt.Println("Channel is close")
	}
	channelPartner := make(chan string)
	go getPartnerNameById("123456", channelPartner)
	channelProduct := make(chan string)
	go getProductNameById("0000001", channelProduct)

	var productName string
	productName = <-channelProduct
	partnerName := <-channelPartner

	fmt.Println(" Product Name : ", productName)
	fmt.Println(" Partner Name : ", partnerName)

}

func getPartnerNameById(idPartner string, channel chan string) {
	// ......
	fmt.Println("idPartner : ", idPartner)
	channel <- fmt.Sprintf("PT . ABCDEF")
}

func getProductNameById(idpelanggan string, channel chan string) {
	// ......
	fmt.Println("idpelanggan : ", idpelanggan)
	channel <- fmt.Sprintf("Laptop")
}

func Cetak(data string, channel chan string) {
	channel <- fmt.Sprintf("Data dari func cetak %s", data)
}
