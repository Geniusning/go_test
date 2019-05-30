package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func myHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "hello world")
}
func myLoginHandle(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		fmt.Println("username=", r.Form["username"])
		fmt.Println("password=", r.Form["password"])
	}
}
func main() {
	http.HandleFunc("/", myHandle)
	http.HandleFunc("/login", myLoginHandle)

	http.ListenAndServe(":8888", nil)
}
