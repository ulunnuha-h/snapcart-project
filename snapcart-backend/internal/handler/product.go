package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ulunnuha-h/snapcart/internal/config"
	"github.com/ulunnuha-h/snapcart/internal/model"
)

func getAllProduct(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	var products []model.Product
	config.DB.Find(&products)

	response := config.GetResponse(products, true, "Successfully get all products!")
	json, err := json.Marshal(response)
	if err != nil {
		return
	}

	w.Write(json)
}

func addProduct(w http.ResponseWriter, r *http.Request){
	fmt.Println(r.Form)

}

func ProductRouter(r *mux.Router) {
	s := r.PathPrefix("/products").Subrouter()
	s.HandleFunc("", getAllProduct).Methods("GET")
	s.HandleFunc("", addProduct).Methods("POST")
}
