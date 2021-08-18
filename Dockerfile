FROM golang:1.16-alpine

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go mod tidy

COPY *.go ./

RUN go build go-restful

EXPOSE 8080

CMD [ "/app/go-restful" ]