FROM golang:1.14 AS build
MAINTAINER John Anthony
WORKDIR /build
ENV CGO_ENABLED=0
RUN go get -u github.com/go-bindata/go-bindata/...
COPY go.mod go.sum ./
RUN go mod download
COPY static static
COPY main.go ./
RUN go generate ./...
RUN CGO_ENABLED=0 go build -o a .

FROM  gcr.io/distroless/static:nonroot
ENV GIN_MODE=release
ENTRYPOINT ["/usr/local/bin/a"]
COPY --from=build /build/a /usr/local/bin/a
