package model

type Product struct {
	ID     uint    `json:"id"`
	Title  string  `json:"title"`
	Price  uint    `json:"price"`
	Rating float32 `json:"rating"`
	Image  string  `json:"image"`
}
