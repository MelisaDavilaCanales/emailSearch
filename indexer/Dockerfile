FROM golang:1.23.1-alpine

RUN apk add --no-cache bash curl

RUN mkdir /app
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

ARG DATA_PATH

ENV ZINC_SEARCH_API_URL=${ZINC_SEARCH_API_URL} \
    DB_USER=${DB_USER} \
    DB_PASSWORD=${DB_PASSWORD} \
    DATA_PATH=${DATA_PATH} \
    APP_MODE=${APP_MODE}

COPY . .

# Cambiar el comando de arranque
ENTRYPOINT ["sh", "-c", "go run . $DATA_PATH"]