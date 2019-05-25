FROM golang:1.11.10

WORKDIR /projects/starter

ENV GO111MODULE=on

COPY go.mod .
COPY go.sum .

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 8080

CMD ["starter"]