package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/z0ff/microblog-backend/db"
	"github.com/z0ff/microblog-backend/db/model"
)

type PostController struct{}

func (pc PostController) Create(c *gin.Context) {
	user_id, _ := strconv.ParseUint(c.PostForm("user_id"), 10, 64)
	content := c.PostForm("content")

	post := model.Post{UserID: uint(user_id), Content: content}

	db := db.GetConnection()
	db.Create(&post)

	c.JSON(http.StatusOK, gin.H{
		"user_id": uint(user_id),
		"content": content,
	})
}
