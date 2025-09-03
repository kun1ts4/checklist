FROM golang:1.25-alpine as builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o db ./cmd/main.go

FROM alpine:3.22
WORKDIR /app
RUN apk add --no-cache netcat-openbsd
COPY --from=builder /src/db ./
EXPOSE 50051
CMD ["./db"]



