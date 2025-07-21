FROM golang:alpine AS build
WORKDIR /build

COPY . .
RUN go build -o garg

FROM alpine
WORKDIR /app

COPY --from=build /build .
ENTRYPOINT ["./garg"]
