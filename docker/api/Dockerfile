FROM golang

LABEL maintainer = "tomomo@eclairpark@gmail.com"

WORKDIR /go/src

COPY ./src .

RUN go get -u github.com/cosmtrek/air && \
  go build -o /go/bin/air github.com/cosmtrek/air && \
  go get -u github.com/gin-gonic/gin

CMD ["air", "-c", ".air.toml"]
