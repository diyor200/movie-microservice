package main

import (
	"github.com/diyor200/movie-microservice/metadata/internal/controller/metadata"
	httphandler "github.com/diyor200/movie-microservice/metadata/internal/handler/http"
	"github.com/diyor200/movie-microservice/metadata/internal/repository/memory"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting movie metadata service...")
	repo := memory.NewRepository()
	ctrl := metadata.NewController(repo)
	h := httphandler.NewHandler(ctrl)
	http.Handle("/metadata", http.HandlerFunc(h.GetMetadata))
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}
