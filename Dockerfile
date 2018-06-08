FROM golang:1.10.2-alpine3.7 as builder

WORKDIR /go/src/github.com/vothanhkiet/noop/
COPY . .

RUN apk --no-cache add git
RUN go get -u github.com/golang/dep/...
RUN dep ensure
RUN go build -o app .

FROM alpine:3.7
WORKDIR /root/
COPY --from=builder /go/src/github.com/vothanhkiet/noop/app .
RUN apk --no-cache add curl
CMD ["./app"]
