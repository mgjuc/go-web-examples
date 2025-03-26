package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func logging() Middleware {
	// Create a new Middleware
	return func(f http.HandlerFunc) http.HandlerFunc {
		// Define the http.HandlerFunc that will be called for each request.
		return func(w http.ResponseWriter, r *http.Request) {
			// Do middleware things
			start := time.Now()

			defer func() { log.Println(r.Method, r.RequestURI, time.Since(start)) }()

			// Call the next middleware handler
			f(w, r)
		}
	}
}

func Method(m string) Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if r.Method != m {
				http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
				return
			}
			f(w, r)
		}
	}
}

func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	// Apply each middleware to the handler in reverse order
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {

	//注意这个 log 和 Method 的顺序，会影响log的输出：这里是从后往前执行的，POST请求会被Method拦截，不会执行log
	http.HandleFunc("/", Chain(Hello, logging(), Method("GET")))
	http.ListenAndServe(":8080", nil)

}
