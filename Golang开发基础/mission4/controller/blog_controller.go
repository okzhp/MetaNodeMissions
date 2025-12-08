package controller

import (
	"myBlog/dto"
	"myBlog/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const CurrentLoginUserId = "currentLoginUserId"

type BlogController struct {
	blogService *service.BlogService
}

func NewBlogController(db *gorm.DB) *BlogController {
	return &BlogController{service.NewBlogService(db)}
}

func (ctl *BlogController) AddPost(c *gin.Context) {
	var req dto.AddPostRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, dto.ErrMsg(err.Error()))
		return
	}

	err = ctl.blogService.AddPost(req, c.GetUint(CurrentLoginUserId))
	if err != nil {
		c.JSON(http.StatusOK, dto.ErrMsg(err.Error()))
		return
	}

	c.JSON(http.StatusOK, dto.SuccessMsg())
}

func (ctl *BlogController) EditPost(c *gin.Context) {
	var req dto.EditPostRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, dto.ErrMsg(err.Error()))
		return
	}
	err = ctl.blogService.EditPost(req, c.GetUint(CurrentLoginUserId))
	if err != nil {
		c.JSON(http.StatusOK, dto.ErrMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, dto.SuccessMsg())
}

func (ctl *BlogController) DeletePost(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	err := ctl.blogService.DeletePost(uint(id), c.GetUint(CurrentLoginUserId))
	if err != nil {
		c.JSON(http.StatusOK, dto.ErrMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, dto.SuccessMsg())
}

func (ctl *BlogController) QueryPostWithCommentsByPostID(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	data, err := ctl.blogService.QueryPostWithCommentsByPostID(uint(id))
	if err != nil {
		c.JSON(http.StatusOK, dto.ErrMsg(err.Error()))
		return
	}

	c.JSON(http.StatusOK, dto.SuccessData(data))
}

func (ctl *BlogController) QueryCommentsByPostID(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	data, err := ctl.blogService.QueryCommentsByPostID(uint(id))
	if err != nil {
		c.JSON(http.StatusOK, dto.ErrMsg(err.Error()))
		return
	}

	c.JSON(http.StatusOK, dto.SuccessData(data))
}

func (ctl *BlogController) AddComment(c *gin.Context) {
	var req dto.AddCommentRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, dto.ErrMsg(err.Error()))
		return
	}

	err = ctl.blogService.AddComment(req)
	if err != nil {
		c.JSON(http.StatusOK, dto.ErrMsg(err.Error()))
		return
	}

	c.JSON(http.StatusOK, dto.SuccessMsg())

}

func (ctl *BlogController) DeleteComment(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	err := ctl.blogService.DeleteComment(uint(id))
	if err != nil {
		c.JSON(http.StatusOK, dto.ErrMsg(err.Error()))
		return
	}

	c.JSON(http.StatusOK, dto.SuccessMsg())
}
