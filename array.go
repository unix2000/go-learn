package main

import (
	"fmt"
)

//var s1 []int //slice
func byteString(p []byte) string {
	for i := 0; i < len(p); i++ {
		if p[i] == 0 {
			return string(p[0:i])
		}
	}
	return string(p)
}

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	b := [10]int{1, 2, 3}
	c := [...]int{3, 4, 5}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c[2])

	//	doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}
	//	easyArray := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	//	fmt.Println(doubleArray)
	//	fmt.Println(doubleArray[1][2])

	//slice
	slice := []byte{'a', 'b', 'c', 'd'}
	fmt.Println(slice[1:3])

	// 声明一个数组
	var array = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	// 声明两个slice
	var aSlice, bSlice []byte

	// 演示一些简便操作
	aSlice = array[:3] // 等价于aSlice = array[0:3] aSlice包含元素: a,b,c
	aSlice = array[5:] // 等价于aSlice = array[5:10] aSlice包含元素: f,g,h,i,j
	aSlice = array[:]  // 等价于aSlice = array[0:10] 这样aSlice包含了全部的元素

	// 从slice中获取slice
	aSlice = array[3:7]  // aSlice包含元素: d,e,f,g，len=4，cap=7
	bSlice = aSlice[1:3] // bSlice 包含aSlice[1], aSlice[2] 也就是含有: e,f
	bSlice = aSlice[:3]  // bSlice 包含 aSlice[0], aSlice[1], aSlice[2] 也就是含有: d,e,f
	bSlice = aSlice[0:5] // 对slice的slice可以在cap范围内扩展，此时bSlice包含：d,e,f,g,h
	bSlice = aSlice[:]   // bSlice包含所有aSlice的元素: d,e,f,g
	fmt.Println(aSlice)
	fmt.Println(bSlice)

	str := byteString(aSlice)
	fmt.Println(str)

	s2 := array[2:4:8]
	fmt.Println(byteString(s2))
	fmt.Println(array)
	fmt.Println(len(s2))
	fmt.Println(cap(s2))
	fmt.Println(cap(array))

	//map
	//	var numbers map[string]int
	numbers := make(map[string]string)
	numbers["one"] = "a"
	numbers["two"] = "b"
	numbers["three"] = "c"
	fmt.Println(numbers)
	fmt.Println("the three is:", numbers["three"])
	k, v := numbers["one"]
	fmt.Println(k)
	fmt.Println(v) //true

	//	make用于内建类型（map、slice 和channel）的内存分配。new用于各种类型的内存分配
	//	new返回指针
	//	make返回初始化后的（非零）值

}
