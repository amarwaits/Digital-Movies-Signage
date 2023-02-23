package test

import (
	"bytes"
	JSON "encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/amarwaits/Syndica-THT/src/app"
	"github.com/amarwaits/Syndica-THT/src/model"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
)

func TestAddMovieController(t *testing.T) {
	movie := model.Movie{
		ID:     1,
		Title:  "The Godfather",
		Genre:  "Crime",
		Year:   1972,
		Length: 143,
	}
	initialData := []model.Movie{movie}
	movieService := app.NewMovieService(initialData)
	movieController := app.NewMovieController(movieService)

	// Set up the RPC server
	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json")
	s.RegisterService(movieController, "")

	// Create a test request
	req, err := http.NewRequest("POST", "/rpc", bytes.NewBufferString(`{
    "jsonrpc": "2.0",
    "method": "MovieController.AddMovie",
    "params": [{"Title": "The Godfather", "Genre": "Crime", "Year": 1972, "Length": 143}],
    "id": 1
  }`))
	if err != nil {
		t.Fatal(err)
	}

	// Create a test response writer
	w := httptest.NewRecorder()

	// Serve the RPC request
	s.ServeHTTP(w, req)

	// Check the response status code
	if w.Code != http.StatusOK {
		t.Fatalf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	// Check the response body
	var resp struct {
		ID     int    `json:"id"`
		Result int    `json:"result"`
		Error  *error `json:"error"`
	}
	if err := JSON.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatal(err)
	}
	if resp.Error != nil {
		t.Fatalf("Got an error: %v", resp.Error)
	}
	if resp.Result != 2 {
		t.Fatalf("Expected movie ID 2, but got %d", resp.Result)
	}
}

func TestGetMovieController(t *testing.T) {
	movie := model.Movie{
		ID:     1,
		Title:  "The Godfather",
		Genre:  "Crime",
		Year:   1972,
		Length: 143,
	}
	initialData := []model.Movie{movie}
	movieService := app.NewMovieService(initialData)
	movieController := app.NewMovieController(movieService)

	// Set up the RPC server
	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json")
	s.RegisterService(movieController, "")

	// Create a test request
	req, err := http.NewRequest("POST", "/rpc", bytes.NewBufferString(`{
    "jsonrpc": "2.0",
    "method": "MovieController.GetMovie",
    "params": [{"ID": 1}],
    "id": 1
  }`))
	if err != nil {
		t.Fatal(err)
	}

	// Create a test response writer
	w := httptest.NewRecorder()

	// Serve the RPC request
	s.ServeHTTP(w, req)

	// Check the response status code
	if w.Code != http.StatusOK {
		t.Fatalf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	// Check the response body
	var resp struct {
		ID     int         `json:"id"`
		Result model.Movie `json:"result"`
		Error  *string     `json:"error"`
	}
	if err := JSON.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatal(err)
	}
	if resp.Error != nil {
		t.Fatalf("Got an error: %v", *resp.Error)
	}
	if resp.Result.ID != 1 {
		t.Fatalf("Expected movie ID 1, but got %d", resp.Result.ID)
	}
}

func TestGetMovieController_NotFound(t *testing.T) {
	initialData := []model.Movie{}
	movieService := app.NewMovieService(initialData)
	movieController := app.NewMovieController(movieService)

	// Set up the RPC server
	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json")
	s.RegisterService(movieController, "")

	// Create a test request
	req, err := http.NewRequest("POST", "/rpc", bytes.NewBufferString(`{
    "jsonrpc": "2.0",
    "method": "MovieController.GetMovie",
    "params": [{"ID": 1}],
    "id": 1
  }`))
	if err != nil {
		t.Fatal(err)
	}

	// Create a test response writer
	w := httptest.NewRecorder()

	// Serve the RPC request
	s.ServeHTTP(w, req)

	// Check the response status code
	if w.Code != http.StatusOK {
		t.Fatalf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	// Check the response body
	var resp struct {
		ID     int         `json:"id"`
		Result model.Movie `json:"result"`
		Error  *string     `json:"error"`
	}
	if err := JSON.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatal(err)
	}
	if resp.Error == nil {
		t.Fatalf("Didn't get an error")
	}
}

func TestGetMoviesController(t *testing.T) {
	movie := model.Movie{
		ID:     1,
		Title:  "The Godfather",
		Genre:  "Crime",
		Year:   1972,
		Length: 143,
	}
	initialData := []model.Movie{movie}
	movieService := app.NewMovieService(initialData)
	movieController := app.NewMovieController(movieService)

	// Set up the RPC server
	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json")
	s.RegisterService(movieController, "")

	// Create a test request
	req, err := http.NewRequest("POST", "/rpc", bytes.NewBufferString(`{
    "jsonrpc": "2.0",
    "method": "MovieController.GetMovies",
    "params": [{"force": true}],
    "id": 1
  }`))
	if err != nil {
		t.Fatal(err)
	}

	// Create a test response writer
	w := httptest.NewRecorder()

	// Serve the RPC request
	s.ServeHTTP(w, req)

	// Check the response status code
	if w.Code != http.StatusOK {
		t.Fatalf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	// Check the response body
	var resp struct {
		ID     int           `json:"id"`
		Result []model.Movie `json:"result"`
		Error  *string       `json:"error"`
	}
	if err := JSON.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatal(err)
	}
	if resp.Error != nil {
		t.Fatalf("Got an error: %v", *resp.Error)
	}
	if resp.Result[0].ID != 1 {
		t.Fatalf("Expected movie ID 1, but got %d", resp.Result[0].ID)
	}
}

// func TestAddMovieController(t *testing.T) {
// 	// Set up the RPC server
// 	s := rpc.NewServer()
// 	s.RegisterCodec(json.NewCodec(), "application/json")
// 	s.RegisterService(new(MoviesService), "")

// 	// Create a test request
// 	req, err := http.NewRequest("POST", "/rpc", bytes.NewBufferString(`{
//     "jsonrpc": "2.0",
//     "method": "MoviesService.AddMovie",
//     "params": [{
//       "title": "The Matrix",
//       "release_year": 1999,
//       "genre": "Sci-Fi"
//     }],
//     "id": 1
//   }`))
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// Create a test response writer
// 	w := httptest.NewRecorder()

// 	// Serve the RPC request
// 	s.ServeHTTP(w, req)

// 	// Check the response status code
// 	if w.Code != http.StatusOK {
// 		t.Fatalf("Expected status code %d, but got %d", http.StatusOK, w.Code)
// 	}

// 	// Check the response body
// 	var resp struct {
// 		ID     int            `json:"id"`
// 		Result int            `json:"result"`
// 		Error  *jsonrpc.Error `json:"error"`
// 	}
// 	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
// 		t.Fatal(err)
// 	}
// 	if resp.Error != nil {
// 		t.Fatalf("Got an error: %v", resp.Error)
// 	}
// 	if resp.Result != 1 {
// 		t.Fatalf("Expected movie ID 1, but got %d", resp.Result)
// 	}
// }
