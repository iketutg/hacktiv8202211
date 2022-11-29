package main

import "fmt"

type member struct {
	userId  string
	name    string
	balance int
}

func (m member) cetakBalance() {
	fmt.Println("Balance : ", m.balance)
}

func (m member) getName() string {
	return m.name
}

func (m *member) updateBalance() {
	m.balance = 1000000
}

func main() {
	var cust = member{"iketutg", "Iketut Gunawan", 10000}
	cust.cetakBalance()

	name := cust.getName()
	fmt.Println("Name : ", name)
	cust.updateBalance()

	fmt.Println(" Balance : ", cust.balance)
}
