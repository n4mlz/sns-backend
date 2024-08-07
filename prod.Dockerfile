FROM golang:latest AS build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/app

FROM debian:12-slim AS prod

ARG FRONTEND_URL
ARG DB_USER
ARG DB_PASSWORD
ARG DB_HOST
ARG DB_PORT
ARG DB_NAME
ARG S3_ACCOUNT_ID
ARG S3_ACCESS_KEY_ID
ARG S3_ACCESS_KEY_SECRET
ARG S3_ENDPOINT
ARG S3_BUCKET_NAME
ARG S3_RESOURCE_URL

ENV FRONTEND_URL=${FRONTEND_URL}
ENV DB_USER=${DB_USER}
ENV DB_PASSWORD=${DB_PASSWORD}
ENV DB_HOST=${DB_HOST}
ENV DB_PORT=${DB_PORT}
ENV DB_NAME=${DB_NAME}
ENV S3_ACCOUNT_ID=${S3_ACCOUNT_ID}
ENV S3_ACCESS_KEY_ID=${S3_ACCESS_KEY_ID}
ENV S3_ACCESS_KEY_SECRET=${S3_ACCESS_KEY_SECRET}
ENV S3_ENDPOINT=${S3_ENDPOINT}
ENV S3_BUCKET_NAME=${S3_BUCKET_NAME}
ENV S3_RESOURCE_URL=${S3_RESOURCE_URL}

COPY --from=build /go/bin/app /go/bin/app
COPY serviceAccountKey.json .
RUN apt update -y && apt-get install -y ca-certificates
EXPOSE 8080

CMD ["/go/bin/app"]
