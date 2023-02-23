package test

import (
	"reflect"
	"testing"

	"github.com/amarwaits/Syndica-THT/src/app"
	"github.com/amarwaits/Syndica-THT/src/model"
)

func TestAddMovie(t *testing.T) {
	var initialData []model.Movie
	movieService := app.NewMovieService(initialData)

	movie := model.MovieInput{
		Title:  "The Godfather",
		Genre:  "Crime",
		Year:   1972,
		Length: 143,
	}
	expectedNewId := len(movieService.Movies) + 1
	id := movieService.AddMovie(movie)

	if id != 1 {
		t.Fatalf("AddMovie returned wrong id. Expected %d, got %d", expectedNewId, id)
	}
	if len(movieService.Movies) != expectedNewId {
		t.Fatalf("AddMovie did not add movie to the list")
	}
}

func TestGetMovie(t *testing.T) {
	movie := model.Movie{
		Title:  "The Godfather",
		Genre:  "Crime",
		Year:   1972,
		Length: 143,
	}
	initialData := []model.Movie{movie}
	movieService := app.NewMovieService(initialData)

	result, err := movieService.GetMovie(1)

	if err != nil {
		t.Fatalf("GetMovie failed with error: %v", err)
	}
	if !reflect.DeepEqual(movie, result) {
		t.Fatalf("GetMovie returned wrong result. Expected %v, got %v", movie, result)
	}
}

func TestGetMovie_NotFound(t *testing.T) {
	initialData := []model.Movie{}
	movieService := app.NewMovieService(initialData)

	_, err := movieService.GetMovie(1)

	if err == nil {
		t.Fatalf("GetMovie didn't fail with error")
	}
}

func TestGetMovies(t *testing.T) {
	movie := model.Movie{
		Title:  "The Godfather",
		Genre:  "Crime",
		Year:   1972,
		Length: 143,
	}
	initialData := []model.Movie{movie}
	movieService := app.NewMovieService(initialData)

	result := movieService.GetMovies()

	if len(result) != len(initialData) {
		t.Fatalf("GetMovies returned wrong no. of movies. Expected %v, got %v", len(initialData), len(result))
	}
	if !reflect.DeepEqual(movie, result[0]) {
		t.Fatalf("GetMovies returned wrong result. Expected %v, got %v", movie, result[0])
	}
}

func TestGetMovies_AfterAddMovie(t *testing.T) {
	initialData := []model.Movie{}
	movieService := app.NewMovieService(initialData)

	result := movieService.GetMovies()

	if len(result) != len(initialData) {
		t.Fatalf("GetMovies returned wrong no. of movies. Expected %v, got %v", len(initialData), len(result))
	}

	movie := model.MovieInput{
		Title:  "The Godfather",
		Genre:  "Crime",
		Year:   1972,
		Length: 143,
	}
	newId := movieService.AddMovie(movie)

	result = movieService.GetMovies()

	if len(result) != len(initialData)+1 {
		t.Fatalf("GetMovies returned wrong no. of movies. Expected %v, got %v", len(initialData), len(result))
	}
	if result[0].ID != newId {
		t.Fatalf("GetMovies returned wrong result. Expected %v, got %v", newId, result[0].ID)
	}
}

func TestGetLastModified(t *testing.T) {
	initialData := []model.Movie{}
	movieService := app.NewMovieService(initialData)

	movie := model.MovieInput{
		Title:  "The Godfather",
		Genre:  "Crime",
		Year:   1972,
		Length: 143,
	}
	oldLastModified := movieService.GetLastModified()
	movieService.AddMovie(movie)
	newLastModified := movieService.GetLastModified()

	if oldLastModified == newLastModified || !newLastModified.After(oldLastModified) {
		t.Fatalf("GetLastModified returned wrong result. Old %v, New %v", oldLastModified, newLastModified)
	}
}

func TestGetETag(t *testing.T) {
	initialData := []model.Movie{}
	movieService := app.NewMovieService(initialData)

	movie := model.MovieInput{
		Title:  "The Godfather",
		Genre:  "Crime",
		Year:   1972,
		Length: 143,
	}
	oldETag := movieService.GetETag()
	movieService.AddMovie(movie)
	newETag := movieService.GetETag()

	if oldETag == newETag {
		t.Fatalf("GetETag returned wrong result. Old %v, New %v", oldETag, newETag)
	}
}
