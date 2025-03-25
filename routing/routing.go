package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "title:%s page:%s\n", title, page)
		// fmt.Fprintf(w, "it is")
	})
	//80 需要root 权限, 大于1023就不用了
	err := http.ListenAndServe(":8080", r) //这个handle 传 mux.NewRouter()
	if err != nil {
		panic(err)
	}
}
