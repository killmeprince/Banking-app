FROM golang:1.23 as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o banking-app ./cmd/server/main.go

FROM golang:1.23
WORKDIR /app
COPY --from=builder /app/banking-app .
EXPOSE 8080
CMD ["./banking-app"]
