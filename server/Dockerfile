FROM golang:1.16
COPY . .
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go get go.mongodb.org/mongo-driver/mongo
RUN go get -u github.com/gorilla/mux
ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor
ENV DB_URI "mongodb://root:root@127.0.0.1:27017/?compressors=disabled&gssapiServiceName=mongodb"
ENV DB_NAME "test"
ENV DB_COLLECTION_NAME "todolist"

WORKDIR /go
EXPOSE 8080
CMD ["go","build"]