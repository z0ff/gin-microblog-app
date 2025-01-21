package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/z0ff/microblog-backend/db"
	"github.com/z0ff/microblog-backend/db/model"
	"github.com/z0ff/microblog-backend/handler"
	//"github.com/z0ff/microblog-backend/utils/session"
)

type Post struct {
	UserID  uint   `json:"user_id"`
	Content string `json:"content"`
}

func main() {
	// DB接続
	database, _ := db.NewDatabase()
	db_conn := database.Connection

	engine := gin.Default()

	// セッションの設定
	store, err := redis.NewStore(10, "tcp", "redis:6379", "", []byte("secret"))
	if err != nil {
		panic("failed to connect redis: " + err.Error())
	}
	engine.Use(sessions.Sessions("session", store))

	engine.Use(cors.New(cors.Config{
		// アクセスを許可するアクセス元
		AllowOrigins: []string{"http://localhost:5173"},
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
		var posts []model.Post

		// ユーザーIDをセッションから取得
		//userID := session.GetUserID(c)
		session := sessions.Default(c)
		userID := session.Get("user_id")
		// ユーザーIDが取得できない場合、認証エラーを返す
		if userID == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			return
		}

		tx := db_conn.Preload("User").Begin()
		tx.Order("created_at desc").Find(&posts)
		//db_conn.Find(&posts)

		c.JSON(http.StatusOK, posts)
		//c.JSON(http.StatusOK, users)
	})
	engine.POST("/post", handler.PostCreate)
	engine.POST("/login", handler.Login)
	engine.POST("/signup", handler.Signup)

	engine.Run(":3000")
}
