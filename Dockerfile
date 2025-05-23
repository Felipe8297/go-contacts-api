FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o go-contacts-api ./cmd/api/main.go

# Etapa 2: imagem final
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/go-contacts-api .
COPY internal/pkg/migrations ./internal/pkg/migrations

EXPOSE 8080

CMD ["./go-contacts-api"]