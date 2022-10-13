FROM golang:1.12-alpine as builder

WORKDIR /app

COPY main.go /app

RUN go build -o main .

FROM alpine:3.16.2

WORKDIR /app

COPY --from=builder /app/main /app/

ENV PORT=8080

EXPOSE 8080

ENTRYPOINT [ "/app/main" ]
