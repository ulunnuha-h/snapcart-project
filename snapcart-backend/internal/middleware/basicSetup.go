package middleware

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func BasicSetup(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set Content-Type Header to applilcation/json
		w.Header().Set("Content-Type", "application/json")

		// Parse form data if method POST or PUT
		if r.Method == "POST" || r.Method == "PUT" {
			if err := r.ParseForm(); err != nil {
				logrus.Errorf("Error parsing form: %v", err)
				http.Error(w, "Invalid form data", http.StatusBadRequest)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}