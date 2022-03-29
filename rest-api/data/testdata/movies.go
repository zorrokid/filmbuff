package testdata

import "github.com/zorrokid/film-db-rest-api/data/models"

func GetMovies() models.Movies {
	return movieList
}

func AddMovie(movie *models.Movie) {
	movieList = append(movieList, movie)
}

var movieList = []*models.Movie{{
	Name: "Zorro",
}, {
	Name: "Star Wars",
}}
