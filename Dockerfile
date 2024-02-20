FROM golang:1.22 AS builder

ARG APP_RELATIVE_PATH

WORKDIR /app

COPY .. /app

RUN rm -rf /app/bin/
RUN go mod tidy && go build -ldflags="-s -w" -o ./bin/server  cmd/server/main.go

RUN mv config /app/bin/


FROM alpine:3.16


RUN apk add tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && apk del tzdata


ARG APP_ENV
ENV APP_ENV=${APP_ENV}

WORKDIR /app
COPY --from=builder /app/bin/server ./server

EXPOSE 8080
ENTRYPOINT [ "./server" ]

#docker build -t  1.1.1.1:5000/demo-api:v1 --build-arg APP_CONF=config/prod.yml --build-arg  APP_RELATIVE_PATH=./cmd/server/...  .
#docker run -it --rm --entrypoint=ash 1.1.1.1:5000/demo-api:v1