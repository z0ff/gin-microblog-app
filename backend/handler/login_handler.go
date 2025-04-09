package handler

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/z0ff/microblog-backend/db"
	"github.com/z0ff/microblog-backend/db/model"
	"github.com/z0ff/microblog-backend/utils/crypto"
	session_mgr "github.com/z0ff/microblog-backend/utils/session"
	"net/http"
)

// type LoginHandler struct{}
type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupUser struct {
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

func Login(c *gin.Context) {
	var loginUser LoginUser
	c.BindJSON(&loginUser)

	// セッションの取得
	session := sessions.Default(c)

	db_conn := db.GetConnection()
	var user model.User
	db_conn.Where("email = ?", loginUser.Email).First(&user)

	// ユーザーが存在しない場合エラーを返す
	if user.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	// ユーザーが休止中の場合エラーを返す
	if user.SuspendedAt != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User is suspended",
		})
		return
	}

	var userAuth model.UserAuth
	db_conn.Where("user_id = ?", user.ID).First(&userAuth)

	if !crypto.ComparePassword(userAuth.Password, loginUser.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email or password",
		})
		return
	} else {
		// セッションにユーザーIDを保存
		session.Set("user_id", user.ID)
		session.Save()
		c.JSON(http.StatusOK, gin.H{
			"message": "Login success",
			"user_id": user.ID,
		})
		return
	}
}

func Logout(c *gin.Context) {
	// 未ログインの場合、認証エラーを返す
	userID := session_mgr.GetUserID(c)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	// セッションの取得
	session := sessions.Default(c)
	// セッションからユーザーIDを削除
	session.Delete("user_id")
	// セッションの保存
	session.Save()

	c.JSON(http.StatusOK, gin.H{
		"message": "Logout success",
	})
}

func Signup(c *gin.Context) {
	var signupUser SignupUser
	c.BindJSON(&signupUser)

	db_conn := db.GetConnection()
	pw_hash, _ := crypto.HashPassword(signupUser.Password)

	user := model.User{
		Name:        signupUser.Name,
		DisplayName: signupUser.DisplayName,
		Email:       signupUser.Email,
	}
	db_conn.Create(&user)
	db_conn.Create(&model.UserAuth{
		UserID:   user.ID,
		Password: pw_hash,
	})

	c.JSON(http.StatusOK, gin.H{
		"message": "Signup success",
	})
}

func GetIsLoggedIn(c *gin.Context) {
	userID := session_mgr.GetUserID(c)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Authorized",
		"user_id": userID,
	})
}
