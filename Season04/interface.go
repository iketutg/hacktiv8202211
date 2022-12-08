package main

import "fmt"

type Duck interface {
	Quack()
	Walk()
}

type Donald struct {
	GatotKaca
}

type GatotKaca struct {
}

func (g GatotKaca) Quack() {
	fmt.Println("HAI Saya Gatot Kaca Tulang kawat!")
}

func (g GatotKaca) Walk() {
	fmt.Println("Saya berjalan dengan tegak")
}

func (d Donald) Quack() {
	fmt.Println("Hi I'm Donald Duck!")
}

func (d Donald) Walk() {
	fmt.Println("I am waddle")
}

func behaveLikeDonald(dc Duck) {
	dc.Quack()
	dc.Walk()
}

//
//func behaveLikeDonal(d Donald) {
//	d.Quack()
//	d.Walk()
//}
//
//func behaveLikeGatot(g GatotKaca) {
//	g.Quack()
//	g.walk()
//}

func main() {
	donald := new(Donald)
	gatot := new(GatotKaca)
	behaveLikeDonald(donald)
	behaveLikeDonald(gatot)
	//donald := Donald{}
	//behaveLikeDonal(donald)
	//gatotkaca := GatotKaca{} //new(GatotKaca)
	//behaveLikeGatot(gatotkaca)
}
