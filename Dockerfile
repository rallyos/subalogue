FROM golang:1.14

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 8000
ENV SUBALOGUE_ENV=development

CMD ["go", "run", "main.go"]
