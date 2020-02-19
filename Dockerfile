FROM golang:latest


LABEL maintainer="mrnaghibi <naghibi.mohammadtaghi@gmail.com>"

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download 

COPY . .

ENV PORT :8000
ENV BASEURL http://discount:8000

RUN go build -o main .


CMD ["./main"]