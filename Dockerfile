FROM alpine:latest

ADD icppp /home/app/
COPY config /home/app/config
COPY template /home/app/template
COPY static /home/app/static

WORKDIR /home/app
ENTRYPOINT ["./icppp"]

EXPOSE 9315