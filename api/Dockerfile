FROM golang:1.25-alpine as builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api ./main.go

FROM alpine:3.22
WORKDIR /app
COPY --from=builder /src/api ./
EXPOSE 8080
CMD ["./api"]
