FROM golang:latest as builder

WORKDIR $GOPATH/src/github.com/snowspice/go-gin-gorm-router

ADD . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main -a -installsuffix cgo .

RUN rm -rf vendor

FROM alpine:latest

RUN apk --no-cache add ca-certificates


RUN mkdir /app
WORKDIR /app

COPY --from=builder /$GOPATH/src/github.com/snowspice/go-gin-gorm-router .
COPY --from=builder /$GOROOT/lib/time/zoneinfo.zip /usr/local/go/lib/time/zoneinfo.zip

EXPOSE 9090

CMD ["./main"]