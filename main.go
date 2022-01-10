package main

import (
	"go-blog/internal/routers"
	"net/http"
	"time"
)

func main() {
	router := routers.NewRouter()

	server := &http.Server{
		Addr:           ":9090",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
