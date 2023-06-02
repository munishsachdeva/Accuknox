package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Note struct {
	ID   uint32 `json:"id"`
	Note string `json:"note"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type CreateNoteRequest struct {
	SID  string `json:"sid"`
	Note string `json:"note"`
}

type DeleteNoteRequest struct {
	SID string `json:"sid"`
	ID  uint32 `json:"id"`
}

var notes []Note

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Process the user creation logic and return the appropriate response
	// ...

	w.WriteHeader(http.StatusOK)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Process the login logic and return the appropriate response
	// ...

	sessionID := "example-session-id"
	response := map[string]string{
		"sid": sessionID,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func listNotesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	sid := r.URL.Query().Get("sid")
	if sid == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Verify the session ID and retrieve the user's notes
	// ...

	// For demonstration purposes, let's assume the notes are pre-populated
	notes := []Note{
		{ID: 1, Note: "Note 1"},
		{ID: 2, Note: "Note 2"},
		{ID: 3, Note: "Note 3"},
	}

	response := map[string][]Note{
		"notes": notes,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func createNoteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req CreateNoteRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Verify the session ID and create a new note
	// ...

	newNoteID := uint32(1) // Example ID of the newly created note

	response := map[string]uint32{
		"id": newNoteID,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func deleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req DeleteNoteRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Verify the session ID and delete the note
	// ...

	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/signup", createUserHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/notes", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			listNotesHandler(w, r)
		case http.MethodPost:
			createNoteHandler(w, r)
		case http.MethodDelete:
			deleteNoteHandler(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Starting server on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
