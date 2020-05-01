package main

import (
	"net/http"
	"fmt"
)

func NewHandler(handler http.HandlerFunc) (http.HandlerFunc) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("start interceptor")
		handler.ServeHTTP(w, r)
	})
}

func logicHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("start login")
}

func main() {
	mutex := http.NewServeMux()
	mutex.Handle("/api", NewHandler(logicHandler))
	http.ListenAndServeTLS(":9920", "server.crt", "server.key", mutex)
}
