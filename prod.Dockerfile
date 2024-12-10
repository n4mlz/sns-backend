FROM golang:1.23 AS build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/app

FROM debian:12.8-slim AS prod

COPY --from=build /go/bin/app /go/bin/app
COPY serviceAccountKey.json .
RUN apt update -y && apt-get install -y ca-certificates
EXPOSE 8080

CMD ["/go/bin/app"]
