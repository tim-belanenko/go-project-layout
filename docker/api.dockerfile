FROM golang:1.21.5 AS builder
WORKDIR /build
COPY go.* .
RUN go mod download
COPY . . 
RUN CGO_ENABLED=0 GOOS=linux go build -o api ./cmd/api/main.go

FROM alpine:3.19.0
COPY --from=builder /build/api /app/api
CMD [ "/app/api" ]