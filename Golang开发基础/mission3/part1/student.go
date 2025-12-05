package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//题目1：基本CRUD操作
//假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
//要求 ：
//编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
//编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
//编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
//编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。

//题目2：事务语句
//假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
//要求 ：
//编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。

type Student struct {
	ID    uint
	Name  string
	Age   int
	Grade string
}

func main() {
	dsn := "root:rootroot@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("连接mysql失败: %v", err)
	}

	err = db.AutoMigrate(&Student{})
	if err != nil {
		log.Fatalf("student建表失败, %v", err)
	}

	//新增一条记录
	db.Debug().Create(&Student{
		Name:  "张三",
		Age:   20,
		Grade: "三年级",
	})

	//查询age>18的所有学生
	var student Student
	db.Debug().Where("age > ?", 18).Find(&student)
	fmt.Println(student)

	//将 name = 张三 的记录年级更新为“四年级”
	db.Debug().Model(&Student{}).Where("name = ?", "张三").Update("grade", "四年级")

	//删除age < 15的所有记录
	db.Debug().Where("age < ?", 15).Delete(&Student{})

}
