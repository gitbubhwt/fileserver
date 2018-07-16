package main

import "net/http"

func main() {
	http.HandleFunc("/file/upload", Upload)
	http.HandleFunc("/file/remove", Remove)
	http.ListenAndServe(":8090", nil)
}
