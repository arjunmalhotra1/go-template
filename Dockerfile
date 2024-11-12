FROM golang:1.23-alpine

WORKDIR /app

COPY ./go.mod .
RUN go mod download

COPY . .

WORKDIR /app/cmd/server

EXPOSE 8080

CMD [ "go","run","main.go" ]