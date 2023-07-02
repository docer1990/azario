package db

import (
	"database/sql"
	"log"

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
}

type PostgressStore struct {
	db *sql.DB
}

func NewPostgressStore() (*PostgressStore, error) {
	connStr := "postgresql://postgres:19DoCerv90@localhost/zarino?sslmode=disable"
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
	return s.createAccountTable()
}
