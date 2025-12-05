package main

import (
  "fmt"
  "log"

  "gorm.io/driver/mysql"
  "gorm.io/gorm"
  "gorm.io/gorm/clause"
)

//题目1：模型定义
//假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
//要求 ：
//使用Gorm定义 User 、 Post 和 Comment 模型，
//其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章），
//Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
//编写Go代码，使用Gorm创建这些模型对应的数据库表。

//题目2：关联查询
//基于上述博客系统的模型定义。
//要求 ：
//编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
//编写Go代码，使用Gorm查询评论数量最多的文章信息。

//题目3：钩子函数
//继续使用博客系统的模型。
//要求 ：
//为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
//为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，
//如果评论数量为 0，则更新文章的评论状态为 "无评论"。

type User struct {
  ID        uint
  UserName  string
  Email     string
  PostCount uint
  Posts     []Post
}

// 新建Post后更新user表 文章数量字段+1
func (p *Post) AfterCreate(db *gorm.DB) (err error) {
  db.Model(&User{}).Where("id = ?", (*p).UserID).Update("post_count", gorm.Expr("post_count + 1"))
  if db.Error != nil {
    return err
  }
  return nil
}

// 删除Post后更新user表 文章数量字段-1
func (p *Post) AfterDelete(db *gorm.DB) (err error) {
  db.Model(&User{}).Where("id = ?", (*p).UserID).Update("post_count", gorm.Expr("post_count - 1"))
  if db.Error != nil {
    return err
  }
  return nil
}

type Post struct {
  ID            uint
  Title         string
  Content       string
  CommentCount  uint
  CommentStatus string `gorm:"default:'无评论'"`
  UserID        uint
  Comments      []Comment
}

// 添加评论后文章评论数+1，并修改评论状态
func (c *Comment) AfterCreate(db *gorm.DB) (err error) {
  db.Model(&Post{}).Where("id = ?", (*c).PostID).Update("comment_count", gorm.Expr("comment_count + 1"))
  if db.Error != nil {
    return err
  }
  db.Model(&Post{}).Where("id = ?", c.PostID).Update("comment_status", "有评论")
  return nil
}

// 删除评论后文章评论数-1
func (c *Comment) AfterDelete(db *gorm.DB) (err error) {
  db.Model(&Post{}).Where("id = ?", (*c).PostID).Update("comment_count", gorm.Expr("comment_count - 1"))
  if db.Error != nil {
    return err
  }
  var post Post
  db.First(&post, c.PostID)
  if post.CommentCount == 0 {
    db.Model(&Post{}).Where("id = ?", c.PostID).Update("comment_status", "无评论")
  }
  return nil
}

type Comment struct {
  ID      uint
  Content string
  PostID  uint
}

func main() {
  dsn := "root:rootroot@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
  if err != nil {
    log.Fatalln("连接MySQL失败", err)
  }

  //初始化数据库
  initDataBase(db)

  //1、查询用户文章及评论
  user := queryPostsAndCommentsByUserId(1, db)
  fmt.Println("1、用户所有文章查询结果：", *user)
  fmt.Println()

  //2、查询最多评论数的文章
  post := queryMostCommentsPost(db)
  fmt.Println("2、评论数量最多的文章是:", *post)

  //3、给用户id=1新增一篇文章
  addedPost := addPost("新增文章title测试", "新增文章content测试", 1, db)
  fmt.Println("新增文章成功：", addedPost)

  //4、给新增的文章添加一条评论
  addedComment := addComment("太强了", addedPost.ID, db)
  fmt.Println("新增评论成功：", addedComment)

  //5、删除刚添加的评论
  deleteComment(addedComment.ID, db)
  fmt.Println("删除评论成功")
}

// 根据userID查询 该用户的所有文章及评论
func queryPostsAndCommentsByUserId(userID uint, db *gorm.DB) *User {
  var user User
  //嵌套预加载、预加载所有关联节点
  db.Preload("Posts.Comments").Preload(clause.Associations).First(&user, userID)
  return &user
}

// 查询评论数量最多的文章
func queryMostCommentsPost(db *gorm.DB) *Post {
  var postID uint
  db.Model(&Comment{}).
    Select("post_id").
    Group("post_id").
    Order("count(1) desc").
    Limit(1).
    Scan(&postID)

  var post Post
  db.Preload("Comments").
    First(&post, postID)
  return &post
}

// 新增文章
func addPost(title, content string, userID uint, db *gorm.DB) *Post {
  post := Post{Title: title, Content: content, UserID: userID, CommentStatus: "无评论"}
  db.Create(&post)
  return &post
}

// 删除文章
func deletePost(id uint, db *gorm.DB) {
  var post Post
  db.First(&post, id)
  db.Delete(&post)
}

// 新增评论
func addComment(content string, postID uint, db *gorm.DB) *Comment {
  comment := Comment{Content: content, PostID: postID}
  db.Create(&comment)
  return &comment
}

// 删除评论
func deleteComment(id uint, db *gorm.DB) {
  var comment Comment
  db.First(&comment, id)
  db.Delete(&comment)
}

func initDataBase(db *gorm.DB) {
  db.AutoMigrate(&User{})
  db.AutoMigrate(&Post{})
  db.AutoMigrate(&Comment{})

  comments1 := []Comment{
    {Content: "牛逼"},
    {Content: "666"},
  }
  comments2 := []Comment{
    {Content: "厉害了"},
  }

  posts1 := []Post{
    {Title: "A", Content: "AA", Comments: comments1},
    {Title: "B", Content: "BB", Comments: comments2},
  }
  posts2 := []Post{
    {Title: "C", Content: "CC"},
  }

  users := []*User{
    {
      UserName: "zhp",
      Email:    "zhp.gmail.com",
      Posts:    posts1,
    },
    {
      UserName: "张三",
      Email:    "zhangsan.gmail.com",
      Posts:    posts2,
    },
  }
  db.Create(&users)

  var ret []User
  db.Preload("Posts.Comments").
    Preload(clause.Associations).
    Find(&ret)
  fmt.Println("初始化所有数据：")
  for _, user := range ret {
    fmt.Println(user)
  }
  fmt.Println("=====================")
  fmt.Println()
}
