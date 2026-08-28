FROM golang:1.25 AS build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/app

FROM debian:12.8-slim AS prod

COPY --from=build /go/bin/app /go/bin/app
RUN apt update -y && apt-get install -y --no-install-recommends ca-certificates curl \
	&& useradd --system --no-create-home --shell /usr/sbin/nologin app \
	&& rm -rf /var/lib/apt/lists/* \
	&& chown app:app /go/bin/app
EXPOSE 8080

USER app
CMD ["/go/bin/app"]
