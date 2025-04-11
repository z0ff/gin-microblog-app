package session

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// ユーザーIDをセッションから取得
func GetUserID(c *gin.Context) uint {
	session := sessions.Default(c)
	if session == nil {
		panic("session is nil")
	}
	userID := session.Get("user_id")
	if userID == nil {
		return 0
	}
	return userID.(uint)
}

// ユーザーIDをセッションから削除
func DeleteUserID(c *gin.Context) {
	session := sessions.Default(c)
	if session == nil {
		panic("session is nil")
	}
	// セッションからユーザーIDを削除
	session.Delete("user_id")
	// セッションの保存
	session.Save()
}

func GetAdminID(c *gin.Context) uint {
	session := sessions.Default(c)
	if session == nil {
		panic("session is nil")
	}
	adminID := session.Get("admin_id")
	if adminID == nil {
		return 0
	}
	return adminID.(uint)
}
