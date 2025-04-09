package service

import (
	"github.com/z0ff/microblog-backend/db"
)

type UserWithIsFollowed struct {
	User
	IsFollowed bool `json:"is_followed"`
}

func SearchUsersByDisplayName(currentUserID uint, searchStr string) ([]UserWithIsFollowed, error) {
	var users []User

	db_conn := db.GetConnection()
	db_conn.Order("created_at desc").Where("display_name like ?", "%"+searchStr+"%").Find(&users)

	usersWithIsFollowed := []UserWithIsFollowed{}
	for _, user := range users {
		usersWithIsFollowed = append(usersWithIsFollowed, UserWithIsFollowed{
			User:       user,
			IsFollowed: isFollowed(currentUserID, user.ID),
		})
	}

	return usersWithIsFollowed, nil
}

func isFollowed(currentUserID uint, userID uint) bool {
	var count int64
	db_conn := db.GetConnection()
	db_conn.Where("user_id = ? AND following_id = ? AND deleted_at is null", currentUserID, userID).Count(&count)
	return count > 0
}
