package main

import (
	"fmt"
)

func main() {

	a := Rectangle{}
	b := Circle{}
	var sharpArr []Shape
	sharpArr = append(sharpArr, a)
	sharpArr = append(sharpArr, b)
	for _, sharp := range sharpArr {
		sharp.Area()
		sharp.Perimeter()
	}
	var e0 = Employee{
		1002,
		Person{
			"李四", // 按顺序给Name赋值
			30,   // 按顺序给Age赋值
		},
	}
	e0.PrintInfo()

	//结构提定义1
	e1 := Employee{
		EmployeeID: 123,
		Person: Person{
			Name: "tom",
			Age:  18,
		},
	}
	e1.PrintInfo()

	// 初始化方式2：简洁写法
	e2 := Employee{
		EmployeeID: 1002,
		Person: Person{
			"李四",
			30,
		},
	}
	e2.PrintInfo()
	e3 := Employee{EmployeeID: 1003}
	e3.Name = "王五"
	e3.Age = 28
	e3.PrintInfo()

	var e4 = new(Employee)
	e4.Name = "王五"
	e4.Age = 28
	e4.PrintInfo()

	var e5 Employee
	e5.Name = "赵六"
	e5.Age = 28
	e5.PrintInfo()

}

// Shape 形状接口
type Shape interface {
	Area()
	Perimeter()
}

type Rectangle struct {
}

func (r Rectangle) Area() {
	fmt.Println("Area of Rectangle")
}

func (r Rectangle) Perimeter() {
	fmt.Println("Perimeter of Rectangle")
}

type Circle struct {
}

func (r Circle) Area() {
	fmt.Println("Area of Circle")
}
func (r Circle) Perimeter() {
	fmt.Println("Perimeter of Circle")
}

type Person struct {
	Name string
	Age  int
}
type Employee struct {
	EmployeeID int
	Person
}

func (e Employee) PrintInfo() {
	fmt.Println("Employee Info", e)
	fmt.Printf("name %s\n", e.Name)
	fmt.Printf("age %d\n", e.Age)
	fmt.Printf("EmployeeID %d\n", e.EmployeeID)
}
