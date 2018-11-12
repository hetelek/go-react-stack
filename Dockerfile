FROM golang:latest 

ADD ./main.go /app/main.go
ADD ./backend /app/backend
ADD ./frontend/build /app/frontend/build

WORKDIR /app/
RUN go get -d .
RUN go build -o start .

ENV PORT 8000
CMD /app/start

EXPOSE 8000
