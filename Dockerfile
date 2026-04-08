# Шаг 1: Сборка приложения (используем легковесный образ golang)
FROM golang:alpine AS builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем файлы зависимостей и скачиваем их
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь остальной исходный код
COPY . .

# Собираем бинарный файл (назовем его todo-app)
RUN go build -o todo-app cmd/main.go

# Шаг 2: Финальный образ (минимального размера)
FROM alpine:latest

WORKDIR /root/

# Копируем собранный бинарник из первого шага
COPY --from=builder /app/todo-app .
# Копируем папку с конфигами, так как viper будет их искать
COPY --from=builder /app/configs ./configs
# Копируем файл .env
COPY --from=builder /app/.env .

# Указываем команду для запуска
CMD ["./todo-app"]