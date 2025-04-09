package admin_handler

import (
	"github.com/gin-gonic/gin"
	"github.com/z0ff/microblog-backend/db"
	"github.com/z0ff/microblog-backend/db/model"
	"net/http"
	"time"
)

type User struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
	DeletedAt   string `json:"deletedAt"`
	SuspendedAt string `json:"suspendedAt"`
}

func GetUsers(c *gin.Context) {
	db_conn := db.GetConnection()
	var users []User
	db_conn.Find(&users)

	c.JSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {
	userID := c.Param("user_id")
	db_conn := db.GetConnection()
	var users []User
	db_conn.Find(&users).Where("id = ?", userID)

	c.JSON(http.StatusOK, users)
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

func SuspendUser(c *gin.Context) {
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

	if user.SuspendedAt != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User is already suspended",
		})
		return
	}

	now := time.Now()
	user.SuspendedAt = &now
	result = db_conn.Save(&user)

	c.JSON(http.StatusOK, gin.H{
		"message": "User suspended",
	})
}
func ResumeUser(c *gin.Context) {
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

	if user.SuspendedAt == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User is not suspended",
		})
		return
	}

	user.SuspendedAt = nil
	result = db_conn.Save(&user)

	c.JSON(http.StatusOK, gin.H{
		"message": "User resumed",
	})
}
