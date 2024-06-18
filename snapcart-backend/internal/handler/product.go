package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/sirupsen/logrus"
	"github.com/ulunnuha-h/snapcart/internal/config"
	"github.com/ulunnuha-h/snapcart/internal/model"
)

var decoder = schema.NewDecoder()

func getAllProduct(w http.ResponseWriter, r *http.Request) {
	var products []model.Product
	if err := config.DB.Find(&products).Error; err != nil {
		logrus.Errorf("Error fetching products: %v", err)
		http.Error(w, "Error fetching products from database", http.StatusInternalServerError)
		return
	}

	response := config.GetResponse(products, true, "Successfully get all products!")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		logrus.Errorf("Error encoding JSON response: %v", err)
		http.Error(w, "Error processing JSON response", http.StatusInternalServerError)
	}
}

func getProductById(w http.ResponseWriter, r *http.Request){
	id := mux.Vars(r)["id"]
	var product model.Product

	if err := config.DB.First(&product, id).Error; err != nil {
		logrus.Errorf("Error no data found: %v", err)
		http.Error(w, "Error no data found in database", http.StatusBadRequest)
		return
	}

	response := config.GetResponse([]model.Product{product}, true, "Successfully get a product!")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		logrus.Errorf("Error encoding JSON response: %v", err)
		http.Error(w, "Error processing JSON response", http.StatusInternalServerError)
	}
}

func addProduct(w http.ResponseWriter, r *http.Request) {
	var product model.Product
	if err := decoder.Decode(&product, r.PostForm); err != nil {
		logrus.Errorf("Error decoding form data: %v", err)
		http.Error(w, "Error decoding form data", http.StatusBadRequest)
		return
	}

	if err := config.DB.Create(&product).Error; err != nil {
		logrus.Errorf("Error creating product in database: %v", err)
		http.Error(w, "Error creating product in database", http.StatusInternalServerError)
		return
	}

	response := config.GetResponse([]model.Product{product}, true, "Successfully added a product!")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		logrus.Errorf("Error encoding JSON response: %v", err)
		http.Error(w, "Error processing JSON response", http.StatusInternalServerError)
	}
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var product model.Product

	if err := config.DB.First(&product, id).Error; err != nil {
		logrus.Errorf("Error no data found: %v", err)
		http.Error(w, "Error no data found in database", http.StatusBadRequest)
		return
	}

	if err := config.DB.Delete(&model.Product{}, id).Error; err != nil {
		logrus.Errorf("Error deleting product from database: %v", err)
		http.Error(w, "Error deleting product from database", http.StatusInternalServerError)
		return
	}

	response := config.GetResponse([]int{}, true, "Successfully deleted a product!")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		logrus.Errorf("Error encoding JSON response: %v", err)
		http.Error(w, "Error processing JSON response", http.StatusInternalServerError)
	}
}


func updateProduct(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var product model.Product

	if err := config.DB.First(&product, id).Error; err != nil {
		logrus.Errorf("Error no data found: %v", err)
		http.Error(w, "Error no data found in database", http.StatusBadRequest)
		return
	}

	if err := decoder.Decode(&product, r.PostForm); err != nil {
		logrus.Errorf("Error decoding form data: %v", err)
		http.Error(w, "Error decoding form data", http.StatusBadRequest)
		return
	}

	if err := config.DB.Model(&model.Product{}).Where("id = ?", id).Updates(&product).Error; err != nil {
		logrus.Errorf("Error updating product in database: %v", err)
		http.Error(w, "Error updating product in database", http.StatusInternalServerError)
		return
	}

	response := config.GetResponse([]model.Product{product}, true, "Successfully updated a product!")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		logrus.Errorf("Error encoding JSON response: %v", err)
		http.Error(w, "Error processing JSON response", http.StatusInternalServerError)
	}
}

func ProductRouter(r *mux.Router) {
	s := r.PathPrefix("/products").Subrouter()
	s.HandleFunc("", getAllProduct).Methods("GET")
	s.HandleFunc("", addProduct).Methods("POST")
	s.HandleFunc("/{id}", deleteProduct).Methods("DELETE")
	s.HandleFunc("/{id}", updateProduct).Methods("PUT")
	s.HandleFunc("/{id}", getProductById).Methods("GET")
}
