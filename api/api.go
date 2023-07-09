package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/docer1990/azario/db"
	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddr string
	store      db.Storage
}

func NewAPIServer(listenAddr string, store db.Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	// user routers
	router.HandleFunc("/account", withJWTAuth(makeHTTPHandleFunc(s.handleAccount), s.store, true))
	router.HandleFunc("/account/{id}", withJWTAuth(makeHTTPHandleFunc(s.handleGetAccountByID), s.store, true))
	router.HandleFunc("/login", makeHTTPHandleFunc(s.handleLogin))

	// doctors routers
	router.HandleFunc("/doctor", withJWTAuth(makeHTTPHandleFunc(s.handleDoctors), s.store, false))
	router.HandleFunc("/doctor/{id}", withJWTAuth(makeHTTPHandleFunc(s.handleGetDoctorByID), s.store, false))

	log.Println("JSON API server running on port: ", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")

	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}

func permissionDenied(w http.ResponseWriter) {
	WriteJSON(w, http.StatusForbidden, ApiError{Error: "permission denied"})
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			// handle the error
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}
