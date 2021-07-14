FROM golang:1.16.6-buster

LABEL maintainer = "tomomo@eclairpark@gmail.com"

WORKDIR /go/src

COPY ./src .

RUN go get -u github.com/cosmtrek/air && \
  go build -o /go/bin/air github.com/cosmtrek/air && \
  go get -u github.com/gin-gonic/gin && \
  go get -u github.com/go-sql-driver/mysql

CMD ["air", "-c", ".air.toml"]
