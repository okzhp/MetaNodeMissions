package main

import (
	"fmt"
	"math"
)

func main() {
	rectangle := Rectangle{long: 20, width: 10}
	fmt.Printf("rectangle: %#v ,面积：%f \n", rectangle, rectangle.Area())
	fmt.Printf("rectangle: %#v ,周长：%f \n", rectangle, rectangle.Perimeter())

	fmt.Println("===========================================")

	circle := &Circle{radius: 5}
	fmt.Printf("circle: %#v ,面积：%f \n", circle, circle.Area())
	fmt.Printf("circle: %#v ,周长：%f \n", circle, circle.Perimeter())

	fmt.Println("===========================================")

	employee := Employee{Person: Person{Name: "zhp", Age: 100}, EmployeeID: 1}
	employee.PrintInfo()

	fmt.Println("===========================================")
}

// 题目 ：
// 定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。
// 然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
// 在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
type shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	long  float64
	width float64
}

func (r *Rectangle) Area() float64 {
	return r.long * r.width
}

func (r *Rectangle) Perimeter() float64 {
	return (r.long + r.width) * 2
}

type Circle struct {
	radius float64
}

func (c *Circle) Area() float64 {
	return 3.14 * math.Pow(c.radius, 2)
}

func (c *Circle) Perimeter() float64 {
	return 2 * 3.14 * c.radius
}

// 题目 ：
// 使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，
// 再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。
// 为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID int
}

func (employee *Employee) PrintInfo() {
	fmt.Printf("员工ID：%d，员工姓名：%s，员工年龄：%d \n", employee.EmployeeID, employee.Name, employee.Age)
}
