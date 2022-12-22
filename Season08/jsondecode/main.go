package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	FullName string `json:"full-name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

func contoh1() {
	var strjson = ` 
				{
             		"full-name" : "IKetut Gunawan",
					"email":"iketutg@email.com", 
					"age": 43
				}
       		`
	fmt.Println(strjson)
	var emp Employee
	err := json.Unmarshal([]byte(strjson), &emp)
	if err != nil {
		fmt.Println("Json unmarshal error : ", err.Error())
	}

	fmt.Println(" Employee Struct : ", emp.FullName)
	fmt.Println(" Employee Struct : ", emp.Age)
	fmt.Println(" Employee Struct : ", emp.Email)
}

type Signer struct {
	APIKey string `json:"apiKey"`
	Email  string `json:"email"`
	EncCVV string `json:"encCVV"`
	KeyID  string `json:"keyId"`
}

type MyJsonName struct {
	PartnerTrxID string `json:"partnerTrxId"`
	//Signer       struct {
	//	APIKey string `json:"apiKey"`
	//	Email  string `json:"email"`
	//	EncCVV string `json:"encCVV"`
	//	KeyID  string `json:"keyId"`
	//} `json:"signer"`
	Signers Signer `json:"signer"`
}

func contoh2() {
	var jsonstr = `
					{
						"partnerTrxId": "1234567890",
						"signer": {
							"email": "iketut.gunawan@vida.id",
							"keyId": "4b6cbf22-5122-46e3-9a0c-fbaf83560ec9",
							"apiKey":"DMaDL4LRD3pqQyyH",
							"encCVV": "h99KMTiC4jEhi5G-9g=="
							}
					}
				`

	var myjson MyJsonName
	err := json.Unmarshal([]byte(jsonstr), &myjson)
	if err != nil {
		fmt.Println("Json unmarshal error : ", err.Error())
	}

	fmt.Println(" Employee Struct : ", myjson.PartnerTrxID)
	fmt.Println(" Employee Struct : ", myjson.Signers.APIKey)

}

func contoh3() {
	var strjsonarr = `[
					{
             		"full-name" : "IKetut Gunawan",
					"email":"iketutg@email.com", 
					"age": 43
					},
					{
             		"full-name" : "ibrahim",
					"email":"ibrahim@email.com", 
					"age": 16
					}
				]`

	var empArr []Employee

	err := json.Unmarshal([]byte(strjsonarr), &empArr)
	if err != nil {
		fmt.Println("Json unmarshal error : ", err.Error())
	}

	for _, emp := range empArr {
		fmt.Println(" Employee Struct : ", emp.FullName)
		fmt.Println(" Employee Struct : ", emp.Age)
		fmt.Println(" Employee Struct : ", emp.Email)
	}

}

func contoh4() {
	var strjson = ` 
				{
             		"full-name" : "IKetut Gunawan",
					"email":"iketutg@email.com", 
					"age": 43
				}
       		`
	fmt.Println(strjson)
	var emp map[string]interface{}
	err := json.Unmarshal([]byte(strjson), &emp)
	if err != nil {
		fmt.Println("Json unmarshal error : ", err.Error())
	}

	fmt.Println(" Employee Struct : ", emp["full-name"])
	fmt.Println(" Employee Struct : ", emp["age"])
	fmt.Println(" Employee Struct : ", emp["email"])
}

func main() {
	//contoh1() simplejson
	//contoh2()
	//contoh3() jsonarray
	contoh4()
}
