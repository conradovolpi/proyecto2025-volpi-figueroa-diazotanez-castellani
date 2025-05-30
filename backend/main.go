package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		enableCors(w)
		fmt.Fprintln(w, "🚀 Backend Go corriendo correctamente!")
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		enableCors(w)
		if r.Method == "OPTIONS" {
			return
		}

		if r.Method != "POST" {
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
			return
		}

		type Credenciales struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		var creds Credenciales
		err := json.NewDecoder(r.Body).Decode(&creds)
		if err != nil {
			http.Error(w, "Error de formato", http.StatusBadRequest)
			return
		}

		if creds.Email == "admin@correo.com" && creds.Password == "admin" {
			user := map[string]interface{}{
				"id":     1,
				"nombre": "Administrador",
				"email":  creds.Email,
				"rol":    "admin",
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(user)
		} else if creds.Email == "user@correo.com" && creds.Password == "user" {
			user := map[string]interface{}{
				"id":     2,
				"nombre": "Usuario Test",
				"email":  creds.Email,
				"rol":    "user",
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(user)
		} else {
			http.Error(w, "Credenciales inválidas", http.StatusUnauthorized)
		}
	})

	

	fmt.Println("🔌 Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func enableCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}
