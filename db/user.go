package db

import (
	"database/sql"
	"fmt"

	"github.com/docer1990/azario/models"
)

func (s *PostgressStore) createAccountTable() error {
	query := `create table if not exists account (
		id serial primary key,
		first_name varchar(100),
		last_name varchar(100),
		email VARCHAR(254) UNIQUE CHECK (email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$'),
		number serial,
		encrypted_password varchar(100),
		balance serial,
		created_at timestamp
	)`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgressStore) CreateAccount(acc *models.Account) error {
	query := (`insert into account 
	(first_name, last_name, email, number, encrypted_password, balance, created_at)
	values ($1, $2, $3, $4, $5, $6, $7)
	RETURNING id`)

	resp, err := s.db.Query(
		query,
		acc.FirstName,
		acc.LastName,
		acc.Email,
		acc.Number,
		acc.EncryptedPassword,
		acc.Balance,
		acc.CreatedAt)

	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", resp)

	return nil
}

func (s *PostgressStore) UpdateAccount(*models.Account) error {
	return nil
}

func (s *PostgressStore) DeleteAccount(id int) error {
	_, err := s.db.Query("DELETE FROM account WHERE id = $1", id)
	return err
}

func (s *PostgressStore) GetAccountByEmail(email string) (*models.Account, error) {
	rows, err := s.db.Query("select * from account where email = $1", email)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanInToAccount(rows)
	}

	return nil, fmt.Errorf("account with number [%v] not found", email)
}

func (s *PostgressStore) GetAccountById(id int) (*models.Account, error) {
	rows, err := s.db.Query("SELECT * FROM account WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanInToAccount(rows)
	}

	return nil, fmt.Errorf("account %d not found", id)
}

func (s *PostgressStore) GetAccounts() ([]*models.Account, error) {
	rows, err := s.db.Query("SELECT * FROM account")
	if err != nil {
		return nil, err
	}

	accounts := []*models.Account{}
	for rows.Next() {
		account, err := scanInToAccount(rows)

		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}

func scanInToAccount(rows *sql.Rows) (*models.Account, error) {
	account := new(models.Account)
	err := rows.Scan(
		&account.ID,
		&account.FirstName,
		&account.LastName,
		&account.Email,
		&account.Number,
		&account.EncryptedPassword,
		&account.Balance,
		&account.CreatedAt)

	return account, err
}
