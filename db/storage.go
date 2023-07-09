package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/docer1990/azario/models"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*models.Account) error
	DeleteAccount(int) error
	UpdateAccount(*models.Account) error
	GetAccounts() ([]*models.Account, error)
	GetAccountById(int) (*models.Account, error)
	GetAccountByEmail(string) (*models.Account, error)
	GetDoctors() ([]*models.Doctor, error)
	CreateDoctor(*models.Doctor) error
	GetDoctorById(int) (*models.Doctor, error)
	DeleteDoctor(int) error
}

type PostgressStore struct {
	db *sql.DB
}

func NewPostgressStore() (*PostgressStore, error) {
	connStr := os.Getenv("DB_HOST")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgressStore{
		db: db,
	}, nil
}

func (s *PostgressStore) Init() error {
	if err := s.createAccountTable(); err != nil {
		return err
	}

	if err := s.createDoctorsTable(); err != nil {
		return err
	}

	return nil
}
