package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Fprintf(writer, string) ==>向指定的writer输出
		fmt.Fprintf(w, "welcom to my server")
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err)
	}
}
