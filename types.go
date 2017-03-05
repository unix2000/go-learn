package main

import (
	"fmt"
)

func main() {
	//各种类型写法
	//array 数组
	//	var array [5]int
	//	array := [5]int{1, 2, 3, 4, 5}
	//	array := [...]int{1, 2, 3, 4, 5, 6}

	//	array := [5]int{1: 11, 2: 22}
	//	array[3]为0

	//	指针数组
	//	array := [5]*int{0: new(int), 1: new(int)}
	//	*array[0] = 1
	//	*array[1] = 2

	//赋值
	//	var array1 [5]string
	//	array2 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}
	//	array1 = array2

	//	var array1 [3]*string
	//	array2 := [3]*string{new(string), new(string), new(string)}
	//	*array2[0] = "Red"
	//	*array2[1] = "Blue"
	//	*array2[2] = "Green"
	//	array1 = array2

	//多维数组
	// 声明一个二维数组
	//var array [4][2]int
	// 使用数组字面值声明并初始化
	//array := [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	// 指定外部数组索引位置初始化
	//array := [4][2]int{1: {20, 21}, 3: {40, 41}}
	// 同时指定内外部数组索引位置初始化
	//array := [4][2]int{1: {0: 20}, 3: {1: 41}}

	//切片 slice
	//empty slice
	//	var slice []int
	// 使用 make 创建
	//	silce := make([]int, 0)
	// 使用 slice 字面值创建
	//	slice := []int{}
	// 创建长度和容量都为4的 slice
	//	slice := []int{10, 20, 30, 40}
	// 附加一个新值到 slice，因为超出了容量，所以会创建新的底层数组
	//	newSlice := slice[1:3]
	//	newSlice = append(newSlice, 50)
	//	fmt.Println(slice)
	//	fmt.Println(newSlice)
	//	source := []string{"apple", "orange", "plum", "banana", "grape"}
	// 接着我们在源 slice 之上创建一个新的 slice
	//	slice := source[2:3:4]
	//	fmt.Println(slice)

	//	slice := []int{10, 20, 30, 40}
	//	for index, value := range slice {
	//		fmt.Printf("Index: %d  Value: %d\n", index, value)
	//	}
	//	for _, value := range slice {
	//		fmt.Printf("Value: %d\n", value)
	//	}
	//	for index := 2; index < len(slice); index++ {
	//		fmt.Printf("Index: %d  Value: %d\n", index, slice[index])
	//	}
	//	for index, value := range slice {
	//		fmt.Printf("Value: %d  Value-Addr: %X  ElemAddr: %X\n", value, &value, &slice[index])
	//	}
	//	多维切片 slice
	//slice := [][]int{{10}, {20, 30}}

	//map
	// 通过 make 来创建
	//	dict := make(map[string]int)
	// 通过字面值创建
	//	dict := map[string]string{"Red": "#da1337", "Orange": "#e95a22"}
	//	colors := map[string]string{}
	//	colors["Red"] = "#da1337"

	//	var colors = map[string]string{}
	//	colors["Red"] = "#da1337"

	//nil map 赋值报错
	//	var colors map[string]string
	//	colors["Red"] = "#da1337"
	//	fmt.Println(colors)

	colors := map[string]string{
		"AliceBlue":   "#f0f8ff",
		"Coral":       "#ff7F50",
		"DarkGray":    "#a9a9a9",
		"ForestGreen": "#228b22",
	}
	for key, value := range colors {
		fmt.Printf("Key: %s  Value: %s\n", key, value)
	}
	value, exists := colors["Coral"]
	if exists {
		fmt.Println(value)
	}
}
