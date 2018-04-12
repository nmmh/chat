FROM golang:latest

RUN go get github.com/nmmh/chat/cmd/chat-client

#EXPOSE 6000
WORKDIR /go/src/github.com/nmmh/chat/cmd/chat-client

RUN go build; go install

WORKDIR /go/bin/

#CMD [""] #chat-client -addr=:6000
