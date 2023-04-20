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

func (r todoRepo) GetByID(id string) (models.ToDo, error) {
	var a models.ToDo

	err := r.db.QueryRow(`SELECT 
		id, 
		title, 
		description,
		created_at, 
		updated_at,
		completed
	FROM todo where id = $1
	`, id).Scan(
		&a.ID,
		&a.Title,
		&a.Description,
		&a.Created,
		&a.Updated,
		&a.Completed,
	)

	if err != nil {
		return a, err
	}
	return a, nil
}
