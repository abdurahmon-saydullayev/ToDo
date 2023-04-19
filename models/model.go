package models

import "time"

type ToDo struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Created     *time.Time `json:"created_at"`
	Updated     *time.Time `json:"updated_at"`
	Completed    bool      `json:"completed"`
}

type Create struct{
	Title string `json:"title"`
	Description string `json:"description"`
}

type Update struct{
	Title string `json:"title"`
	Description string `json:"description"`
	Completed bool `json:"completed"`
}