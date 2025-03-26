package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func secret(w http.ResponseWriter, r *http.Request) {
	sessions, _ := store.Get(r, "cookie-name")

	if auth, ok := sessions.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	fmt.Fprintln(w, "The cake is a lie!")
}

func login(w http.ResponseWriter, r *http.Request) {
	sessions, _ := store.Get(r, "cookie-name")

	//"cookie-name" 名字是可见的，值是用"super-secret-key"签名加密的，如 MTc0Mjk3OTg1MnxEWDhFQVFMX2dBQUJF.....
	//解密后里面有 authenticated 信息
	sessions.Values["authenticated"] = true
	sessions.Save(r, w)
	fmt.Fprintln(w, "You are logged in")
}

func logout(w http.ResponseWriter, r *http.Request) {
	sessions, _ := store.Get(r, "cookie-name")

	log.Println(sessions.Values)
	sessions.Values["authenticated"] = false
	sessions.Save(r, w)
}

func main() {
	http.HandleFunc("/secret", secret)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)

	http.ListenAndServe(":8080", nil)
}
