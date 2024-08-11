package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastName"`
}

var movies []Movie

func main() {
	r := mux.NewRouter() // for routing
	// these are methods
	movies = append(movies, Movie{ID: "1", Isbn: "4632225", Title: "Movie 1", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	r.HandleFunc("/movies", getMovie).Methods("GET")
	r.HandleFunc("/movies/{ID}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{ID}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{ID}", deleteMovie).Methods("DELETE")
	fmt.Print("Starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
