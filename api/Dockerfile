FROM golang:1.23.1-alpine

RUN mkdir app

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o api .

CMD ["./api"]