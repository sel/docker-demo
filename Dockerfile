ARG GO_VERSION
FROM golang:${GO_VERSION} AS build
ARG VERSION

WORKDIR /usr/local/src/docker-demo

COPY go.mod go.sum .
RUN go mod download && go mod verify

COPY main.go .
RUN CGO_ENABLED=0 \
      go build -v -o /usr/local/bin/docker-demo \
      -ldflags="-s -w -X main.VERSION=${VERSION}" \
      main.go

FROM scratch AS final
COPY --from=build /usr/local/bin/docker-demo /usr/local/bin/docker-demo
USER 65534:65534
ENTRYPOINT [ "/usr/local/bin/docker-demo" ]
