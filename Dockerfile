FROM golang:1.9.4-alpine as builder

WORKDIR /go/src/github.com/rogersole/payments-basic-api
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o payments-basic-api ./cmd/api/main.go

# From a brand new image, copy only the payments-basic-api binary
FROM scratch

COPY --from=builder /go/src/github.com/Typeform/payments-basic-api/payments-basic-api .

CMD ["./payments-basic-api"]