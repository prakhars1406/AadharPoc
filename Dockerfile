FROM golang:1.16.6-alpine3.14
COPY . /go/src/AadharPoc/
WORKDIR "/go/src/AadharPoc/app"
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' .

FROM alpine:latest
WORKDIR /root
COPY --from=0 "/go/src/AadharPoc/app" .
CMD ["./app"]
#testing
