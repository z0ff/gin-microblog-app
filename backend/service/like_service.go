package service

import (
	"github.com/z0ff/microblog-backend/db"
	"github.com/z0ff/microblog-backend/db/model"
)

func isLiked(userID uint, postID uint) bool {
	db_conn := db.GetConnection()

	var like model.Like

	result := db_conn.Where("user_id = ? AND post_id = ?", userID, postID).First(&like)

	if result.Error != nil {
		return false
	}

	return true
}
