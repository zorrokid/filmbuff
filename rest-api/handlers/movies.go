package handlers

import (
	"log"
	"net/http"

	"github.com/zorrokid/film-db-rest-api/data"
	"github.com/zorrokid/film-db-rest-api/data/models"
)

type Movies struct {
	logger     *log.Logger
	repository data.IMoviesRepository
}

func NewMovies(logger *log.Logger, repository data.IMoviesRepository) *Movies {
	return &Movies{logger, repository}
}

func (m *Movies) GetMovies(rw http.ResponseWriter, r *http.Request) {
	m.logger.Println("Handle GET movies")
	movies := m.repository.GetMovies()

	err := movies.ToJson(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal JSON", http.StatusInternalServerError)
	}
}

type KeyMovie struct{}

func (m *Movies) AddMovie(rw http.ResponseWriter, r *http.Request) {
	m.logger.Println("Handle POST movies")
	movie := r.Context().Value(KeyMovie{}).(models.Movie)
	m.repository.AddMovie(&movie)
}
