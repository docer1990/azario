package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/docer1990/azario/models"
	"github.com/gorilla/mux"
)

func (s *APIServer) handleDoctors(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetDoctors(w, r)
	}

	if r.Method == "POST" {
		return s.handleCreateDoctor(w, r)
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

func (s *APIServer) handleCreateDoctor(w http.ResponseWriter, r *http.Request) error {
	createDoctorReq := new(models.CreateDoctorRequest)
	if err := json.NewDecoder(r.Body).Decode(createDoctorReq); err != nil {
		return err
	}

	doctor, err := models.NewDoctor(createDoctorReq.FirstName, createDoctorReq.LastName, createDoctorReq.Email, createDoctorReq.TelNumber, createDoctorReq.DocType)
	if err != nil {
		return err
	}

	if err := s.store.CreateDoctor(doctor); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, doctor)
}

func (s *APIServer) handleGetDoctorByID(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		id, err := getDoctorID(r)
		fmt.Println(id)
		if err != nil {
			return err
		}

		doctor, err := s.store.GetDoctorById(id)
		if err != nil {
			return err
		}

		return WriteJSON(w, http.StatusOK, doctor)
	}

	if r.Method == "DELETE" {
		return s.handleDeleteDoctor(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleDeleteDoctor(w http.ResponseWriter, r *http.Request) error {
	id, err := getDoctorID(r)
	if err != nil {
		return err
	}

	if err := s.store.DeleteDoctor(id); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, map[string]int{"deleted": id})
}

func getDoctorID(r *http.Request) (int, error) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return id, fmt.Errorf("invalid id given %s", idStr)
	}

	return id, nil
}
