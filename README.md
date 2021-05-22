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
- BS_SNAKE_COLOR - the desired color of the snake, such as `#FF00AB`. If not provided, it will be randomized at startup.
- BS_SNAKE_HEAD - the desired snake head. If not provided, it will be randomized at startup.
- BS_SNAKE_TAIL - the desired snake tail. If not provided, it will be randomized at startup.
- BS_AUTHOR - your name, defaults to `Someone Online`
- BS_VERSION - the version of your API, such as a tag or semver number, defaults to v0.0.1

## Running

### Docker

A Docker image can be pulled with:

`docker pull kevineaton/go-battlesnake`. Running it is as simple as passing in the configuration in the environment:

`docker run -e BS_API_PORT=8987 -e BS_SNAKE_COLOR="#FFFFFF" kevineaton/go-battlesnake`

### From Scratch

If you have the code but not the binary, first build the code:

`$ go build .`

Next, run the binary:

`$ BS_API_PORT=8888 ./go-battlesnake`

This will start the server. To test if it is up, you can issue a command against the status or health endpoints:

```bash
$ curl http://localhost:8888/
{"status":"up"}
```

*NOTE*: Please ensure that, when you go live, you run this over an SSL connection. I personally prefer to place this behind nginx as a reverse proxy with a free Let's Encrypt certificate for TLS.

### Developing

The majority of the code is in the `api` subdirectory. The routes are set up in the `*_routes.go` files, while the logic is in the appropriately named files. Always write tests, although 100% coverage is not likely. Test the main areas to ensure new code works and doesn't introduce regressions.

#### Testing

Several test helpers are found in `http.go` to test the Chi endpoints. To run the tests, make sure you point at the appropriate directory:

`$ go test -v ./api`

For coverage:

`go test -coverprofile=coverage.out -v ./api && go tool cover -html=coverage.out -o coverage.html`

### Contributing

Contributions welcome but may not be merged. It's best to reach out first. This is more of a fun side-project playground than something designed to be battle-tested or serious. Feel free to send something in, or fork and build on your own.
