package service

import (
	"errors"
	"myBlog/dto"
	"myBlog/model"
	"myBlog/util"

	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db}
}

func (s *UserService) Register(req dto.RegisterRequest) error {

	var dbUser model.User
	result := s.db.Where("username = ?", req.Username).Find(&dbUser)
	if result.RowsAffected > 0 {
		return errors.New("用户名重复")
	}

	user := model.User{
		Username: req.Username,
		Password: util.CryptoStrWithDefaulSalt(req.Password),
		Email:    req.Email,
	}

	result = s.db.Create(&user)

	if result.RowsAffected == 0 {
		return errors.New("user数据库新增错误")
	}
	return nil
}

func (s *UserService) Login(req dto.LoginRequest) (string, error) {
	var dbUser model.User
	tx := s.db.Where("username = ?", req.Username).First(&dbUser)
	if tx.RowsAffected == 0 || dbUser.Password != util.CryptoStrWithDefaulSalt(req.Password) {
		return "", errors.New("用户名或密码错误")
	}

	token, err := util.GenJwtToken(&dbUser)
	if err != nil {
		return "", errors.New(err.Error())
	}
	return token, nil
}
