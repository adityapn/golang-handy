package main

import("io"
		"net/http")

func homePage(w http.ResponseWriter,r *http.Request) {
	io.WriteString(w,"<h1>Home</h1>")
}

func main() {
	http.HandleFunc("/",homePage)
	http.ListenAndServe(":8080",nil)
}