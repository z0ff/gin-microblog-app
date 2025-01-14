package handler

// package post_handler

import (
	"github.com/z0ff/microblog-backend/db"
	"github.com/z0ff/microblog-backend/db/model"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

// type PostHandler struct{}

type Post struct {
	UserID  uint   `json:"user_id"`
	Content string `json:"content"`
}

// func (pc PostHandler) Create(c *gin.Context) {
func PostCreate(c *gin.Context) {
	var post Post
	c.BindJSON(&post)

	slog.Debug("debug", "post", post)

	db_conn := db.GetConnection()
	//db_conn.Create(&post)
	db_conn.Create(&model.Post{
		UserID:  post.UserID,
		Content: post.Content,
	})

	c.JSON(http.StatusOK, gin.H{
		"user_id": post.UserID,
		"content": post.Content,
	})
}
