package main

import (
	"math/rand"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director ` json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params  := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}


func main() {
	routes := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "Scary Movie", Director: &Director{Firstname: "John", Lastname: "Ufdo"}})
	movies = append(movies, Movie{ID: "2", Isbn: "585793", Title: "Superhero", Director: &Director{Firstname: "Alex", Lastname: "Crasinski"}})

	// routes
	routes.HandleFunc("/movies", getMovies).Methods("GET")
	routes.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	routes.HandleFunc("/movies", createMovie).Methods("POST")
	routes.HandleFunc("/movie/{id}", updateMovie).Methods("PUT")
	routes.HandleFunc("/movie/{id}", deleteMovie).Methods("DELETE")
	
	
	// server
	fmt.Printf("starting server on port 3000\n")
	log.Fatal(http.ListenAndServe(":3000", routes))

}
