package admin_handler

import (
	"github.com/gin-gonic/gin"
	"github.com/z0ff/microblog-backend/db"
	"github.com/z0ff/microblog-backend/db/model"
	"net/http"
)

type Post struct {
	ID        uint   `json:"id"`
	UserID    uint   `json:"user_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	DeletedAt string `json:"deletedAt"`
}

func GetPosts(c *gin.Context) {
	db_conn := db.GetConnection()
	var posts []Post
	//tx := db_conn.Preload("User").Begin()
	db_conn.Find(&posts)

	c.JSON(http.StatusOK, posts)
}

func DeletePost(c *gin.Context) {
	postID := c.Param("post_id")

	db_conn := db.GetConnection()
	var post model.Post
	result := db_conn.Where("id = ?", postID).First(&post)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Post not found",
		})
		return
	}

	db_conn.Delete(&post)

	c.JSON(http.StatusOK, gin.H{
		"message": "Post deleted",
	})
}
