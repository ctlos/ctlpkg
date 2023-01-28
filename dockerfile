FROM golang:1.19.1 AS build
WORKDIR /
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o go-sqlc ./main.go
FROM alpine:3.16.2
COPY --from=build /go-sqlc .
ENTRYPOINT ["/go-sqlc"]
CMD ["--help"]