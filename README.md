# Syndica-THT

#### Objective
You are requested to build a Go web service which helps in fetching the list of movies from the
server and serving the content on a digital signage of cinemas.

#### Requirements
You are expected to work on the following areas:
1. Designing API endpoints.
	a. Build an API that provides functionality to add / get movies.
	b. You would need to provide endpoints through which the client could add and get movie(s).
2. Create the data (in-memory or mock api).
	a. Add 100 movies at most hard-coded saved in a json file / mock api.
	b. Read data when the server is first started.
3. Write handler to return all movies.
a. GET (/movies) - Fetch list of all movies.
4. Write handler to add a new movie.
a. POST (/movies) - Add a new movie to the list.
5. Write handler to return a specific movie.
a. GET (/movies/:id) - Fetch a movie by its entity id.

Take the following points into consideration:
- Data should be cached.
- As this digital signage screen refreshes every 15 seconds, we need to ensure that we
only fetch the data in case there is a change on the server.
- And by default every hour it will fetch the latest data from the server regardless.
- There should be an independent boolean flag (force) which should fetch the latest data if set true.
- The requests / responses should be in JSON-RPC format.

\* Please provide a repository URL (e.g Github) to send it after completion.

#### Tech Design
We have used in-memory data structures to implement this assignment.
* `src/app/app.go` defines the JSON-RPC implementation.
* `src/app/movie_controller.go` implements the RPC controllers.
* `src/app/movie_service.go` contains service layer / business logic for Add/Get Movie(s).
* `cmd/server.go` contains the code to start the server.

##### Dependencies:
1. Go v1.12 or above (Installation guide: https://go.dev/doc/install)

##### Running the code
- Run server after moving the root directory of repository
	>`go run cmd/server.go`
