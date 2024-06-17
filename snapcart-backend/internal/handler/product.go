package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/ulunnuha-h/snapcart/internal/config"
	"github.com/ulunnuha-h/snapcart/internal/model"
)

var decoder = schema.NewDecoder()

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
	r.PostFormValue("")

	var product model.Product
	err := decoder.Decode(&product, r.PostForm)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	err = config.DB.Create(&product).Error
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	response := config.GetResponse([]model.Product{product}, true, "Successfully add a product!")
	json, err := json.Marshal(response)
	if err != nil {
		return
	}

	w.Write(json)
}

func ProductRouter(r *mux.Router) {
	s := r.PathPrefix("/products").Subrouter()
	s.HandleFunc("", getAllProduct).Methods("GET")
	s.HandleFunc("", addProduct).Methods("POST")
}
