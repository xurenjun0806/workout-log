# ビルド
FROM golang:1.27.4-alpine3.24 AS builder

RUN apk update && apk upgrade && \
    apk --update add git make bash build-base

WORKDIR /app

COPY . .

RUN make build

# 配布
FROM alpine:3.24

RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata && \
    mkdir /app 

WORKDIR /app 

EXPOSE 9090

COPY --from=builder /app/bin/workout-log /app/

CMD ["/app/workout-log"]