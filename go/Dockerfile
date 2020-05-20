FROM golang:1.14

RUN apt-get update && apt-get install -y unzip vim

ENV GO111MODULE=on

RUN go get -u google.golang.org/grpc && \
    go get -u github.com/golang/protobuf/protoc-gen-go

RUN wget https://github.com/protocolbuffers/protobuf/releases/download/v3.9.0-rc1/protoc-3.9.0-rc-1-linux-x86_64.zip && \
    unzip protoc-3.9.0-rc-1-linux-x86_64.zip -d protoc && \
    cd protoc && \
    mv bin/protoc /usr/bin

ENV PATH $PATH:$GOPATH/bin

WORKDIR /go/src/grpc-training-ruby-go

COPY . .
#RUN apk add --no-cache alpine-sdk git && go get -u github.com/oxequa/realize

EXPOSE 8083
CMD ["go", "run", "server.go"]
