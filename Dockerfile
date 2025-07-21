FROM golang:alpine AS build
WORKDIR /build

COPY . .
RUN go build -o garg

FROM alpine
WORKDIR /app

RUN apk update && \
  apk add git

COPY --from=build /build .
ENTRYPOINT ["./garg"]
