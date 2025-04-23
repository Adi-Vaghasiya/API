// package main

// import (
// 	"encoding/json"
// 	"net/http"
// )

// type Credentials struct {
// 	Username string `json:"username"`
// 	Passwoed string `json:"password"`
// }

// func LoginHandler(w http.ResponseWriter, r *http.Request) {
// 	var creds Credentials
// 	_ = json.NewDecoder(r.Body).Decode(&creds)

// 	if creds.Username != "admin" || creds.Passwoed != "pass" {
// 		http.Error(w, "Unauthorized", http.StatusUnauthorized)
// 		return
// 	}

//		tokenStr, err := GenerateJWT(creds.Username)
//		if err != nil {
//			http.Error(w, "Error Generating Token", http.StatusInternalServerError)
//			return
//		}
//		json.NewDecoder(w).Encode(map[string]string{
//			"token": tokenString,
//		})
//	}
package handler

import (
	"encoding/json"
	"jwt-api/middleware"
	"jwt-api/utils"
	"net/http"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	var creds Credentials
	_ = json.NewDecoder(r.Body).Decode(&creds)

	// Dummy authentication
	if creds.Username != "admin" || creds.Password != "pass" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	tokenString, err := utils.GenerateJWT(creds.Username)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"token": tokenString,
	})
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the ProfilePage\n"))
	value := r.Context().Value(middleware.UserContextKey)
	user, ok := value.(string)
	if !ok || user == "" {
		http.Error(w, "Unauthorized - user not found in context", http.StatusUnauthorized)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Welcome, " + user,
	})
}
