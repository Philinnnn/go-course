APP_NAME = go-course-app
BIN_DIR = ../bin

.PHONY: run build link

## Запуск приложения
run:
	go run ./cmd/go-course/main.go

## Сборка бинарника в ../bin
build:
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(APP_NAME) ./cmd/go-course/main.go

## Проверка форматирования и статики
link:
	go fmt ./...
	go vet ./...