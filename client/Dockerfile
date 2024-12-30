FROM node:18.8.0-alpine

RUN mkdir app

WORKDIR /app

RUN npm i -g http-server

ARG VITE_API_URL

ENV VITE_API_URL=${VITE_API_URL} 

COPY ./package*.json ./

RUN npm install

COPY . .

RUN npm run build

CMD [ "http-server", "dist"]