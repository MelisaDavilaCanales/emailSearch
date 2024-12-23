FROM golang:1.23.1-alpine

RUN apk add --no-cache bash curl

RUN mkdir /app
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Cambiar el comando de arranque
ENTRYPOINT ["go", "run", "."]
CMD ["${DATA_PATH}"]
