package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "gorm.io/driver/mysql"
)

//题目2：实现类型安全映射
//假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
//要求 ：
//定义一个 Book 结构体，包含与 books 表对应的字段。
//编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。

type Book struct {
	ID     uint
	Title  string
	Author string
	Price  uint
}

func main() {

	dsn := "root:rootroot@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalln("连接MySQL失败")
	}

	initDataBase2(db)

	//1、条件查询
	fmt.Println("=======================")
	rows, err := db.Queryx("select * from books where price between ? and ?", 20, 50)
	if err != nil {
		log.Fatalln("查询失败,err", err)
	}

	fmt.Println("1、查询价格20~50的图书有：")
	for rows.Next() {
		var book Book
		rows.StructScan(&book)
		fmt.Println(book)
	}
	fmt.Println("=======================")
	fmt.Println()

	//2、更新记录
	_, err = db.Exec("update books set author = ? where title = ?", "加西亚·马尔克斯", "百年孤独")
	if err != nil {
		fmt.Println("更新失败", err)
	}
	fmt.Println("2、更新成功")
	fmt.Println("=======================")
	fmt.Println()

	//3、删除记录
	_, err = db.Exec("delete from books where author = ?", "qwe")
	if err != nil {
		fmt.Println("删除失败", err)
	}
	fmt.Println("3、删除成功")
	fmt.Println("=======================")
	fmt.Println()
}

// 初始化数据库
func initDataBase2(db *sqlx.DB) {
	schema := `create table books
            (
                id     bigint unsigned auto_increment
                    primary key,
                title  varchar(100)        null,
                author varchar(50)        null,
                price  bigint unsigned null
            );`
	db.Exec(schema)

	//初始化一批数据
	db.Exec("INSERT INTO db1.books (id, title, author, price) VALUES (1, 'Java编程', 'abc', 20),(2, '百年孤独', '马克', 30),(3, '追风筝的人', 'yy', 25),(4, 'go编程', '哈哈哈', 999),(5, '吹牛逼呢', 'qwe', 10);")
	var books []Book
	err := db.Select(&books, "select * from books")
	if err != nil {
		log.Fatalln("查询初始化数据失败,", err)
	}
	fmt.Println("初始化数据如下：")
	for _, b := range books {
		fmt.Println(b)
	}
}
