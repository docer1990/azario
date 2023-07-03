package db

import (
	"database/sql"

	"github.com/docer1990/azario/models"
)

func (s *PostgressStore) createDoctorsTable() error {
	query := `create table if not exists doctors (
		doctor_id serial primary key,
		first_name varchar(100),
		last_name varchar(100),
		email VARCHAR(254) UNIQUE CHECK (email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$'),
		tel_number varchar(100),
		type varchar(100),
		created_at timestamp
	)`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgressStore) CreateDoctor(acc *models.Account) error {
	query := (`insert into doctors 
	(first_name, last_name, email, tel_number, type, created_at)
	values ($1, $2, $3, $4, $5)
	RETURNING id`)

	err := s.db.QueryRow(
		query,
		acc.FirstName,
		acc.LastName,
		acc.Email,
		acc.EncryptedPassword,
		acc.CreatedAt).Scan(&acc.ID)

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

func scanInToDoctors(rows *sql.Rows) (*models.Doctor, error) {
	doctor := new(models.Doctor)
	err := rows.Scan(
		&doctor.ID,
		&doctor.FirstName,
		&doctor.LastName,
		&doctor.Email,
		&doctor.TelNumber,
		&doctor.Type,
		&doctor.CreatedAt)

	return doctor, err

}
