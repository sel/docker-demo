# Docker Demo

A small application to demonstrate the use of [Docker](https://www.docker.com/).


## Prerequisites

- `docker`
- `go`


## Usage

- `go run main.go` runs the application natively on your host machine.
  By default the application is available from http://127.0.0.1:8000/

- `make run` runs the application inside a container on your host machine.
  By default the application is available from http://127.0.0.1:8000/

- `make stop` stops the running container.

- `make build` builds a container image from the instructions in [`Dockerfile`](Dockerfile).

- `make push` builds and pushes the container image to a registry.
  >**Note**
  >Modify the `REPO` variable in [`Makefile`](Makefile) to that of your own container repository.


## Configuration

The port on which the server listens for requests is configured through the `LISTEN_PORT`
environment variable and defaults to `8000`.
Modify the `LISTEN_PORT` variable in [`Makefile`](Makefile) to change to a different value.


## References

- https://docs.docker.com/get-started/
- https://12factor.net/
