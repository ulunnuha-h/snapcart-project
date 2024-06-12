package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/ulunnuha-h/snapcart/internal/handler"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome to snapcart API folks!")
}

func main() {
	r := mux.NewRouter()
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err.Error())
	}

	// db := config.OpenDatabase()

	r.HandleFunc("/product", handler.GetAllProduct).Methods("GET")
	r.HandleFunc("/", index)

	http.Handle("/", r)
	err := http.ListenAndServe(":81", r)
	if err != nil {
		log.Panic(err.Error())
	} else {
		fmt.Print("Server is running!")
	}

}
