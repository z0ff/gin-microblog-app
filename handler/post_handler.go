package handler

// package post_handler

import (
	"errors"
	"github.com/z0ff/microblog-backend/db"
	"github.com/z0ff/microblog-backend/db/model"
	"github.com/z0ff/microblog-backend/service"
	"github.com/z0ff/microblog-backend/utils/session"
	"gorm.io/gorm"
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

	//var followings []uint

	userId := session.GetUserID(c)

	if userId == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}
	//followings = append(followings, userId)
	//
	//db_conn := db.GetConnection()
	//tx := db_conn.Preload("User").Begin()
	//tx.Order("created_at desc").Where("content LIKE ?", "%"+query+"%").Find(&posts)

	posts, _ := service.SearchPostsByString(userId, query)

	c.JSON(http.StatusOK, posts)
}

func GetTimeline(c *gin.Context) {
	var posts []service.PostWithIsLiked
	var followings []uint

	// ユーザーIDをセッションから取得
	userID := session.GetUserID(c)
	// ユーザーIDが取得できない場合、認証エラーを返す
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	db_conn := db.GetConnection()
	// フォローしているユーザーのIDを取得
	db_conn.Table("follows").Select("following_id").Where("user_id = ? AND deleted_at is null", userID).Find(&followings)
	followings = append(followings, userID)

	//tx := db_conn.Preload("User").Begin()
	//tx.Order("created_at desc").Where("user_id in ?", followings).Find(&posts)

	posts, _ = service.GetPostsByUserIDs(userID, followings)

	c.JSON(http.StatusOK, posts)
}

func GetPostsOfUser(c *gin.Context) {
	var userId uint
	userId = session.GetUserID(c)

	if userId == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	username := c.Param("username")

	// ユーザーの存在確認
	var user model.User
	db_conn := db.GetConnection()
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

	posts, _ := service.GetPosts(userId, user.ID)

	c.JSON(http.StatusOK, posts)
}

func LikePost(c *gin.Context) {
	var userId uint
	userId = session.GetUserID(c)

	if userId == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	postId := c.Param("post_id")

	var post model.Post
	db_conn := db.GetConnection()
	result := db_conn.Where("id = ?", postId).First(&post)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Post not found",
		})
		return
	} else if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return
	}

	db_conn.Create(&model.Like{
		UserID: userId,
		PostID: post.ID,
	})

	c.JSON(http.StatusOK, gin.H{
		"message": "Like post success",
	})
}

func UnLikePost(c *gin.Context) {
	var userId uint
	userId = session.GetUserID(c)

	if userId == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	postId := c.Param("post_id")

	var post model.Post
	db_conn := db.GetConnection()
	result := db_conn.Where("id = ?", postId).First(&post)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Post not found",
		})
		return
	} else if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return
	}

	db_conn.Where("user_id = ? AND post_id = ?", userId, post.ID).Delete(&model.Like{})

	c.JSON(http.StatusOK, gin.H{
		"message": "Unlike post success",
	})
}
