package main

/*
Try out some samples from:
http://blog.scottlogic.com/2017/02/28/building-a-web-app-with-go.html
 */


import (
	"fmt"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8080", nil)
}
