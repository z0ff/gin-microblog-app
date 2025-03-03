package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"

	"github.com/z0ff/microblog-backend/db"
	"github.com/z0ff/microblog-backend/handler"
	"github.com/z0ff/microblog-backend/handler/admin_handler"
	"github.com/z0ff/microblog-backend/service"
	"github.com/z0ff/microblog-backend/utils/session"
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

		// フォローしているユーザーのIDを取得
		db_conn.Table("follows").Select("following_id").Where("user_id = ? AND deleted_at is null", userID).Find(&followings)
		followings = append(followings, userID)

		//tx := db_conn.Preload("User").Begin()
		//tx.Order("created_at desc").Where("user_id in ?", followings).Find(&posts)

		posts, err = service.GetPostsByUserIDs(userID, followings)

		c.JSON(http.StatusOK, posts)
	})
	engine.POST("/post", handler.CreatePost)
	engine.GET("/search", handler.SearchPost)
	engine.POST("/login", handler.Login)
	engine.POST("/signup", handler.Signup)
	engine.POST("/logout", handler.Logout)
	engine.GET("/is_logged_in", handler.GetIsLoggedIn)
	engine.GET("/me", handler.GetMe)
	engine.GET("/user/:username", handler.GetUserInfo)
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
		admin.DELETE("/user/:user_id", admin_handler.DeleteUser)
		admin.GET("/posts", admin_handler.GetPosts)
		admin.DELETE("/post/:post_id", admin_handler.DeletePost)
	}

	engine.Run(":3000")
}
