package main

import (
	"fmt"
	"net/http"
)

func myHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "hello world")
}
func main() {
	http.HandleFunc("/", myHandle)

	http.ListenAndServe(":8888", nil)
}
