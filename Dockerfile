FROM golang:1.19 as build

WORKDIR /k8s-observer
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o k8s-observer

FROM scratch
COPY --from=build /k8s-observer/k8s-observer .
ENTRYPOINT ["./k8s-observer"]