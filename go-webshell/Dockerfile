FROM golang:1.23-bullseye AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY *.go ./

COPY . .
RUN go build -o go-webshell -ldflags='-w -s' .

##
## Deploy
##
FROM debian:testing-slim

WORKDIR /

COPY --from=build /app/go-webshell .

EXPOSE 8080

CMD ["./go-webshell"]