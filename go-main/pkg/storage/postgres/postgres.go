package postgres

import (
	"GoNews/pkg/storage"
	"database/sql"
	"errors"
)

type Store struct {
	db *sql.DB
}

func New(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) Posts() ([]storage.Post, error) {
	return nil, errors.New("not implemented yet")
}

func (s *Store) AddPost(post storage.Post) error {
	return errors.New("not implemented yet")
}

func (s *Store) UpdatePost(post storage.Post) error {
	return errors.New("not implemented yet")
}

func (s *Store) DeletePost(post storage.Post) error {
	return errors.New("not implemented yet")
}
