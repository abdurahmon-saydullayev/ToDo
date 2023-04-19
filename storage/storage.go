package storage

import "GoProjects/ToDoList/models"

type StorageI interface{
	ToDo() ToDoRepoI
}

type ToDoRepoI interface {
	Create(entity models.Create) (err error)
	// GetList(query models.Query) (resp []models.ToDo, err error)
	// GetByID(ID int) (resp models.ToDo, err error)
	// Update(entity models.Update) (effectedRowsNum int, err error)
	// Delete(ID int) (effectedRowsNum int, err error)
}