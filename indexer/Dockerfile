FROM golang:1.23.1-alpine

RUN apk add --no-cache bash curl

# Descarga el script `wait-for` para esperar la disponibilidad del servicio DB
# ADD https://raw.githubusercontent.com/eficode/wait-for/master/wait-for /usr/local/bin/wait-for
# RUN chmod +x /usr/local/bin/wait-for

RUN mkdir app
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main .

# CMD ["wait-for", "db:5432", "--", "./main"]