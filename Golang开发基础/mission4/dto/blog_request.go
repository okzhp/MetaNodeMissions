package dto

type AddPostRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type EditPostRequest struct {
	ID      uint   `json:"id" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type AddCommentRequest struct {
	Content string `json:"content" binding:"required"`
	PostID  uint   `json:"postId" binding:"required"`
}
