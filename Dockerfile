FROM golang:1.21.5-alpine AS builder

COPY . /github.com/FreylGit/chat-server/sourse/
WORKDIR /github.com/FreylGit/chat-server/sourse/

RUN go mod download
RUN go build -o ./bin/chat-server ./cmd/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/FreylGit/chat-server/sourse/bin/chat-server .

CMD ["./chat-server"]