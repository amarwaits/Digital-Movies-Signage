package app

import (
	"log"
	"net/http"

	"github.com/amarwaits/Syndica-THT/src/config"
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
)

type App struct {
	Config          *config.Config
	MovieService    *MovieService
	MovieController *MovieController
}

func NewApp() *App {
	// Initialize configuration
	config := config.NewConfig()

	// Load initial movie data from JSON file
	initialData, err := config.LoadMoviesFromFile()
	if err != nil {
		log.Fatalf("Failed to load initial movie data: %v", err)
	}

	// Initialize movie service with the initial data
	movieService := NewMovieService(initialData)

	// Initialize movie controller with the movie service
	movieController := NewMovieController(movieService)

	return &App{
		Config:          config,
		MovieService:    movieService,
		MovieController: movieController,
	}
}

func (a *App) Start() error {
	rpcServer := rpc.NewServer()
	rpcServer.RegisterCodec(json.NewCodec(), "application/json")
	rpcServer.RegisterService(a.MovieController, "")

	r := mux.NewRouter()
	r.Handle("/movies", rpcServer)
	return http.ListenAndServe("localhost:8080", r)
}
