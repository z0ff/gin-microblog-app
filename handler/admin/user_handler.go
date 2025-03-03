package admin_handler

import (
	"github.com/gin-gonic/gin"
	"github.com/z0ff/microblog-backend/db"
	"github.com/z0ff/microblog-backend/db/model"
	"net/http"
)

func GetUsers(c *gin.Context) {
	db_conn := db.GetConnection()
	var users []model.User
	db_conn.Find(&users)

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func DeleteUser(c *gin.Context) {
	userID := c.Param("user_id")

	db_conn := db.GetConnection()
	var user model.User
	result := db_conn.Where("id = ?", userID).First(&user)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	db_conn.Delete(&user)

	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted",
	})
}
