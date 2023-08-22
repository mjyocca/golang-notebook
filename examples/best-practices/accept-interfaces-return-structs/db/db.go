package db

import "database/sql"

type Store struct {
	db *sql.DB
}

func NewDB() *Store {
	return &Store{
		db: &sql.DB{},
	}
}

func (s *Store) Insert(item interface{}) error {

	//... to-do
	return nil
}

func (s *Store) Get(id string, val interface{}) error {

	//... to-do
	return nil
}
