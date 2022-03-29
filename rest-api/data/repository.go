package data

import (
	"log"

	"github.com/zorrokid/film-db-rest-api/data/models"
	"github.com/zorrokid/film-db-rest-api/data/testdata"
	"github.com/zorrokid/film-db-rest-api/database"
)

type IMoviesRepository interface {
	GetMovies() models.Movies
	AddMovie(movie *models.Movie)
}

func NewMoviesTestDataRepository(logger *log.Logger) IMoviesRepository {
	return &MoviesTestDataRepository{logger}
}

type MoviesTestDataRepository struct {
	logger *log.Logger
}

func (mr *MoviesTestDataRepository) GetMovies() models.Movies {
	return testdata.GetMovies()
}

func (mr *MoviesTestDataRepository) AddMovie(movie *models.Movie) {
	testdata.AddMovie(movie)
}

func NewMoviesDataRepository(logger *log.Logger, db *database.Database) IMoviesRepository {
	return &MoviesDataRepository{logger: logger, db: db}
}

type MoviesDataRepository struct {
	db     *database.Database
	logger *log.Logger
}

func (mr *MoviesDataRepository) GetMovies() models.Movies {
	var movies []*models.Movie
	mr.db.GetConnection().Find(&movies)
	return movies
}

func (mr *MoviesDataRepository) AddMovie(movie *models.Movie) {
	mr.db.GetConnection().Create(movie)
}
