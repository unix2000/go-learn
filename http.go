package main

import (
	_ "fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static/")))
	http.ListenAndServe(":8090", nil)
}
