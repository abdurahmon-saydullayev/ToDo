package api

import (

	"GoProjects/ToDoList/api/handlers"
	"GoProjects/ToDoList/config"

	"github.com/gin-gonic/gin"
	
)

func SetUpAPI(r *gin.Engine, h handlers.Handler, cfg config.Config) {
	r.POST("/todo", h.CreateToDO)
	r.GET("/todo/:id", h.GetByID)
}
