package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/ulunnuha-h/snapcart/internal/config"
	"github.com/ulunnuha-h/snapcart/internal/handler"
	"github.com/ulunnuha-h/snapcart/internal/middleware"
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

	// Open Database
	config.OpenDatabase()

	// Register subrouter
	r.Use(middleware.BasicSetup)
	handler.ProductRouter(r)
	handler.AuthRouter(r)
	r.HandleFunc("/", index)

	http.Handle("/", r)

	log.Println("Server is running on :81")
	err := http.ListenAndServe(":81", r)
	if err != nil {
		log.Panic(err.Error())
	}
}
