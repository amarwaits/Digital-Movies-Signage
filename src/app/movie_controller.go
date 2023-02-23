package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/amarwaits/Syndica-THT/src/model"
)

type MovieController struct {
	movieService *MovieService
}

func NewMovieController(movieService *MovieService) *MovieController {
	return &MovieController{
		movieService: movieService,
	}
}

func (c *MovieController) AddMovie(r *http.Request, data *model.MovieInput, result *int) error {
	*result = c.movieService.AddMovie(*data)
	return nil
}

func (c *MovieController) GetMovie(r *http.Request, data *model.MovieID, result *model.Movie) error {
	var err error
	*result, err = c.movieService.GetMovie(data.ID)
	return err
}

func (c *MovieController) GetMovies(r *http.Request, data *model.ForceGetMovies, result *[]model.Movie) error {
	fmt.Println(r.Context())

	// w := r.Context().Value("http.ResponseWriter").(http.ResponseWriter)
	// w.Header().Set("Content-Type", "application/json")

	// Generate ETag based on the content of movies and last modified time
	etag := c.movieService.GetETag()

	if !data.Force {
		// Check If-None-Match and If-Modified-Since headers to see if content has been modified
		if match := r.Header.Get("If-None-Match"); match != "" && match == etag {
			// w.WriteHeader(http.StatusNotModified)
			return nil
		}
		if modifiedSince := r.Header.Get("If-Modified-Since"); modifiedSince != "" {
			t, err := time.Parse(http.TimeFormat, modifiedSince)
			if err == nil && !t.Before(c.movieService.GetLastModified()) {
				// w.WriteHeader(http.StatusNotModified)
				return nil
			}
		}
	}

	*result = c.movieService.GetMovies()

	// w.Header().Set("ETag", etag)
	// w.Header().Set("Last-Modified", c.movieService.GetLastModified().UTC().Format(http.TimeFormat))

	return nil
}
