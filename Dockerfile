ARG GO_VERSION
FROM golang:${GO_VERSION} AS build
ARG VERSION

WORKDIR /build

COPY go.mod go.sum .
RUN go mod download

COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux GOFLAGS=-ldflags=-w go build -o hello \
      -ldflags="-s -X main.VERSION=${VERSION}" \
      main.go

FROM gcr.io/distroless/static-debian11 AS final
COPY --from=build /build/hello /usr/local/bin/hello
USER 65534:65534
ENTRYPOINT [ "/usr/local/bin/hello" ]
