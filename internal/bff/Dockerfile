FROM golang:1.22-alpine

WORKDIR $GOPATH/src/doggo-track

COPY . .
RUN go mod download

WORKDIR $GOPATH/src/doggo-track/cmd/bff

RUN go build -o $GOPATH/bin/svc

WORKDIR $GOPATH/bin

ENTRYPOINT ["/go/bin/svc", "--cfgPath=/go/src/doggo-track/internal/bff/.env.dev"]