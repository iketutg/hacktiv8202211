package host

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
)

type User struct {
	UserName string
	Password string
	Email    string
}

func SendCreateUser(user User) {

	var url = "http://localhost:8087/user"

	client := resty.New()
	respo, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(user).
		Post(url)

	if err == nil {
		if respo.StatusCode() == http.StatusCreated {
			fmt.Println("Create Account success")
			return
		}
	}
	fmt.Println("Create Account Failed ")
}
