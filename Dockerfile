FROM golang:1.16-alpine

WORKDIR /go/src/k8s-observer
COPY . .
RUN go build -o k8s-observer

FROM alpine:latest
COPY --from=0 /go/src/k8s-observer .
ENTRYPOINT ./k8s-observer