FROM alpine:latest

ADD icppp /home/app/
COPY conf /home/app/conf

WORKDIR /home/app
ENTRYPOINT ["./icppp", "web"]

EXPOSE 9315