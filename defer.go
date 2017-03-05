package main

import "fmt"

func fn(n int) int {
	defer func() {
		n++
		fmt.Println("3st:", n)
	}()
	defer func() {
		n++
		fmt.Println("2st:", n)
	}()
	defer func() {
		n++
		fmt.Println("1st:", n)
	}()
	return n //没有做任何事情
}
func main() {
	fmt.Println("函数返回值：", fn(0))
}
