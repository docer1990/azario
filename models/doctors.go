package models

import "time"

type CreateDoctorRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	TelNumber string `json:"telNumber"`
	DocType   string `json:"docType"`
}

type Doctor struct {
	DoctorID  int64     `json:"doctorId"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	TelNumber string    `json:"telNumber"`
	DocType   string    `json:"docType"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewDoctor(firstName, lastName, email, telNumber, docType string) (*Doctor, error) {

	return &Doctor{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		TelNumber: telNumber,
		DocType:   docType,
		CreatedAt: time.Now().UTC(),
	}, nil
}
