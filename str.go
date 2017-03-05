package main

import (
	"fmt"
	"strings"
)

func main() {
	var s []string
	str := "liner.xie"
	//	s := strings.Split(str, ".")
	s = strings.Split(str, ".")
	fmt.Println(s[0])
	fmt.Println(strings.Contains("widuu", "wi"))          //true
	fmt.Println(strings.Contains("wi", "widuu"))          //false
	fmt.Println(strings.ContainsAny("widuu", "w&d"))      //true
	fmt.Println(strings.ContainsRune("widuu", rune('w'))) //true
	fmt.Println(strings.ContainsRune("widuu", 20))        //fasle
	fmt.Println(strings.Count("widuu", "uu"))             //1
	fmt.Println(strings.Count("widuu", "u"))              //2
	fmt.Println(strings.Index("widuu", "i"))              //1
	fmt.Println(strings.Index("widuu", "u"))              //3
	fmt.Println(strings.IndexAny("widuu", "u"))           //3
	fmt.Println(strings.IndexByte("hello xiaowei", 'x'))  //6
	fmt.Println(strings.IndexRune("widuu", rune('w')))    //0
	fmt.Println(strings.IndexFunc("nihaoma", split))      //3
	fmt.Println(strings.LastIndex("widuu", "u"))          // 4
	fmt.Println(strings.LastIndexAny("widuu", "u"))       // 4
}
func split(r rune) bool {
	if r == 'a' {
		return true
	}
	return false
}
