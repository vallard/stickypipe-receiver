FROM golang:onbuild
EXPOSE 8888
WORKDIR /go/src/app
COPY . /go/src/app


