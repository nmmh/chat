FROM golang:latest

RUN go get github.com/nmmh/chat/cmd/chat

EXPOSE 6000
WORKDIR /go/src/github.com/nmmh/chat/cmd/chat

RUN go build; go install

CMD ["chat"]
#docker run -i -t -p 127.0.0.1:6000:6000/tcp chat # needed for windows?