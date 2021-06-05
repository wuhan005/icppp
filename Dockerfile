FROM golang:alpine as builder

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk --no-cache add git

WORKDIR /go/src/github.com/go/icppp/

COPY . .

ENV GOPROXY https://goproxy.cn

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o icppp .

FROM alpine:latest as prod

RUN apk --no-cache add ca-certificates

WORKDIR /home/app/

COPY --from=0 /go/src/github.com/go/icppp/icppp .

COPY config /home/app/config

COPY template /home/app/template

COPY static /home/app/static

EXPOSE 9315

CMD ["./icppp"]
