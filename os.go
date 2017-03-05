package main

import (
	"fmt"
	"os"
)

func main() {
	host, err := os.Hostname()
	if err != nil {
		//		fmt.Printf("%s", err)
	} else {
		//		fmt.Printf("%s", host)
		fmt.Println(host)
	}
	path := os.Getenv("GOPATH")
	fmt.Println("环境变量GOPATH的值是:", path)
	dir, _ := os.Getwd()
	fmt.Println("当前的目录是:", dir)

	fmt.Println(os.Getegid())
	fmt.Println(os.Geteuid())
	fmt.Println(os.Getgid())
	g, _ := os.Getgroups()
	fmt.Println(g)
	fmt.Println(os.Getpagesize())
	fmt.Println(os.Getppid())
	fmt.Println(os.Getuid())

	filemode, _ := os.Stat("io.go")
	fmt.Println(filemode.Mode())

	data := os.Environ()
	fmt.Println(data)
	//	os.Clearenv() //清空环境变量
	//	data = os.Environ()
	//	fmt.Println(data)

	//	f, err := os.Open("c:\\a.txt")
	//	if err != nil {
	//		return nil, err
	//	}
	//	defer f.Close()
}
