package service

import (
	//"gorm.io/gorm"

	"github.com/z0ff/microblog-backend/db"
	"github.com/z0ff/microblog-backend/db/model"
)

type PostWithIsLiked struct {
	model.Post
	IsLiked bool
}

func GetPosts(currentUserID uint, targetUserID uint) ([]PostWithIsLiked, error) {
	var posts []model.Post

	db_conn := db.GetConnection()
	tx := db_conn.Preload("User").Begin()
	tx.Order("created_at desc").Where("user_id = ?", targetUserID).Find(&posts)

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
	var posts []model.Post

	db_conn := db.GetConnection()
	tx := db_conn.Preload("User").Begin()
	tx.Order("created_at desc").Where("user_id in ?", targetUserIDs).Find(&posts)

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
	var posts []model.Post

	db_conn := db.GetConnection()
	tx := db_conn.Preload("User").Begin()
	tx.Order("created_at desc").Where("content like ?", "%"+searchStr+"%").Find(&posts)

	postsWithIsLiked := []PostWithIsLiked{}
	for _, post := range posts {
		postsWithIsLiked = append(postsWithIsLiked, PostWithIsLiked{
			Post:    post,
			IsLiked: isLiked(currentUserID, post.ID),
		})
	}

	return postsWithIsLiked, nil
}
