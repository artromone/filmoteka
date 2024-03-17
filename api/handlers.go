package main

import (
	// "encoding/json"
	"fmt"
	"log"
	"net/http"
)

func AddActor(w http.ResponseWriter, r *http.Request) {
}

func GetActors(w http.ResponseWriter, r *http.Request) {
}

func UpdateActor(w http.ResponseWriter, r *http.Request) {
}

func DeleteActor(w http.ResponseWriter, r *http.Request) {
}

func AddMovie(w http.ResponseWriter, r *http.Request) {
}

func GetMovies(w http.ResponseWriter, r *http.Request) {
}

func SearchMovies(w http.ResponseWriter, r *http.Request) {
}

func GetActorMovies(w http.ResponseWriter, r *http.Request) {
}

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

func main() {
	http.Handle("/", logger(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Filmoteca API!")
	})))

	http.HandleFunc("/actors", AddActor)
	http.HandleFunc("/actors", GetActors)
	http.HandleFunc("/actors/{actor_id}", UpdateActor)
	http.HandleFunc("/actors/{actor_id}", DeleteActor)
	http.HandleFunc("/movies", AddMovie)
	http.HandleFunc("/movies", GetMovies)
	http.HandleFunc("/movies/search", SearchMovies)
	http.HandleFunc("/actors/{actor_id}/movies", GetActorMovies)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
