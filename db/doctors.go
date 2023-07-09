package db

import (
	"database/sql"
	"fmt"

	"github.com/docer1990/azario/models"
)

func (s *PostgressStore) createDoctorsTable() error {
	query := `create table if not exists doctors (
		doctor_id serial primary key,
		first_name varchar(100),
		last_name varchar(100),
		email VARCHAR(254) UNIQUE CHECK (email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$'),
		tel_number varchar(100),
		doc_type varchar(100),
		created_at timestamp
	)`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgressStore) CreateDoctor(doc *models.Doctor) error {
	query := (`insert into doctors 
	(first_name, last_name, email, tel_number, doc_type, created_at)
	values ($1, $2, $3, $4, $5, $6)
	RETURNING doctor_id`)

	err := s.db.QueryRow(
		query,
		doc.FirstName,
		doc.LastName,
		doc.Email,
		doc.TelNumber,
		doc.DocType,
		doc.CreatedAt).Scan(&doc.DoctorID)

	if err != nil {
		return err
	}

	return nil
}

func (s *PostgressStore) GetDoctors() ([]*models.Doctor, error) {
	rows, err := s.db.Query("SELECT * FROM doctors")
	if err != nil {
		return nil, err
	}

	doctors := []*models.Doctor{}
	for rows.Next() {
		doctor, err := scanInToDoctors(rows)

		if err != nil {
			return nil, err
		}

		doctors = append(doctors, doctor)
	}

	return doctors, nil
}

func (s *PostgressStore) GetDoctorById(id int) (*models.Doctor, error) {
	rows, err := s.db.Query(`SELECT * FROM doctors WHERE doctor_id = $1`, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanInToDoctors(rows)
	}

	return nil, fmt.Errorf("doctor_id %d not found", id)
}

func (s *PostgressStore) DeleteDoctor(id int) error {
	_, err := s.db.Query("DELETE FROM doctors WHERE doctor_id = $1", id)
	return err
}

func scanInToDoctors(rows *sql.Rows) (*models.Doctor, error) {
	doctor := new(models.Doctor)
	err := rows.Scan(
		&doctor.DoctorID,
		&doctor.FirstName,
		&doctor.LastName,
		&doctor.Email,
		&doctor.TelNumber,
		&doctor.DocType,
		&doctor.CreatedAt)

	return doctor, err

}
