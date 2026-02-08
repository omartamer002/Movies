package handlers

import (
	"encoding/json"
	"math/rand"
	"movies/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Movies)
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parameters := mux.Vars(r)
	for _, item := range models.Movies {
		if item.ID == parameters["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&models.Movie{})
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie models.Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	models.Movies = append(models.Movies, movie)
	json.NewEncoder(w).Encode(models.Movies)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parameters := mux.Vars(r)
	for index, item := range models.Movies {
		if item.ID == parameters["id"] {
			models.Movies = append(models.Movies[:index], models.Movies[index+1:]...)
			var movie models.Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = parameters["id"]
			models.Movies = append(models.Movies, movie)
			json.NewEncoder(w).Encode(models.Movies)
			return
		}
	}
	json.NewEncoder(w).Encode(models.Movies)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parameters := mux.Vars(r)
	for index, item := range models.Movies {
		if item.ID == parameters["id"] {
			models.Movies = append(models.Movies[:index], models.Movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(models.Movies)
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	models.Movies = nil
	json.NewEncoder(w).Encode(models.Movies)
}
