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
	r.Use(middleware.ContentTypeApplicationJsonMiddleware)
	handler.ProductRouter(r)
	r.HandleFunc("/", index)

	http.Handle("/", r)
	err := http.ListenAndServe(":81", r)
	if err != nil {
		log.Panic(err.Error())
	} else {
		log.Println("Server is running on :81")
	}

}
