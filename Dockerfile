FROM registry.access.redhat.com/ubi8/ubi

#Install golang
RUN dnf install tar wget go-toolset -y

ENV PATH="/usr/local/go/bin:$PATH"  GOLANG_VERSION=1.9 GOPATH=/go

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go mod init go-restful

RUN go mod tidy

COPY *.go ./

RUN go build go-restful

EXPOSE 8080

CMD [ "/app/go-restful" ]