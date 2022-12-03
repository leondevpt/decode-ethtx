# syntax=docker/dockerfile:1.2

FROM golang:1.18.3-alpine3.16 as binary-builder
ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn,direct
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk update && apk upgrade && apk --update add git make ca-certificates  build-base libc6-compat && \
    mkdir  -p /app/build
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN make build

# Stage 2: # Run container
FROM alpine:3.16
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk --no-cache add  tzdata bash vim ca-certificates && \
    update-ca-certificates && \
    ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone

# Create a group and user
RUN addgroup appuser --gid 1000 && \
    adduser appuser --ingroup appuser --uid 1000 --home /home/appuser --disabled-password

ENV APP_PORT=3000
WORKDIR /app
COPY --from=binary-builder --chown=appuser:appuser /app/build/server .
USER appuser
EXPOSE $APP_PORT

ENTRYPOINT ["./server"]