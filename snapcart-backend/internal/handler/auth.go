package handler

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
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

	if err := config.DB.First(&user).Where("email = ?", user.Email).Error; err == nil {
		logrus.Errorf("Error email have been used: %v", err)
		http.Error(w, "Error email have been used", http.StatusBadRequest)
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

func loginUser(w http.ResponseWriter, r *http.Request) {
    var creds model.Credential
	var user model.User

	if err := schema.NewDecoder().Decode(&creds, r.PostForm); err != nil {
		logrus.Errorf("Error decoding form data: %v", err)
		http.Error(w, "Error decoding form data", http.StatusBadRequest)
		return
	}

	err := config.DB.Where("email = ?", creds.Email).First(&user).Error;

	if err != nil || !CheckPasswordHash(creds.Password, user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		logrus.Errorf("Error email or password not found: %v", err)
		http.Error(w, "Error email or password not found", http.StatusBadRequest)
		return
	}

    expirationTime := time.Now().Add(15 * time.Minute)
    claims := &model.Claim{
		ID: user.ID ,
        Email: user.Email,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtKey := []byte(os.Getenv("JWT_KEY"))
    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
		logrus.Errorf("Error: %v", err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

	if os.Getenv("APP_ENV") == "local" {
		http.SetCookie(w, &http.Cookie{
			Name:    "jwtToken",
			Value:   tokenString,
			Expires: expirationTime,
			Path: "/",
			HttpOnly: true,
		})	
	} else {
		http.SetCookie(w, &http.Cookie{
			Name:    "jwtToken",
			Value:   tokenString,
			Expires: expirationTime,
			Path: "/",
			HttpOnly: true,
			Secure: true,  // Set to true in production
		})
	}
}


func AuthRouter(r *mux.Router) {
	s := r.PathPrefix("/auth").Subrouter()
	s.HandleFunc("/register", registerUser).Methods("POST")
	s.HandleFunc("/login", loginUser).Methods("POST")
}