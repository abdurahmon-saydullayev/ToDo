package postgres

import (
	"GoProjects/ToDoList/storage"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Store struct {
	db      *sqlx.DB
	todo storage.ToDoRepoI
}

func NewPostgres(psqlConnString string) storage.StorageI {
	db, err := sqlx.Connect("postgres", psqlConnString)
	if err != nil {
		log.Panic(err)
	}

	return &Store{
		db: db,
	}
}

func (s *Store) ToDo() storage.ToDoRepoI {
	if s.todo == nil {
		s.todo = &todoRepo{db: s.db}
	}
	return s.todo
}

