FROM golang:1.14

# RUN go version
ENV GOPATH=/

COPY ./ ./

# install psql
RUN apt-get update
RUN apt-get -y install postgresql-client

# make wait-for-postgres.sh executable
RUN chmod +x wait-for-postgres.sh

# build go app
RUN go mod download
RUN go build -o main ./cmd/crud-api-server/main.go

# WORKDIR /cmd/crud-api-server/
CMD ["./main"]
