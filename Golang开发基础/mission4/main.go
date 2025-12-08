package main

import (
	"encoding/json"
	"fmt"
	"log"
	"myBlog/controller"
	"myBlog/dto"
	"myBlog/model"
	"myBlog/util"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db := initMysqlDataBase()
	r := gin.Default()

	userController := controller.NewUserController(db)
	blogController := controller.NewBlogController(db)

	//用户注册
	r.POST("/user/register", userController.Register)

	//用户登录
	r.POST("/user/login", userController.Login)

	//添加文章
	r.POST("/addPost", checkAuthentication(), blogController.AddPost)

	//编辑文章
	r.POST("/editPost", checkAuthentication(), blogController.EditPost)

	//删除文章
	r.POST("/deletePost/:id", checkAuthentication(), blogController.DeletePost)

	//查询文章及评论
	r.GET("/queryPost/:id", blogController.QueryPostWithCommentsByPostID)

	//查询文章的评论
	r.GET("/queryPostComment/:id", blogController.QueryCommentsByPostID)

	//添加评论
	r.POST("/addComment", checkAuthentication(), blogController.AddComment)

	//删除评论
	r.POST("/deleteComment/:id", checkAuthentication(), blogController.DeleteComment)

	r.Run(":8080")
}

type DBConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	DBName   string `json:"dbName"`
	IP       string `json:"ip"`
}

func initMysqlDataBase() *gorm.DB {
	dsn := readDBConfig()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("连接数据库失败")
	}

	db.AutoMigrate(&model.Comment{})
	db.AutoMigrate(&model.Post{})
	db.AutoMigrate(&model.User{})

	fmt.Println("✌️✌️✌️数据库初始化成功")
	return db
}

// 读取数据库配置
func readDBConfig() string {
	data, err := os.ReadFile("db.json")
	if err != nil {
		panic(err)
	}

	var cfg DBConfig
	if err := json.Unmarshal(data, &cfg); err != nil {
		panic(err)
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Username,
		cfg.Password,
		cfg.IP,
		cfg.DBName,
	)

	return dsn
}

// jwt鉴权
func checkAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusOK, dto.ErrMsg("token不能为空"))
			c.Abort()
			return
		}
		//解析token
		userClaims, err := util.ParseJwtToken(token)
		if err != nil {
			c.JSON(http.StatusOK, dto.ErrMsg(err.Error()))
			c.Abort()
			return
		}
		//将当前登录用户保存上下文
		c.Set(controller.CurrentLoginUserId, userClaims.ID)
		c.Next()
	}
}
