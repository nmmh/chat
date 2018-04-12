FROM golang:latest

RUN go get github.com/nmmh/chat/cmd/chat-client

EXPOSE 6000
WORKDIR /go/src/github.com/nmmh/chat/cmd/chat-client

RUN go build; go install

CMD ["chat-client -addr=:6000"]
#docker run -i -t -p 127.0.0.1:6000:6000/tcp chat # needed for windows?