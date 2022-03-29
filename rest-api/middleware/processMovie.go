package middleware

import (
	"context"
	"net/http"

	"github.com/zorrokid/film-db-rest-api/data/models"
	"github.com/zorrokid/film-db-rest-api/handlers"
)

func ProcessMovie(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		movie := models.Movie{}
		err := movie.FromJson(r.Body)
		if err != nil {
			http.Error(rw, "Unable to  marshal JSON", http.StatusBadRequest)
		}

		ctx := context.WithValue(r.Context(), handlers.KeyMovie{}, movie)
		req := r.WithContext(ctx)
		next.ServeHTTP(rw, req)
	})
}
