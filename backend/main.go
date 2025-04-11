package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"os"

	"github.com/z0ff/microblog-backend/db"
	"github.com/z0ff/microblog-backend/handler"
	"github.com/z0ff/microblog-backend/handler/admin_handler"
)

type Post struct {
	UserID  uint   `json:"user_id"`
	Content string `json:"content"`
}

func main() {
	// DB接続
	_, err := db.NewDatabase()
	if err != nil {
		return
	}

	engine := gin.Default()

	// セッションの設定
	store, err := redis.NewStore(
		10,
		"tcp",
		os.Getenv("REDIS_HOST")+":"+os.Getenv("REDIS_PORT"),
		os.Getenv("REDIS_PASSWORD"),
		[]byte(os.Getenv("REDIS_SECRET")))
	if err != nil {
		panic("failed to connect redis: " + err.Error())
	}
	engine.Use(sessions.Sessions("session", store))

	engine.Use(cors.New(cors.Config{
		// アクセスを許可するアクセス元
		AllowOrigins: []string{os.Getenv("APP_ORIGIN"), os.Getenv("ADMIN_ORIGIN")},
		// アクセスを許可するHTTPメソッド
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		// 許可するHTTPリクエストヘッダ
		AllowHeaders: []string{"Origin", "Content-Length", "Content-Type"},
		// cookieなどの情報を必要とするかどうか
		AllowCredentials: true,
		// preflightリクエストの結果をキャッシュする秒数
		MaxAge: 24 * 3600,
	}))

	engine.GET("/", handler.GetTimeline)
	engine.POST("/post", handler.CreatePost)
	engine.GET("/search", handler.SearchPost)
	engine.POST("/login", handler.Login)
	engine.POST("/signup", handler.Signup)
	engine.POST("/logout", handler.Logout)
	engine.GET("/is_logged_in", handler.GetIsLoggedIn)
	engine.GET("/me", handler.GetMe)
	engine.GET("/user/:username", handler.GetUserInfo)
	engine.GET("/searchuser", handler.SearchUser)
	engine.GET("/post/:username", handler.GetPostsOfUser)
	engine.POST("/follow/:username", handler.FollowUser)
	engine.POST("/unfollow/:username", handler.UnFollowUser)
	engine.GET("/followers/:username", handler.GetFollowers)
	engine.GET("/followings/:username", handler.GetFollowings)
	engine.GET("/is_following/:username", handler.GetIsFollowing)
	engine.POST("/like/:post_id", handler.LikePost)
	engine.POST("/unlike/:post_id", handler.UnLikePost)

	admin := engine.Group("/admin")
	{
		admin.GET("/users", admin_handler.GetUsers)
		admin.GET("/user/:user_id", admin_handler.GetUser)
		admin.DELETE("/user/:user_id", admin_handler.DeleteUser)
		admin.POST("/suspend/:user_id", admin_handler.SuspendUser)
		admin.POST("/resume/:user_id", admin_handler.ResumeUser)
		admin.GET("/posts", admin_handler.GetPosts)
		admin.DELETE("/post/:post_id", admin_handler.DeletePost)

		auth := admin.Group("/auth")
		{
			auth.POST("/login", admin_handler.Login)
			auth.POST("/logout", admin_handler.Logout)
			auth.GET("/is_logged_in", admin_handler.GetIsLoggedIn)
			auth.POST("/signup", admin_handler.Signup)
			auth.GET("/me", admin_handler.GetMe)
		}
	}

	engine.Run(":3000")
}
