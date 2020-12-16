FROM golang:1.15 AS build
MAINTAINER John Anthony
WORKDIR /build
ENV CGO_ENABLED=0
COPY go.mod go.sum ./
RUN go mod download
COPY static static
COPY main.go ./
RUN CGO_ENABLED=0 go build -o a .

FROM gcr.io/distroless/static:nonroot
ENV GIN_MODE=release
ENTRYPOINT ["/usr/local/bin/a"]
COPY --from=build /build/a /usr/local/bin/a
