FROM golang:1.20 as builder

WORKDIR /app
ADD . .
RUN go mod download
RUN go build -o cmd ./main.go

FROM bitnami/kubectl:1.26-debian-11
COPY --from=builder /app/cmd /
CMD ["/cmd"]
