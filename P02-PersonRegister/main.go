package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type Server struct {
	manager *PersonManager
}

func NewServer(manager *PersonManager) *Server {
	return &Server{manager: manager}
}

func (s *Server) writeJSON(w http.ResponseWriter, statusCode int, payload Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(payload)
}

func (s *Server) RegisterPersonHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s.writeJSON(w, http.StatusMethodNotAllowed, Response{
			Success: false,
			Message: "Method tidak diizinkan, gunakan POST",
		})
		return
	}

	var dto PersonDto
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		s.writeJSON(w, http.StatusBadRequest, Response{
			Success: false,
			Message: "Format JSON request tidak valid",
		})
		return
	}

	person, err := s.manager.NewPerson(dto.RawName, dto.RawAge)
	if err != nil {
		statusCode := http.StatusBadRequest
		if errors.Is(err, ErrInvalidRange) {
			statusCode = http.StatusUnprocessableEntity
		}

		s.writeJSON(w, statusCode, Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	s.writeJSON(w, http.StatusCreated, Response{
		Success: true,
		Message: "Person berhasil didaftarkan!",
		Data: map[string]any{
			"id":   person.ID(),
			"name": person.Name(),
			"age":  person.Age(),
		},
	})
}

func (s *Server) DeletePersonHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		s.writeJSON(w, http.StatusMethodNotAllowed, Response{
			Success: false,
			Message: "Method tidak diizinkan, gunakan DELETE",
		})
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		s.writeJSON(w, http.StatusBadRequest, Response{
			Success: false,
			Message: "ID person tidak boleh kosong",
		})
		return
	}

	err := s.manager.DeletePerson(id)
	if err != nil {
		s.writeJSON(w, http.StatusBadRequest, Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	s.writeJSON(w, http.StatusOK, Response{
		Success: true,
		Message: "Person berhasil dihapus!",
	})
}

func main() {
	manager := NewPersonManager()
	server := NewServer(manager)

	mux := http.NewServeMux()
	mux.HandleFunc("/register", server.RegisterPersonHandler)
	mux.HandleFunc("/person", server.DeletePersonHandler)

	mux.Handle("/", http.FileServer(http.Dir("./wwwroot")))

	port := ":8080"
	fmt.Printf("Server running at http://localhost%s\n", port)

	if err := http.ListenAndServe(port, mux); err != nil {
		fmt.Println("Server error:", err)
	}
}