package main

import "net/http"

func main() {

	fileServer := http.FileServer(http.Dir("D:\\www\\eui\\dist"))
	http.Handle("/", fileServer)
	http.ListenAndServe(":9999", nil)
}
