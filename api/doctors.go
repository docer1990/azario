package api

import (
	"fmt"
	"net/http"
)

func (s *APIServer) handleDoctors(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetDoctors(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleGetDoctors(w http.ResponseWriter, r *http.Request) error {
	doctors, err := s.store.GetDoctors()
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, doctors)
}
