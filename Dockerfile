FROM golang:1.22 AS builder

WORKDIR /go/src

COPY .. .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server cmd/server/main.go

FROM alpine:latest as prod

RUN apk --no-cache add ca-certificates
RUN apk add tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && apk del tzdata


WORKDIR /root/

COPY --from=builder /go/src/server .

EXPOSE 8080

ENTRYPOINT [ "./server" ]