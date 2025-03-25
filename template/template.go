package main

import (
	"net/http"
	"text/template"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {

	tmpl := template.Must(template.ParseFiles("layout.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "my todo list",
			Todos: []Todo{
				{Title: "吃饭", Done: true},
				{Title: "睡觉", Done: false},
				{Title: "打豆豆", Done: false},
				{Title: "拉屎", Done: false},
			},
		}
		tmpl.Execute(w, data)
	})

	http.ListenAndServe(":8080", nil)
}
