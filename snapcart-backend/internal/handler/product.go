package handler

import (
	"fmt"
	"net/http"
)

func GetAllProduct(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	// products := []model.Product{
	// 	{ID: 1, Title: "The Blue Shoes", Price: 20000, Rating: 2.3, Image: ""},
	// }

	// json, err := json.Marshal(response)
	// if err != nil {
	// 	return
	// }

	w.Write(json)
}

func Test() {
	fmt.Print("Print from Product handlers")
}
