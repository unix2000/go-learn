package main

import (
	"encoding/json"
	"fmt"
	//	"log"
)

//type Movie struct {
//	Title  string
//	Year   int  `json:"released"`
//	Color  bool `json:"color,omitempty"`
//	Actors []string
//}

//func main() {
//	var movies = []Movie{
//		{Title: "Casablanca", Year: 1942, Color: false,
//			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
//		{Title: "Cool Hand Luke", Year: 1967, Color: true,
//			Actors: []string{"Paul Newman"}},
//		{Title: "Bullitt", Year: 1968, Color: true,
//			Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
//	}
//	data, err := json.Marshal(movies)
//	if err != nil {
//		log.Fatalf("JSON marshaling failed: %s", err)
//	}
//	fmt.Printf("%s\n", data)
//}

type Person struct {
	Name        string `json:"username"`
	Age         int
	Gender      bool `json:",omitempty"`
	Profile     string
	OmitContent string `json:"-"`
	Count       int    `json:",string"`
}

func main() {
	var p *Person = &Person{
		Name:        "brainwu",
		Age:         21,
		Gender:      true,
		Profile:     "I am Wujunbin",
		OmitContent: "OmitConent",
	}
	if bs, err := json.Marshal(&p); err != nil {
		panic(err)
	} else {
		//result --> {"username":"brainwu","Age":21,"Gender":true,"Profile":"I am Wujunbin","Count":"0"}
		fmt.Println(string(bs))
	}

	var p2 *Person = &Person{
		Name:        "brainwu",
		Age:         21,
		Profile:     "I am Wujunbin",
		OmitContent: "OmitConent",
	}
	if bs, err := json.Marshal(&p2); err != nil {
		panic(err)
	} else {
		//result --> {"username":"brainwu","Age":21,"Profile":"I am Wujunbin","Count":"0"}
		fmt.Println(string(bs))
	}

	// slice 序列化为json
	var aStr []string = []string{"Go", "Java", "Python", "Android"}
	if bs, err := json.Marshal(aStr); err != nil {
		panic(err)
	} else {
		//result --> ["Go","Java","Python","Android"]
		fmt.Println(string(bs))
	}

	//map 序列化为json
	var m map[string]string = make(map[string]string)
	m["Go"] = "No.1"
	m["Java"] = "No.2"
	m["C"] = "No.3"
	if bs, err := json.Marshal(m); err != nil {
		panic(err)
	} else {
		//result --> {"C":"No.3","Go":"No.1","Java":"No.2"}
		fmt.Println(string(bs))
	}
}
