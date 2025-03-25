package main

import (
	"fmt"
	"net/http"
)

func main() {
	//匿名函数 func () {}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})
	//需要root权限
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		//panic: listen tcp 127.0.0.1:80: bind: permission denied
		panic(err)
	}

}
