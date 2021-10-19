FROM golang:alpine

ENV GO111MODULE=off \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY . .

RUN go build -o app .

WORKDIR /dist

RUN cp /build/app .

EXPOSE 9000

CMD ["/dist/app"]
