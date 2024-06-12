package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ulunnuha-h/snapcart/internal/config"
	"github.com/ulunnuha-h/snapcart/internal/model"
)

func GetAllProduct(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	db := config.OpenDatabase()
	var product []model.Product
	db.Find(&product)

	response := config.GetResponse[model.Product](product)
	json, err := json.Marshal(response)
	if err != nil {
		return
	}

	w.Write(json)
}

func Test() {
	fmt.Print("Print from Product handlers")
}
