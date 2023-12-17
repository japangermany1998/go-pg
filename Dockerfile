# build Goland binary to AMD architect to deploy
FROM golang:1.19-alpine AS build-env

ENV WDIR app

WORKDIR /$WDIR

RUN apk add upx

RUN apk add git

COPY go.mod .

COPY go.sum .

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags '-w'

RUN upx -9 /$WDIR/go-pg

FROM alpine:latest

RUN mkdir -p /app/bin

WORKDIR /app/bin

COPY --from=build-env /app/go-pg /app/bin/

ENTRYPOINT ["./go-pg"]

EXPOSE 8080
