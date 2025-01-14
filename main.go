package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/z0ff/microblog-backend/db"
	"github.com/z0ff/microblog-backend/db/model"
	"github.com/z0ff/microblog-backend/handler"
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
		tx := db_conn.Preload("User").Begin()
		tx.Order("created_at desc").Find(&posts)
		//db_conn.Find(&posts)

		c.JSON(http.StatusOK, posts)
		//c.JSON(http.StatusOK, users)
	})
	engine.POST("/post", handler.PostCreate)

	engine.Run(":3000")
}
