# base image for building the binary
FROM golang:1.11-alpine AS base
COPY . /go/src/github.com/davidkroell/tictacgo

WORKDIR /go/src/github.com/davidkroell/tictacgo
RUN apk add git gcc && GO111MODULE=off go get -u

# run tests before build
RUN go test ./... -race

# binary output path: /go/bin/tictacgo
RUN CGO_ENABLED=0 go build -a -tags netgo -ldflags '-w' -o /go/bin/tictacgo *.go


# add the binary to an empty image
FROM scratch
COPY --from=base /go/bin/tictacgo /tictacgo
ENTRYPOINT ["/tictacgo"]
