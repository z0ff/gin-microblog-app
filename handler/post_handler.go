package handler

// package post_handler

import (
	"github.com/z0ff/microblog-backend/db"
	"github.com/z0ff/microblog-backend/db/model"
	"github.com/z0ff/microblog-backend/utils/session"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

// type PostHandler struct{}

type Post struct {
	//UserID  uint   `json:"user_id"`
	Content string `json:"content"`
}

// func (pc PostHandler) Create(c *gin.Context) {
func CreatePost(c *gin.Context) {
	var userId uint
	userId = session.GetUserID(c)

	if userId == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	var post Post
	c.BindJSON(&post)

	slog.Debug("debug", "post", post)

	db_conn := db.GetConnection()
	//db_conn.Create(&post)
	db_conn.Create(&model.Post{
		UserID:  userId,
		Content: post.Content,
	})

	c.JSON(http.StatusOK, gin.H{
		"user_id": userId,
		"content": post.Content,
	})
}

func SearchPost(c *gin.Context) {
	// 検索クエリを取得
	query := c.Query("query")

	var posts []model.Post
	//var followings []uint

	userId := session.GetUserID(c)

	if userId == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}
	//followings = append(followings, userId)

	db_conn := db.GetConnection()
	tx := db_conn.Preload("User").Begin()
	tx.Order("created_at desc").Where("content LIKE ?", "%"+query+"%").Find(&posts)

	c.JSON(http.StatusOK, posts)
}
