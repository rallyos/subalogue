FROM golang:1.14

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

RUN go get github.com/cespare/reflex

EXPOSE 8000
ENV SUBALOGUE_ENV=development

CMD reflex -r '\.go$' -s -- sh -c 'go build . && ./subalogue'
