package model

import "fmt"

type Product struct {
	ID     uint
	Title  string
	Price  uint
	Rating float32
	Image  string
}

func Test() {
	fmt.Print("Print from Product struct")
}
