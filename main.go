package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"Id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "43877", Title: "ABCD", Director: &Director{FirstName: "Ashish", LastName: "Sawant"}})
	movies = append(movies, Movie{ID: "2", Isbn: "43887", Title: "EFGH", Director: &Director{FirstName: "Rutvik", LastName: "Mohite"}})
	movies = append(movies, Movie{ID: "3", Isbn: "58887", Title: "XYZ", Director: &Director{FirstName: "Satyajeet", LastName: "Barale"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("GET")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("GET")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("GET")

	fmt.Printf("Starting serer at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))

}
