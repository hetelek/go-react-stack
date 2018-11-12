FROM golang:latest 

ADD . /app/
WORKDIR /app/
RUN go get -d .
RUN go build -o start .

ENV PORT 8000
CMD /app/start

EXPOSE 8000
