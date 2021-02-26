package main

import (
	"fmt"
	"net/http"
)

func HelloWorld(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello World8")
}

func main() {
	http.HandleFunc("/", HelloWorld)
	http.ListenAndServe(":8080", nil)
}
