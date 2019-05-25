FROM golang:1.11.10

WORKDIR /projects/starter

ENV GO111MODULE=on

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

EXPOSE 8080

CMD ["/projects/starter"]