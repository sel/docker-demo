# Docker Demo

A small application to demonstrate the use of [Docker](https://www.docker.com/).


## Prerequisites

- `docker`
- `go`


## Usage

- Run `go run main.go` to run the application natively on your host machine.

- Run `make run` to run the application inside a container on your host machine.
  By default the application is available from http://127.0.0.1:8000/

- Run `make stop` to stop the running container.

- Run `make build` to build a container image from the instructions in [`Dockerfile`](Dockerfile).

- Run `make push` to build and push the container image to a registry.
  >**Note**
  >Modify the `REPO` variable in [`Makefile`](Makefile) to that of your own container repository.


## Configuration

The port on which the server listens for requests is configured through the `LISTEN_PORT`
environment variable and defaults to `8000`.
Modify the `LISTEN_PORT` variable in [`Makefile`](Makefile) to change to a different value.


## References

- https://docs.docker.com/get-started/
- https://12factor.net/
