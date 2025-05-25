FROM golang:1.24 as builder

WORKDIR /app
COPY . .

# Сначала копируем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем конфиги в правильное место
COPY config/.env /app/config/.env

# Копируем остальной код и собираем
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /bookit ./cmd/server/

FROM alpine:latest
WORKDIR /app

# Копируем необходимые файлы
COPY --from=builder /bookit /app/bookit
COPY --from=builder /app/migrations /app/migrations
COPY --from=builder /app/config/.env /app/config/.env

CMD ["sh", "-c", "/app/bookit migrate && /app/bookit"]