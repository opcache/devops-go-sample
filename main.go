package main

import (
	"fmt"
	"net/http"
)

func HelloWorld(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "HTTP1")
}

func main() {
	http.HandleFunc("/", HelloWorld)
	http.ListenAndServe(":80", nil)
}
