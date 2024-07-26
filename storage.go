package main

import (
	"database/sql"
	// "log"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(account *Account) error
	GetAccountId(id string) (*Account, error)
	DeleteAccount(int) error
	UpdateAccount(account *Account) error
}

type PostgreStore struct {
	db *sql.DB
}

func NewPostgreStore() (* PostgreStore, error) {
	connstr := "user=postgres dbname=postgres password=gobank sslmode=verify-full"
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &PostgreStore{
		db: db,
	}, nil
}

