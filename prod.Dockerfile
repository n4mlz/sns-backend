FROM golang:1.27 AS build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/app

FROM debian:13.6-slim AS prod

COPY --from=build /go/bin/app /go/bin/app
RUN apt update -y && apt-get install -y --no-install-recommends ca-certificates curl \
    && rm -rf /var/lib/apt/lists/*
EXPOSE 8080

CMD ["/go/bin/app"]
