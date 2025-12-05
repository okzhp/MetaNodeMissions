package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "gorm.io/driver/mysql"
)

////题目1：使用SQL扩展库进行查询
////假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
////要求 ：
////编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
////编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。

type Employee struct {
	ID         uint
	Name       string
	Department string
	Salary     uint
}

func main() {

	dsn := "root:rootroot@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalf("连接mysql失败")
	}

	initDataBase1(db)

	//1、查询所有“技术部”的员工信息
	var employees []Employee
	err = db.Select(&employees, "select * from employees where department = ?", "技术部")
	if err != nil {
		log.Fatalln("查询失败，err: ", err)
	}
	fmt.Println("=============")
	fmt.Println("1、查询技术部员工有:")
	for _, e := range employees {
		fmt.Println(e)
	}
	fmt.Println("=============")

	//2、查询工资最高的员工信息
	var highestSalaryEmployee Employee
	err = db.Get(&highestSalaryEmployee, "select * from employees order by salary desc limit 1")
	if err != nil {
		log.Fatalln("查询失败，err: ", err)
	}
	fmt.Println("=============")
	fmt.Println("2、工资最高的员工是:")
	fmt.Println(highestSalaryEmployee)
	fmt.Println("=============")

}

// 初始化数据库
func initDataBase1(db *sqlx.DB) {

	schema := `create table employees
            (
                id         bigint unsigned auto_increment
                    primary key,
                name       varchar(50)        null,
                department varchar(50)        null,
                salary     bigint unsigned null
            );`
	db.Exec(schema)

	//初始化一批数据
	db.Exec("INSERT INTO db1.employees (id, name, department, salary) VALUES (1, '张三', '技术部', 3500),(2, '李四', '技术部', 5000),(3, '小红', '运营部', 3000),(4, '小明', '市场部', 5600),(5, 'zhp', '技术部', 7000)")

	var employees []Employee
	err := db.Select(&employees, "select * from employees")
	if err != nil {
		log.Fatalln("查询初始化数据失败,", err)
	}
	fmt.Println("初始化数据如下：")
	for _, e := range employees {
		fmt.Println(e)
	}
}
