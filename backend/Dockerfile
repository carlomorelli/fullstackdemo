# Build stage

FROM golang:1.10
RUN mkdir /build
ADD . /build/ 
WORKDIR /build 

RUN go get -u -v \
    github.com/gorilla/mux \
    github.com/mattn/go-sqlite3 \
    github.com/rs/cors \
    github.com/stretchr/testify

RUN go test -v -tags unit

RUN CGO_ENABLED=0 GOOS=linux go build -o backend -a -installsuffix cgo

# Execution stage

FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN mkdir /app
WORKDIR /app
COPY --from=0 /build/backend .

EXPOSE 9000

CMD [ "/app/backend" ]
