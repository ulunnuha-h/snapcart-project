package model

import "fmt"

type Product struct {
	ID     uint		`json:"id"`
	Title  string	`json:"title"`
	Price  uint		`json:"price"`
	Rating float32	`json:"rating"`
	Image  string	`json:"image"`
}

func Test() {
	fmt.Print("Print from Product struct")
}
