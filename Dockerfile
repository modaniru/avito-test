FROM golang:latest as builder

ENV GOPATH=/

WORKDIR /app

COPY cmd cmd
COPY internal internal
COPY docs docs
COPY go.mod .
COPY go.sum .

RUN go get ./...
RUN CGO_ENABLED=0 go build -o main cmd/main.go

FROM centos:latest

COPY --from=builder /app/main .

CMD [ "./main" ]


