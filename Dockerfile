FROM alpine:latest

ADD icppp /home/app/
COPY config /home/app/config
WORKDIR /home/app
ENTRYPOINT ["./icppp"]

EXPOSE 9315