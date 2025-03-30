package admin_handler

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/z0ff/microblog-backend/db"
	"github.com/z0ff/microblog-backend/db/model"
	"github.com/z0ff/microblog-backend/utils/crypto"
	session_mgr "github.com/z0ff/microblog-backend/utils/session"
)

type LoginAdmin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupAdmin struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var loginAdmin LoginAdmin
	c.BindJSON(&loginAdmin)

	// セッションの取得
	session := sessions.Default(c)

	// DB接続
	db_conn := db.GetConnection()
	var admin model.Admin
	db_conn.Where("email = ?", loginAdmin.Email).First(&admin)
	var adminAuth model.AdminAuth
	db_conn.Where("admin_id = ?", admin.ID).First(&adminAuth)

	// パスワードの比較
	if !crypto.ComparePassword(adminAuth.Password, loginAdmin.Password) {
		c.JSON(401, gin.H{
			"error": "Invalid email or password",
		})
		return
	} else {
		// セッションに管理者IDを保存
		session.Set("admin_id", admin.ID)
		session.Save()
		c.JSON(200, gin.H{
			"message":  "Login success",
			"admin_id": admin.ID,
		})
		return
	}
}

func Logout(c *gin.Context) {
	// 未ログインの場合、認証エラーを返す
	adminID := session_mgr.GetAdminID(c)
	if adminID == 0 {
		c.JSON(401, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	// セッションから管理者IDを削除
	session := sessions.Default(c)
	session.Delete("admin_id")
	session.Save()

	c.JSON(200, gin.H{
		"message": "Logout success",
	})
}

func Signup(c *gin.Context) {
	var signupAdmin SignupAdmin
	c.BindJSON(&signupAdmin)

	// DB接続
	db_conn := db.GetConnection()

	// 管理者の登録
	admin := model.Admin{
		Name:  signupAdmin.Name,
		Email: signupAdmin.Email,
	}
	db_conn.Create(&admin)

	// パスワードのハッシュ化
	pw_hash, _ := crypto.HashPassword(signupAdmin.Password)

	// 管理者l認証情報の登録
	adminAuth := model.AdminAuth{
		AdminID:  admin.ID,
		Password: pw_hash,
	}
	db_conn.Create(&adminAuth)

	c.JSON(200, gin.H{
		"message":  "Signup success",
		"admin_id": admin.ID,
	})
}

func GetIsLoggedIn(c *gin.Context) {
	adminID := session_mgr.GetAdminID(c)
	if adminID == 0 {
		c.JSON(401, gin.H{
			"error": "Unauthorized",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "Logged in",
		})
	}
}

func GetMe(c *gin.Context) {
	adminID := session_mgr.GetAdminID(c)
	if adminID == 0 {
		c.JSON(401, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	db_conn := db.GetConnection()
	var admin model.Admin
	db_conn.Where("id = ?", adminID).First(&admin)

	c.JSON(200, gin.H{
		"admin_id":   admin.ID,
		"name":       admin.Name,
		"email":      admin.Email,
		"created_at": admin.CreatedAt,
		"updated_at": admin.UpdatedAt,
	})
}
