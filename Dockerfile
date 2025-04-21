FROM golang:1.24@sha256:d9db32125db0c3a680cfb7a1afcaefb89c898a075ec148fdc2f0f646cc2ed509

WORKDIR /go/src/k8s-observer
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o k8s-observer

FROM scratch
COPY --from=0 /go/src/k8s-observer .
ENTRYPOINT ["./k8s-observer"]