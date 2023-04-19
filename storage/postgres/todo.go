package postgres

import (
	"GoProjects/ToDoList/models"
	"GoProjects/ToDoList/storage"

	"github.com/jmoiron/sqlx"
)

type todoRepo struct {
	db *sqlx.DB
}

func NewToDoRepo(db *sqlx.DB) storage.ToDoRepoI {
	return todoRepo{
		db: db,
	}
}

func (r todoRepo) Create(entity models.Create) (err error) {
	insertQuery := `INSERT INTO todo (
		title,
		description
	) VALUES (
		$1,
		$2
	)`

	_, err = r.db.Exec(insertQuery,
		entity.Title,
		entity.Description,
	)

	return err
}

// func (r todoRepo) GetList(entity models.Update) (err error){

// }