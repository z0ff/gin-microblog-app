package main

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	// "github.com/z0ff/microblog-backend/controller"
	"github.com/z0ff/microblog-backend/db"
	"github.com/z0ff/microblog-backend/db/model"
)

type Post struct {
	UserID  uint   `json:"user_id"`
	Content string `json:"content"`
}

func main() {
	database, _ := db.NewDatabase()
	db_conn := database.Connection

	var users []model.User
	db_conn.Find(&users)

	var posts []model.Post
	db_conn.Find(&posts)

	engine := gin.Default()

	engine.Use(cors.New(cors.Config{
		// アクセスを許可するアクセス元
		AllowOrigins: []string{"*"},
		// アクセスを許可するHTTPメソッド
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		// 許可するHTTPリクエストヘッダ
		AllowHeaders: []string{"Origin", "Content-Length", "Content-Type"},
		// cookieなどの情報を必要とするかどうか
		AllowCredentials: true,
		// preflightリクエストの結果をキャッシュする秒数
		MaxAge: 24 * 3600,
	}))

	engine.GET("/", func(c *gin.Context) {
		// c.JSON(http.StatusOK, gin.H{
		// 	"message": "Hello, World!",
		// })

		tx := db_conn.Preload("User").Begin()
		tx.Find(&posts)

		c.JSON(http.StatusOK, posts)
		// c.JSON(http.StatusOK, users)
	})
	// engine.POST("/post", controller.PostController.Create)

	engine.POST("/post", func(c *gin.Context) {
		// user_id, _ := strconv.ParseUint(c.PostForm("user_id"), 10, 64)
		// content := c.PostForm("content")

		// post := model.Post{UserID: uint(user_id), Content: content}

		var post Post
		c.BindJSON(&post)

		buf := make([]byte, 2048)
		n, _ := c.Request.Body.Read(buf)
		b := string(buf[0:n])
		fmt.Println(b)

		slog.Debug("debug", "post", post)

		// db := db.GetConnection()
		db_conn.Create(&post)

		c.JSON(http.StatusOK, gin.H{
			"user_id": post.UserID,
			"content": post.Content,
		})
	})

	engine.Run(":3000")
}
