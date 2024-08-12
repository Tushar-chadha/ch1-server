package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

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

func getMoviesList(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, MovieItem := range movies {
		if MovieItem.ID == params["ID"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)

}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, MovieItem := range movies {
		if MovieItem.ID == params["ID"] {
			json.NewEncoder(w).Encode(MovieItem) // Encode the matched movie
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("Movie not found")
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	var movie Movie
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(r)
	return
}
func updateMovie(w http.ResponseWriter, r *http.Request) {
	var updatedMovie Movie
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	err := json.NewDecoder(r.Body).Decode(&updatedMovie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for index, movieItem := range movies {
		if movieItem.ID == params["ID"] {

			movies = append(movies[:index], movies[index+1:]...)

			updatedMovie.ID = params["ID"]

			if updatedMovie.Director == nil {
				updatedMovie.Director = &Director{}
			}
			updatedMovie.Director.Firstname = "Tushar"

			movies = append(movies, updatedMovie)

			json.NewEncoder(w).Encode(updatedMovie)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("Movie not found")
}

func main() {
	r := mux.NewRouter() // for routing
	// these are methods
	movies = append(movies, Movie{ID: "1", Isbn: "4632225", Title: "Movie 1", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "4632225", Title: "Movie 1", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	r.HandleFunc("/movies", getMoviesList).Methods("GET")
	r.HandleFunc("/movies/{ID}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{ID}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{ID}", deleteMovie).Methods("DELETE")
	fmt.Print("Starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
