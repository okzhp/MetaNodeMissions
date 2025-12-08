package controller

import (
	"myBlog/dto"
	"myBlog/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{service.NewUserService(db)}
}

func (ctl *UserController) Register(c *gin.Context) {
	var req dto.RegisterRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, dto.ErrMsg(err.Error()))
		return
	}

	err = ctl.userService.Register(req)
	if err != nil {
		c.JSON(http.StatusOK, dto.ErrMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, dto.SuccessMsg())
}

func (ctl *UserController) Login(c *gin.Context) {
	var req dto.LoginRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, dto.ErrMsg(err.Error()))
		return
	}

	token, err := ctl.userService.Login(req)
	if err != nil {
		c.JSON(http.StatusOK, dto.ErrMsg(err.Error()))
		return
	}

	c.JSON(http.StatusOK, dto.SuccessData(token))
}
