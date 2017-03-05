package main

import (
	"fmt"
)

//type person struct {
//	Name    string
//	Age     int
//	Contact struct {
//		Phone, Email string
//	}
//}

//struct类似于class,go没有继承
type human struct {
	Sex int
}
type teacher struct {
	human
	Name string
	Age  int
}

type student struct {
	human
	Name string
	Age  int
}

func main() {
	//	a := person{Name: "liner", Age: 20}
	//	a.Contact.Phone = "13600020001"
	//	a.Contact.Email = "liner.xie@qq.com"
	//	fmt.Println(a)

	a := teacher{Name: "liner.xie", Age: 20, human: human{Sex: 1}}
	b := student{Name: "xie", Age: 21, human: human{Sex: 2}}
	a.Name = "xiaolin"
	a.Age = 22
	a.Sex = 1
	fmt.Println(a, b)
	a.Print()

	//闭包使用
	f := closure(20)
	fmt.Println(f(11))
	fmt.Println(f(12))
}

//struct 函数定义,类似于class中的函数
func (t teacher) Print() {
	fmt.Println("teacher print func")
}

//go闭包
func closure(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

//go接口

//go切片

//go map
