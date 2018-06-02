FROM golang:1.10

RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 

RUN go get -u -v \
    github.com/gorilla/mux \
    github.com/mattn/go-sqlite3 \
    github.com/rs/cors \
    github.com/stretchr/testify

RUN go build -o api
RUN go test

EXPOSE 9000

CMD ["/app/api"]
