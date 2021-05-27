FROM golang:1.15 AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /go/src/gin-sample
COPY . .
RUN go build -ldflags="-s -w" -o /app/gin-sample ./cmd/gin-sample


FROM alpine

RUN apk add --no-cache tzdata
ENV TZ Asia/Shanghai
#ENV GIN_MODE=release

WORKDIR /app
COPY --from=builder /app/gin-sample /app/gin-sample

ENTRYPOINT /app/gin-sample