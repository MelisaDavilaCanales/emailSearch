FROM node:18.8.0-alpine

RUN mkdir app

WORKDIR /app

RUN npm i -g http-server

COPY ./package*.json ./

RUN npm install

COPY . .

RUN npm run build

CMD [ "http-server", "dist"]