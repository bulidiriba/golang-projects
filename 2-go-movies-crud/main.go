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
	Director *Director `json:"director"` // here * shows pointer to struct Director

}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// lets create variable movies that hold list of Movie
var movies []Movie

// implement getMovies function
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

// function to get specific movie
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item) // return the required movie
			return
		}
	}
}

// now implement function to create movie
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie                                // initialize movie
	_ = json.NewDecoder(r.Body).Decode(&movie)     // decode the request body to the address of created movie
	movie.ID = strconv.Itoa(rand.Intn(1000000000)) // create random ID and convert to string
	movies = append(movies, movie)                 // append the created movie to movies
	json.NewEncoder(w).Encode(movie)
}

// implement function to update movie
func updateMovie(w http.ResponseWriter, r *http.Request) {
	// set the json content type
	// get the params
	// loop over the movies until you get the movie the same with the requested one
	// delete the movie
	// add a new movie, with the requested data with the same ID on the place of deleted one
	// then return the updated movie
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
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

// implement deleteMovie
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			// now lets delete the movie, by appending movies from first to current and from next to end
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies) // then return the movies after the deleted one
}

// now lets create main function
func main() {
	r := mux.NewRouter() // lets create router from gorilla mux

	// lets create two Movies manually and then append/add into the variable movies
	// &(ampresand) is used to give the address, but *(asterix) is used to access the address/pointer
	// & -- to create/give the address, * -- to access the address

	var movie1 Movie = Movie{ID: "1", Isbn: "438227", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}}
	var movie2 Movie = Movie{ID: "2", Isbn: "45455", Title: "Movie Two", Director: &Director{Firstname: "Stever", Lastname: "Smith"}}

	movies = append(movies, movie1)
	movies = append(movies, movie2)

	r.HandleFunc("/movies", getMovies).Methods("GET")           // to get all the movies
	r.HandleFunc("/movies", createMovie).Methods("POST")        // to create new movie
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")       // to get given movie with specific id
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")    // to update given movie with specific id
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE") // to delete given movie with specific id

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r)) // this used to listen the router in port 8000, colon(:) is used to put localhost before

}
