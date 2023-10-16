package main

import (
	"github.com/diyor200/movie-microservice/movie/internal/controller/movie"
	metadateGateway "github.com/diyor200/movie-microservice/movie/internal/gateway/metadata/http"
	ratingGateway "github.com/diyor200/movie-microservice/movie/internal/gateway/rating/http"
	httpHandler "github.com/diyor200/movie-microservice/movie/internal/handler/http"
	"log"
	"net/http"
)

func main() {
	log.Println("Staring the movie service")
	mGateway := metadateGateway.NewGateway("localhost:8081")
	rGateway := ratingGateway.NewGateway("localhost:8082")
	ctrl := movie.NewController(rGateway, mGateway)
	handler := httpHandler.NewHandler(ctrl)
	http.Handle("/movie", http.HandlerFunc(handler.GetMovieDetails))
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		panic(err)
	}
}
