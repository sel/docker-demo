FROM golang:1.19 AS build
WORKDIR /build

COPY go.mod go.sum .
RUN go mod download

COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux GOFLAGS=-ldflags=-w go build -o hello \
      -ldflags="-s" \
      main.go

FROM gcr.io/distroless/static-debian11 AS final
COPY --from=build /build/hello /usr/local/bin/hello
USER 65534:65534
ENTRYPOINT [ "/usr/local/bin/hello" ]
