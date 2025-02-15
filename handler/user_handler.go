package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/z0ff/microblog-backend/db"
	"github.com/z0ff/microblog-backend/db/model"
	session_mgr "github.com/z0ff/microblog-backend/utils/session"
	"gorm.io/gorm"
	"net/http"
)

func GetMe(c *gin.Context) {
	userID := session_mgr.GetUserID(c)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	db_conn := db.GetConnection()
	var user model.User
	db_conn.Where("id = ?", userID).First(&user)

	c.JSON(http.StatusOK, gin.H{
		"user_id":      user.ID,
		"name":         user.Name,
		"display_name": user.DisplayName,
	})
}

func GetUserInfo(c *gin.Context) {
	username := c.Param("username")

	db_conn := db.GetConnection()
	var user model.User
	result := db_conn.Where("name = ?", username).First(&user)

	// ユーザーが存在しない場合、404エラーを返す
	// その他のエラーの場合、500エラーを返す
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	} else if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id":      user.ID,
		"name":         user.Name,
		"display_name": user.DisplayName,
	})
}

func FollowUser(c *gin.Context) {
	userID := session_mgr.GetUserID(c)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	username := c.Param("username")

	db_conn := db.GetConnection()
	var user model.User
	result := db_conn.Where("name = ?", username).First(&user)

	// ユーザーが存在しない場合、404エラーを返す
	// その他のエラーの場合、500エラーを返す
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	} else if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return
	}

	var follow model.Follow
	db_conn.Where("user_id = ? AND following_id = ?", userID, user.ID).First(&follow)
	if follow.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Already followed",
		})
		return
	}

	follow = model.Follow{
		UserID:      userID,
		FollowingID: user.ID,
	}
	db_conn.Create(&follow)

	c.JSON(http.StatusOK, gin.H{
		"message": "Follow success",
	})
}
