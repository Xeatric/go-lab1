# Этап сборки
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Копируем модули и загружаем зависимости
COPY go.mod go.sum* ./
RUN go mod download

# Копируем исходный код и собираем
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o main .

# Финальный минимальный образ
FROM alpine:latest

# Устанавливаем CA-сертификаты и создаём пользователя
RUN apk --no-cache add ca-certificates && \
    adduser -D -u 1000 appuser

WORKDIR /app
COPY --from=builder /app/main .
RUN chown -R appuser:appuser /app

USER appuser
EXPOSE 8080
CMD ["./main"]