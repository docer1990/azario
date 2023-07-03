package models

import "time"

type Doctor struct {
	ID        int64     `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	TelNumber string    `json:"telNumber"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"createdAt"`
}
