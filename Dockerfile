# Initial stage
FROM golang:1.18 as modules

ADD go.mod go.sum /m/
RUN cd /m && go mod download


# Build stage
FROM golang:1.18-alpine as builder

COPY --from=modules /go/pkg /go/pkg

RUN mkdir -p /application
ADD . /application
WORKDIR /application
RUN go mod vendor

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
    go build -mod=vendor -o ./bin/main.out ./cmd/main.go

WORKDIR /application/bin
CMD ["./main.out"]