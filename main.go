package main

import (
	"io"
	"net/http"
	"os"
	"time"

	"go-web-boilerplate/config"

	"github.com/gorilla/handlers"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)

	file, _ := os.OpenFile("server.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	outs := io.MultiWriter(file, os.Stdout)

	logged := handlers.LoggingHandler(outs, mux)
	server := &http.Server{
		Addr:         config.Cfg.HTTPAddr,
		Handler:      logged,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
