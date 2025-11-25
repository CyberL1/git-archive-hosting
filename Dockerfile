FROM node:alpine AS build-web
WORKDIR /build

COPY frontend .

RUN npm i && \
    npm run build

FROM golang:alpine AS build
WORKDIR /build

COPY . .
COPY --from=build-web /build frontend

RUN go build -o garg

FROM alpine
WORKDIR /app

RUN apk add git --no-cache

COPY --from=build /build .
ENTRYPOINT ["./garg"]
