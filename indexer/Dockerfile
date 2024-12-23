FROM golang:1.23.1-alpine

RUN apk add --no-cache bash curl

RUN mkdir app
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main .

# Define un argumento para la ruta del directorio con la data
ARG DATA_DIR

# Usa el argumento en el comando de inicio
CMD ["sh", "-c", "./main ${DATA_DIR}"]

