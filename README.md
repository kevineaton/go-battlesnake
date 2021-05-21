# Battlesnake Implementation

This repo is my exploration of the [https://play.battlesnake.com/](Battlesnake) service in Go. It is a simple API designed to provide responses to the Battlesnake
HTTP calls for games. The logic is simple, and I may or may not commit my current algorithms.

As the API is pretty straight-forward, this is also a place to experiment with extras, such as persistnce and analysis (TODO).

## Technologies

- [github.com/go-chi/chi](Chi) - router
- [github.com/stretchr/testify](Testify) - test helpers
- [github.com/mitchellh/mapstructure](Mapstructure) - Data management

### Vendoring

I vendor my libraries for ease of use and development.

## Configuration

Currently, the following fields are configurable from the environment:

- BS_API_PORT - the port for the server to listen on, default to 7000

## Running

If you have the code but not the binary, first build the code:

`$ go build .`

Next, run the binary:

`$ BS_API_PORT=8888 ./go-battlesnake`

This will start the server. To test if it is up, you can issue a command against the status or health endpoints:

```bash
$ curl http://localhost:8888/
{"status":"up"}
```

### Developing

The majority of the code is in the `api` subdirectory. The routes are set up in the `*_routes.go` files, while the logic is in the appropriately named files. Always write tests, although 100% coverage is not likely. Test the main areas to ensure new code works and doesn't introduce regressions.

#### Testing

Several test helpers are found in `http.go` to test the Chi endpoints. To run the tests, make sure you point at the appropriate directory:

`$ go test -v ./api`

For coverage:

`go test -coverprofile=coverage.out -v ./api && go tool cover -html=coverage.out -o coverage.html`

### Contributing

Contributions welcome but may not be merged. It's best to reach out first. This is more of a fun side-project playground than something designed to be battle-tested or serious. Feel free to send something in, or fork and build on your own.
