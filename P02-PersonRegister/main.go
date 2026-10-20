package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func registerPersonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(Response{
			Success: false,
			Message: "Method tidak diizinkan, gunakan POST",
		})
		return
	}

	var dto PersonDto
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Success: false,
			Message: "Format JSON request tidak valid",
		})
		return
	}

	person, err := NewPerson(dto.RawName, dto.RawAge)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Response{
		Success: true,
		Message: "Person berhasil didaftarkan!",
		Data: map[string]any{
			"id":   person.ID(),
			"name": person.Name(),
			"age":  person.Age(),
		},
	})
}

func main() {
	http.HandleFunc("/register", registerPersonHandler)

	port := ":8080"
	fmt.Printf("http://localhost%s\n", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Println("Server error:", err)
	}
}
