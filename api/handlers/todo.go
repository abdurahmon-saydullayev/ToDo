package handlers

import (
	"GoProjects/ToDoList/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

// CreateArticle godoc
// @tags article
// @ID create-article-handler
// @Summary Create Article
// @Description Create Article By Given Info and Author ID
// @Param data body models.ArticleCreateModel true "Article Body"
// @Accept  json
// @Produce  json
// @Success 200 {object} models.SuccessResponse{data=string}
// @Failure default {object} models.DefaultError
// @Router /articles [POST]
func (h *Handler) CreateToDO(c *gin.Context) {
	var entity models.Create
	err := c.BindJSON(&entity)
	if err != nil {
		c.JSON(400, models.DefaultError{
			Message: err.Error(),
		})
		return
	}

	fmt.Println(entity)

	err = h.strg.ToDo().Create(entity)

	if err != nil {
		c.JSON(400, models.DefaultError{
			Message: err.Error(),
		})
	}

	c.JSON(200, models.SuccessResponse{
		Message: "article has been created",
		Data:    "ok",
	})
}

