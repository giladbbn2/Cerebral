package main

import (
	"fmt"
	"io/ioutil"
)

func handler(w http.ResponseWriter, r *http.Request) {
    //fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
    fmt.Fprintf(w, "<!DOCTYPE html><html><head></head><body><h1>Hello</h1><div>ggg</div></body></html>");
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":80", nil)
}