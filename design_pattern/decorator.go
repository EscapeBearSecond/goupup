package main

import (
	"log"
	"net/http"
	"time"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello world"))
}
func NiHaoHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("你好"))
}

type Handler func(w http.ResponseWriter, r *http.Request)

func LoggerHandler(handler Handler) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		handler(w, r)
		log.Printf("request: %s %s %s %s", r.Method, r.URL.Path, r.RemoteAddr, time.Since(now))
	}
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", LoggerHandler(HelloHandler))
	mux.HandleFunc("/nihao", LoggerHandler((NiHaoHandler)))
	srv := http.Server{
		Addr:    ":8008",
		Handler: mux,
	}
	srv.ListenAndServe()
}
