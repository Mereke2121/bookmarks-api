FROM golang:1.17

RUN go version
ENV GOPATH=/

COPY ./ ./

# install psql
RUN apt-get update
RUN apt-get -y install postgresql-client

RUN chmod +x wait-for-postgres.sh


RUN go mod download
RUN go build -o bookmarks-api ./cmd/app/main.go

CMD ["./bookmarks-api"]