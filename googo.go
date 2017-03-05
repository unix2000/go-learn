package main

import (
	"fmt"
	//	"html"
	//	"html/template"
	//	"net/http"
	"os"

	"math/rand"
	"strconv"
	"strings"
	"time"

	"./trans"
)

var v1 int
var v2 string
var v3 [10]int //数组
var v4 []int   //数组切片
var v5 struct {
	f int
}
var v6 *int           //指针
var v7 map[string]int //map,key为string类型,value为int类型
var v8 func(a int) int

//复合类型
//指针 pointer
//数组 array
//切片 slice
//字典 map
//通道 chan
//结构体 struct
//接口 interface

func a() {
	str := "Hello,世界"
	//	n := len(str)
	//	for i := 1; i < n; i++ {
	//		ch := str[i] //依据下标取字符串中的字符,类型为byte
	//		fmt.Println(i, ch)
	//	}
	for i, ch := range str { //str类型为rune
		fmt.Println(i, ch)
	}
}

var twoPi = 2 * trans.Pi

const ss string = "liner.xie"

var global_a = "global string"
var ch int = '\u0041'
var ch2 int = '\u03B2'
var ch3 int = '\U00101234'

const (
	//关键字iota从0开始自增枚举值
	Sunday    = iota //0
	Monday           //1
	Tuesday          //2
	Wednesday        //3
	Thursday         //4
	Friday           //5
	Saturday         //6
)

const (
	_        = iota
	KB int64 = 1 << (10 * iota)
	MB
	GB
	TB
)

//type
type Color int

const (
	Black Color = iota
	Red
	Blue
)

//匿名函数
//f := func(x,y int) int {
//	return x + y
//}

func n() { print(global_a) }

func m() {
	a := "O"
	print(a)
}
func init() {

}
func test() (int, string) {
	return 1, "liner.xie"
}
func main() {
	//切片使用
	mySlice := make([]int, 5, 10)
	fmt.Println("len(mySlice):", len(mySlice))
	fmt.Println("cap(mySlice):", cap(mySlice))

	a()
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)

	_, sa := test()
	fmt.Println(sa)
	a, b, c := 5, 7, "abc"
	var goos string = os.Getenv("GOOS")
	fmt.Printf("The OS is: %s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)
	fmt.Printf("a:%d,b:%d,c:%s", a, b, c)
	fmt.Printf("2*Pi = %g\n", twoPi)
	fmt.Printf("const ss string is:%s \n", ss)
	m()
	n()
	for i := 0; i < 10; i++ {
		a := rand.Int()
		fmt.Printf("%d / ", a)
	}
	for i := 0; i < 5; i++ {
		r := rand.Intn(8)
		fmt.Printf("%d / ", r)
	}
	fmt.Println()
	timens := int64(time.Now().Nanosecond())
	rand.Seed(timens)
	for i := 0; i < 10; i++ {
		fmt.Printf("%2.2f / ", 100*rand.Float32())
	}

	fmt.Printf("%d - %d - %d\n", ch, ch2, ch3) // integer
	fmt.Printf("%c - %c - %c\n", ch, ch2, ch3) // character
	fmt.Printf("%X - %X - %X\n", ch, ch2, ch3) // UTF-8 bytes
	fmt.Printf("%U - %U - %U", ch, ch2, ch3)   // UTF-8 code point

	// := 不能定义在func外,内部简略申明变量
	st := "中文hel" + "lo,"
	st += "world!中文"
	fmt.Println(st)

	//var str string = "This is an example of a string"
	//	str := "This is an example of a string"
	//	fmt.Printf("T/F? Does the string \"%s\" have prefix %s? ", str, "Th")
	//	fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))
	str := "The quick brown fox jumps over the lazy dog"
	sl := strings.Fields(str)
	fmt.Printf("Splitted in slice: %v\n", sl)
	for _, val := range sl {
		fmt.Printf("%s - ", val)
	}
	fmt.Println()
	str2 := "GO1|The ABC of Go|25"
	sl2 := strings.Split(str2, "|")
	fmt.Printf("Splitted in slice: %v\n", sl2)
	for _, val := range sl2 {
		fmt.Printf("%s - ", val)
	}
	fmt.Println()
	str3 := strings.Join(sl2, ";")
	fmt.Printf("sl2 joined by ;: %s\n", str3)

	//字符串与其它类型的转换strconv
	var orig string = "666"
	var an int
	var newS string

	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)

	an, _ = strconv.Atoi(orig)
	fmt.Printf("The integer is: %d\n", an)
	an = an + 5
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", newS)

	//指针
	var i1 = 5
	fmt.Printf("An integer: %d, its location in memory: %p\n", i1, &i1)
	var intP *int
	intP = &i1
	fmt.Printf("The value at memory location %p is %d\n", intP, *intP)

	s := "good bye"
	var p *string = &s
	*p = "ciao"
	fmt.Printf("Here is the pointer p: %p\n", p)  // prints address 内存地址
	fmt.Printf("Here is the string *p: %s\n", *p) // prints string
	fmt.Printf("Here is the string s: %s\n", s)   // prints same string

	k := 6
	switch k {
	case 4:
		fmt.Println("was <= 4")
		fallthrough
	case 5:
		fmt.Println("was <= 5")
		fallthrough
	case 6:
		fmt.Println("was <= 6")
		fallthrough
	case 7:
		fmt.Println("was <= 7")
		fallthrough
	case 8:
		fmt.Println("was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}
}
