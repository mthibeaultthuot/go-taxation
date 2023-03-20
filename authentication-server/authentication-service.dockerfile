# base go image

FROM golang:1.19-alpine as builder


RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLE=0 go build -o authenticationApp ./cmd/api

RUN chmod +x /app/authenticationApp

FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/authenticationApp /app

CMD [ "/app/authenticationApp" ]

EXPOSE 9001