# base go image

FROM golang:1.19-alpine as builder


RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLE=0 go build -o verificationApp ./cmd/api

RUN chmod +x /app/verificationApp

FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/verificationApp /app

CMD [ "/app/verificationApp" ]

EXPOSE 8081