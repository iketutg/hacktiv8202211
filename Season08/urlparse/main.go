package main

import (
	"fmt"
	"net/url"
)

func main() {
	var urlstr = "ipfs://iketutg.com/notifcation?phone=081213456778&message=hello golang"
	parse, err := url.Parse(urlstr)
	if err != nil {
		fmt.Println("Parsing Error : ", err.Error())
	}

	fmt.Println(" Parse : ", parse)
	fmt.Println(" Host : ", parse.Host)
	fmt.Println(" Path : ", parse.Path)
	fmt.Println(" Port : ", parse.Port())
	fmt.Println(" Scheme : ", parse.Scheme)
	fmt.Println(" Phone : ", parse.Query()["phone"][0])
	fmt.Println(" Message : ", parse.Query()["message"][0])

}
