FROM golang:1.15.5-buster AS builder
ARG VERSION=dev
COPY . /go/src/app
WORKDIR /go/src/app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main -ldflags=-X=main.version=${VERSION} main.go

FROM alpine:latest
COPY --from=builder /go/src/app/main /go/bin/main
ENV PATH="/go/bin:${PATH}"
EXPOSE 9010
CMD ["main"]