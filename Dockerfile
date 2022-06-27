FROM golang:1.18

WORKDIR /app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

RUN go install github.com/cespare/reflex@latest

EXPOSE 3000
