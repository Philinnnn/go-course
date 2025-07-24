# Базовый образ
FROM golang:1.24.5

# Рабочая директория внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем весь проект
COPY . .

# Собираем приложение
RUN go build -o app ./cmd/go-course

# Открываем порт
EXPOSE 8080

# Команда запуска
CMD ["./app"]