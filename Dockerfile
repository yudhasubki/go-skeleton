FROM golang:1.12 AS builder

WORKDIR $GOPATH/src/github.com/yudhasubki/go-skeleton
COPY . .
RUN go get ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /bin/app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /bin/app .
CMD ["./app"]
