package main

import (
	"fmt"
	"net/http"
	"strconv"
)

type Genre string

const (
	Action  Genre = "Action"
	Sci_Fi  Genre = "Sci-Fi"
	Western Genre = "Western"
	Comedy  Genre = "Comedy"
	Romance Genre = "Romance"
)

type Movie struct {
	Title string `json:"title"`
	Genre Genre  `json:"genre"`
	Year  int64  `json:"year"`
}

func createMovieHandler(w http.ResponseWriter, r *http.Request) {
	var newMovie Movie

	newMovie.Title = r.FormValue("title")
	newMovie.Genre = Genre(r.FormValue("genre"))
	yearStr := r.FormValue("year")

	year, err := strconv.ParseInt(yearStr, 10, 64)
	if err != nil {
		http.Error(w, "failed to convert yearStr to int64", http.StatusBadRequest)
	}

	newMovie.Year = year

	fmt.Fprintf(w, "New movie creared: %v", newMovie)

}

func main() {

	http.HandleFunc("POST /movies", createMovieHandler)
	mux := http.NewServeMux()
	fmt.Println("Listening....")
	http.ListenAndServe(":8080", mux)

}
