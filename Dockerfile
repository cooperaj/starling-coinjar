# Stage1
FROM golang:1.11-alpine AS build

RUN apk --no-cache add git
WORKDIR /go/src/github.com/cooperaj/starling-coinjar
ADD . .
RUN go get ./... && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o coinjar cmd/coinjar/main.go


# Stage2
FROM alpine

RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=build /go/src/github.com/cooperaj/starling-coinjar/coinjar .

EXPOSE 5000
HEALTHCHECK --interval=10s --timeout=3s CMD wget -q -O- http://localhost:5000/health || exit 1
CMD ["./coinjar"]