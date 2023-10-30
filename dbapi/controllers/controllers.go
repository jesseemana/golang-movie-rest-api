package controllers

import (
	"encoding/json"
	"net/http"

	model "github.com/jesseemana/gomongoapi/models"
	"github.com/jesseemana/gomongoapi/service"
)

func GetAllMovies(w http.ResponseWriter, r *http.Response) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := services.GetAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Response) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Movies
	_ = json.NewDecoder(r.Body).Decode(movie)
	services.InsertMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkWatched(w http.ResponseWriter, r *http.Response) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	services.UpdateMovie(params["id"])
	json.NewEncoder(w).Encode("Movie marked as watched")
}

func DeleteOneMovie(w http.ResponseWriter, r *http.Response) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	services.DeleteOneMovie(params["id"])
	json.NewEncoder(w).Encode("Movie deleted")
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Response) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := services.DeleteAllMovies()
	json.NewEncoder(w).Encode(count)
}
