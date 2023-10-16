package main

import (
	"github.com/diyor200/movie-microservice/rating/internal/controller/rating"
	httphandler "github.com/diyor200/movie-microservice/rating/internal/handler/handler/http"
	"github.com/diyor200/movie-microservice/rating/internal/repository/memory"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting the rating service")
	repo := memory.NewRepository()
	ctrl := rating.NewController(repo)
	h := httphandler.NewController(ctrl)
	http.Handle("/rating", http.HandlerFunc(h.Handle))
	if err := http.ListenAndServe(":8082", nil); err != nil {
		panic(err)
	}
}
