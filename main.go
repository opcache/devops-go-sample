package main

import (
	"fmt"
	"net/http"
)

func HelloWorld(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello World")
}

func main() {
	http.HandleFunc("/", HelloWorld1)
	http.ListenAndServe(":8080", nil)
}
