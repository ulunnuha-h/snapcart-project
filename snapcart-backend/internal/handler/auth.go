package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/sirupsen/logrus"
	"github.com/ulunnuha-h/snapcart/internal/config"
	"github.com/ulunnuha-h/snapcart/internal/model"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func registerUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := schema.NewDecoder().Decode(&user, r.PostForm); err != nil {
		logrus.Errorf("Error decoding form data: %v", err)
		http.Error(w, "Error decoding form data", http.StatusBadRequest)
		return
	}

	hashPassword, err := HashPassword(user.Password)

	if err != nil{
		logrus.Errorf("Error hashing password: %v", err)
		http.Error(w, "Error hashing password", http.StatusBadRequest)
		return
	}

	user.Password = hashPassword

	if err := config.DB.Create(&user).Error; err != nil {
		logrus.Errorf("Error creating user in database: %v", err)
		http.Error(w, "Error creating user in database", http.StatusInternalServerError)
		return
	}

	response := config.GetResponse([]int{}, true, "Successfully registering new user!")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		logrus.Errorf("Error encoding JSON response: %v", err)
		http.Error(w, "Error processing JSON response", http.StatusInternalServerError)
	}
}

func AuthRouter(r *mux.Router) {
	s := r.PathPrefix("/auth").Subrouter()
	s.HandleFunc("/register", registerUser).Methods("POST")
}