package router

import (
	"github.com/gorilla/mux"
	"github.com/jesseemana/movieapi/controllers"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controllers.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movies", controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controllers.MarkWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controllers.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/api/movies/delete_all", controllers.DeleteAllMovies).Methods("DELETE")

	return router
}
