package handler

import (
	"github.com/z0ff/microblog-backend/db"
	"github.com/z0ff/microblog-backend/db/model"
	"net/http"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/z0ff/microblog-backend/utils/crypto"
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

	if !crypto.ComparePassword(user.Password, loginUser.Password) {
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
	c.JSON(http.StatusOK, gin.H{
		"message": "Logout success",
	})
}

func Signup(c *gin.Context) {
	var signupUser SignupUser
	c.BindJSON(&signupUser)

	db_conn := db.GetConnection()
	pw_hash, _ := crypto.HashPassword(signupUser.Password)
	db_conn.Create(&model.User{
		Name:        signupUser.Name,
		DisplayName: signupUser.DisplayName,
		Email:       signupUser.Email,
		Password:    pw_hash,
	})

	c.JSON(http.StatusOK, gin.H{
		"message": "Signup success",
	})
}
