FROM golang:1.21.5 AS builder
WORKDIR /build
COPY go.* .
RUN go mod download
COPY . . 
RUN CGO_ENABLED=0 GOOS=linux go build -o ./migration ./cmd/migration/main.go

FROM alpine:3.19.0
WORKDIR /app
COPY --from=builder /build/migration /app/migration
CMD [ "/app/migration" ]