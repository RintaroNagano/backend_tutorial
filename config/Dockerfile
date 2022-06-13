FROM golang:1.18.3

ENV DOCKERIZE_VERSION v0.6.1
RUN apt-get update && apt-get install -y wget \
 && wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
 && tar -C /usr/local/bin -xzvf dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
 && rm dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz

run mkdir /app
workdir /app

run go env -w GO111MODULE=on
run go mod init sample
run go get github.com/gin-gonic/gin
run go get github.com/go-sql-driver/mysql
run go get github.com/jinzhu/gorm
run go get crypto/sha256
run go get encoding/hex
run go get github.com/dgrijalva/jwt-go

run go mod tidy