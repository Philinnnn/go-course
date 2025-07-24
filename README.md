# Go-Course

Сервис для учёта транзакций мерчантов. Реализован на Go, использует PostgreSQL и Gin.

## Быстрый старт

1. **Настройте .env**
   - Пример в файле .env.dist

2. **Соберите и запустите приложение**
   - Локально:
     ```bash
     make build
     make run
     ```
   - В Docker:
     ```bash
     make docker-build
     docker run -p 8080:8080 go-course-app
     ```

3. **Генерация и просмотр документации**
   - Swagger:
     ```bash
     make swagger
     ```
   - Откройте [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

## API

- **POST /api/transactions** — создать транзакцию
- **GET /api/transactions/{id}** — получить транзакцию по ID
- **GET /api/transactions?start=...&end=...** — получить транзакции за период (формат дат RFC3339)
- **PUT /api/transactions/{id}/status** — изменить статус транзакции
- **POST /api/terminals/** — создать терминал
- **GET /api/terminals/** — получить все терминалы
- **GET /api/terminals/{id}** — получить терминал по UUID
- **PUT /api/terminals/{id}** — обновить терминал
- **DELETE /api/terminals/{id}** — удалить терминал

## Импорт команд для Postman

В файле `api_commands.json` содержатся коллекции запросов для Postman. Импортируйте его через интерфейс Postman для быстрого тестирования API.

## Миграции

- Автоматические миграции включаются через переменную `AUTO_MIGRATE=true` в `.env`.

## Зависимости

- Go 1.24+
- PostgreSQL
- swaggo/swag (для генерации Swagger)
- gin-gonic/gin
- gorm.io/gorm

## Сборка и пересборка

- Пересборка Swagger-документации:
  ```bash
  make swagger
  ```
- Пересборка Docker-контейнера:
  ```bash
  make docker-build
  ```

