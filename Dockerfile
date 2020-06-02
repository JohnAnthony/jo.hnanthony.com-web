FROM golang:1.14 AS build
MAINTAINER John Anthony
WORKDIR /build
ENV CGO_ENABLED=0
COPY main.go static.go go.mod go.sum ./
RUN CGO_ENABLED=0 go build -o a .

FROM  gcr.io/distroless/static:nonroot
ENV GIN_MODE=release
ENTRYPOINT ["/usr/local/bin/a"]
COPY --from=build /build/a /usr/local/bin/a
