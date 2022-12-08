package main

import "fmt"

type MahlukHidup interface {
	Bersuara() string
}

type Kucing struct {
}

type Bebek struct {
}

type Developer struct {
}

func (b Bebek) Bersuara() string {
	return "WK .... Wk.... Wk..."
}

func (k Kucing) Bersuara() string {
	return "Meooooowwwww !"
}

func (d Developer) Bersuara() string {
	return "Design Patern !"
}

func main() {
	fmt.Println("Interface with slice struct")
	hidups := []MahlukHidup{new(Bebek), new(Developer), Kucing{}}

	for _, mahluk := range hidups {
		fmt.Println(mahluk.Bersuara())
	}
}
