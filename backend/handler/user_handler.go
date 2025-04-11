package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/z0ff/microblog-backend/db"
	"github.com/z0ff/microblog-backend/db/model"
	"github.com/z0ff/microblog-backend/service"
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
	db_conn.Debug().Where("id = ?", userID).First(&user)

	// ユーザーが存在しない場合、ログアウトし404エラーを返す
	if user.ID == 0 {
		session_mgr.DeleteUserID(c)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	// ユーザーが停止中の場合、ログアウトし403エラーを返す
	if user.SuspendedAt != nil {
		session_mgr.DeleteUserID(c)
		c.JSON(http.StatusForbidden, gin.H{
			"error": "User suspended",
		})
		return
	}

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

func SearchUser(c *gin.Context) {
	// 検索クエリを取得
	query := c.Query("query")

	userId := session_mgr.GetUserID(c)

	if userId == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	users, _ := service.SearchUsersByDisplayName(userId, query)

	c.JSON(http.StatusOK, users)
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
	db_conn.Where("user_id = ? AND following_id = ? AND deleted_at is null", userID, user.ID).First(&follow)
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

func UnFollowUser(c *gin.Context) {
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
	if follow.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Not followed",
		})
		return
	}

	db_conn.Delete(&follow)

	c.JSON(http.StatusOK, gin.H{
		"message": "Unfollow success",
	})
}

func GetFollowings(c *gin.Context) {
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

	var followings []model.Follow
	db_conn.Where("user_id = ?", user.ID).Find(&followings)

	var followingUsers []model.User
	for _, following := range followings {
		var followingUser model.User
		db_conn.Where("id = ?", following.FollowingID).First(&followingUser)
		followingUsers = append(followingUsers, followingUser)
	}

	c.JSON(http.StatusOK, followingUsers)
}

func GetFollowers(c *gin.Context) {
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

	var followers []model.Follow
	db_conn.Where("following_id = ?", user.ID).Find(&followers)

	var followerUsers []model.User
	for _, follower := range followers {
		var followerUser model.User
		db_conn.Where("id = ?", follower.UserID).First(&followerUser)
		followerUsers = append(followerUsers, followerUser)
	}

	c.JSON(http.StatusOK, followerUsers)
}

func GetIsFollowing(c *gin.Context) {
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
	db_conn.Where("user_id = ? AND following_id = ? AND deleted_at is null", userID, user.ID).First(&follow)

	if follow.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"is_following": false,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"is_following": true,
		})
	}
}
