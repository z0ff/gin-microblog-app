package service

import (
	//"gorm.io/gorm"

	"github.com/z0ff/microblog-backend/db"
	"github.com/z0ff/microblog-backend/db/model"
)

type User struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
}

type Post struct {
	ID        uint   `json:"id"`
	UserID    uint   `json:"user_id"`
	User      User   `json:"user"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdAt"`
}

type PostWithIsLiked struct {
	Post
	IsLiked bool `json:"is_liked"`
}

func GetPosts(currentUserID uint, targetUserID uint) ([]PostWithIsLiked, error) {
	var posts []Post

	db_conn := db.GetConnection()
	tx := db_conn.Preload("User").Begin()
	tx.Model(&model.Post{}).Order("created_at desc").Where("user_id = ?", targetUserID).Find(&posts)

	postsWithIsLiked := []PostWithIsLiked{}
	for _, post := range posts {
		postsWithIsLiked = append(postsWithIsLiked, PostWithIsLiked{
			Post:    post,
			IsLiked: isLiked(currentUserID, post.ID),
		})
	}

	return postsWithIsLiked, nil
}

func GetPostsByUserIDs(currentUserID uint, targetUserIDs []uint) ([]PostWithIsLiked, error) {
	var posts []Post

	db_conn := db.GetConnection()
	tx := db_conn.Preload("User").Begin()
	tx.Model(&model.Post{}).Order("created_at desc").Where("user_id in ?", targetUserIDs).Find(&posts)

	postsWithIsLiked := []PostWithIsLiked{}
	for _, post := range posts {
		postsWithIsLiked = append(postsWithIsLiked, PostWithIsLiked{
			Post:    post,
			IsLiked: isLiked(currentUserID, post.ID),
		})
	}

	return postsWithIsLiked, nil
}

func SearchPostsByString(currentUserID uint, searchStr string) ([]PostWithIsLiked, error) {
	var posts []Post

	db_conn := db.GetConnection()
	tx := db_conn.Preload("User").Begin()
	tx.Model(&model.Post{}).Order("created_at desc").Where("content like ?", "%"+searchStr+"%").Find(&posts)

	postsWithIsLiked := []PostWithIsLiked{}
	for _, post := range posts {
		postsWithIsLiked = append(postsWithIsLiked, PostWithIsLiked{
			Post:    post,
			IsLiked: isLiked(currentUserID, post.ID),
		})
	}

	return postsWithIsLiked, nil
}
