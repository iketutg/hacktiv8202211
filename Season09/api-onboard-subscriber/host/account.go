package host

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
)

type Account struct {
	FullName    string
	Email       string
	PhoneNumber string
}

func SendCreateAccount(acc Account) {

	var url = "http://localhost:8085/account"

	client := resty.New()
	respo, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(acc).
		Post(url)

	if err == nil {
		if respo.StatusCode() == http.StatusCreated {
			fmt.Println("Create Account success")
			return
		}
	}
	fmt.Println("Create Account Failed ")
}
