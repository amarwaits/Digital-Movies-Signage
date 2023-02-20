package app

import (
	"fmt"
	"sync"
	"time"

	"github.com/amarwaits/Syndica-THT/src/model"
)

type MovieService struct {
	Movies       []model.Movie
	movies       map[int]*model.Movie
	lastModified time.Time
	mutex        sync.Mutex
}

func NewMovieService(initialData []model.Movie) *MovieService {
	return &MovieService{
		Movies:       initialData,
		lastModified: time.Now().UTC(),
	}
}

func (s *MovieService) AddMovie(movie model.Movie) int {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	id := len(s.Movies) + 1
	s.movies[id] = &movie
	s.Movies = append(s.Movies, movie)
	s.lastModified = time.Now().UTC()
	return id
}

func (s *MovieService) GetMovie(id int) (model.Movie, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	movie, ok := s.movies[id]
	if !ok {
		return model.Movie{}, fmt.Errorf("Movie not found with ID %d", id)
	}
	return *movie, nil
}

func (s *MovieService) GetMovies() []model.Movie {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.Movies
}

func (s *MovieService) GetLastModified() time.Time {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.lastModified
}

func (s *MovieService) GetETag() string {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return fmt.Sprintf("%d:%d", len(s.Movies), s.lastModified.UnixNano())
}
