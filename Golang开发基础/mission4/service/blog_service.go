package service

import (
	"errors"
	"myBlog/dto"
	"myBlog/model"

	"gorm.io/gorm"
)

type BlogService struct {
	db *gorm.DB
}

func NewBlogService(db *gorm.DB) *BlogService {
	return &BlogService{db}
}

// 添加文章
func (s *BlogService) AddPost(req dto.AddPostRequest, currentLoginUserId uint) error {
	post := model.Post{
		Title:   req.Title,
		Content: req.Content,
		UserID:  currentLoginUserId,
	}

	tx := s.db.Create(&post)

	if tx.RowsAffected == 0 {
		return errors.New("post数据库新增失败")
	}
	return nil
}

// 编辑文章
func (s *BlogService) EditPost(req dto.EditPostRequest, currentLoginUserId uint) error {
	var oldPost model.Post
	s.db.First(&oldPost, req.ID)

	if currentLoginUserId != oldPost.UserID {
		return errors.New("仅作者可编辑")
	}

	post := model.Post{
		Model:   gorm.Model{ID: req.ID},
		Title:   req.Title,
		Content: req.Content,
	}
	tx := s.db.Model(&post).Updates(post)

	if tx.RowsAffected == 0 {
		return errors.New("post数据库更新失败")
	}
	return nil
}

// 删除文章
func (s *BlogService) DeletePost(id uint, currentLoginUserId uint) error {
	var post model.Post
	tx := s.db.First(&post, id)

	if tx.RowsAffected == 0 {
		return errors.New("不存在的文章")
	}

	if currentLoginUserId != post.UserID {
		return errors.New("仅作者可删除")
	}

	tx = s.db.Delete(&post)

	if tx.RowsAffected == 0 {
		return errors.New("post数据库删除失败")
	}
	return nil
}

// 根据文章id查询文章及其所有评论
func (s *BlogService) QueryPostWithCommentsByPostID(id uint) (any, error) {

	var posts model.Post
	tx := s.db.Debug().Preload("Comments").
		Where("id = ?", id).
		First(&posts)

	if tx.RowsAffected == 0 {
		return nil, nil
	}

	return posts, nil
}

// 根据文章id查询所有评论
func (s *BlogService) QueryCommentsByPostID(id uint) ([]model.Comment, error) {
	var comments []model.Comment
	s.db.Where("post_id = ?", id).Find(&comments)

	return comments, nil
}

// 添加评论
func (s *BlogService) AddComment(req dto.AddCommentRequest) error {
	var post model.Post
	tx := s.db.First(&post, req.PostID)
	if tx.RowsAffected == 0 {
		return errors.New("评论文章不存在")
	}

	comment := model.Comment{
		Content: req.Content,
		PostID:  req.PostID,
	}

	tx = s.db.Create(&comment)
	if tx.RowsAffected == 0 {
		return errors.New("comment数据库新增失败")
	}

	return nil
}

// 删除评论
func (s *BlogService) DeleteComment(id uint) error {
	var comment model.Comment
	tx := s.db.First(&comment, id)

	if tx.RowsAffected == 0 {
		return errors.New("评论不存在")
	}

	tx = s.db.Delete(&comment)
	if tx.RowsAffected == 0 {
		return errors.New("comment数据库删除失败")
	}

	return nil
}
