package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)
func NewAPIServer(listenAddr string, store Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store: store,
	}
}

// Run starts the API server.
func (s *APIServer) Run() error {
	router := mux.NewRouter()

	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))
	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.handleGetAccount))

	log.Printf("API server is running on %s", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)

	return nil
}

// handleAccount handles the account creation.
func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetAccount(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateAccount(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteAccount(w, r)
	}
	return fmt.Errorf("method not allowed %s", r.Method)
}

// handleGetAccount handles the retrieval of an account.
func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	// account := NewAccount("John", "Doe")
	id := mux.Vars(r)["id"]

	fmt.Printf("Getting account with id %s\n", id)

	idInt, err := strconv.Atoi(id)
	if err != nil {
		// handle the error
	}
	return WriteJSON(w, http.StatusOK, &Account{ID: idInt})
}

// handleCreateAccount handles the creation of an account.
func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// handleDeleteAccount handles the deletion of an account.
func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// handleTransfer handles the transfer of funds between accounts.
func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

type ApiError struct {
	Error string 
}

// apiFunc is a function type that handles API requests.
type apiFunc func(w http.ResponseWriter, r *http.Request) error

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusInternalServerError, ApiError{Error: err.Error()})
		} 
	}
}

// APIServer represents the API server.
type APIServer struct {
	listenAddr string
	store Storage
	// We will add Postgres connection here
}
