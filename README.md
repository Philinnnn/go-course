# Go-Course

Go-Course — сервис для учёта транзакций мерчантов на Go с использованием PostgreSQL и Gin.

## Быстрый старт

1. **Настройте переменные окружения**
   - Скопируйте `.env.dist` в `.env` и заполните параметры подключения к базе данных.

2. **Сборка и запуск**
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

3. **Документация**
   - Сгенерируйте Swagger:
     ```bash
     make swagger
     ```
   - Откройте [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

## API

- **POST /api/transactions** — создать транзакцию
- **GET /api/transactions/{id}** — получить транзакцию по ID
- **GET /api/transactions?start=...&end=...** — получить транзакции за период
- **PUT /api/transactions/{id}/status** — изменить статус транзакции
- **POST /api/terminals/** — создать терминал
- **GET /api/terminals/** — получить все терминалы
- **GET /api/terminals/{id}** — получить терминал по UUID
- **PUT /api/terminals/{id}** — обновить терминал
- **DELETE /api/terminals/{id}** — удалить терминал
- **GET /api/currency/convert** — конвертация валют (демо)
- **GET /api/currency/rates** — получить текущие курсы валют

## Пример запроса курсов валют

```bash
curl http://localhost:8080/api/currency/rates
```

## Postman

Импортируйте коллекцию запросов из файла `api_commands.json` для быстрого тестирования.

## Миграции

Автоматические миграции включаются через `AUTO_MIGRATE=true` в `.env`.

## Зависимости

- Go 1.24+
- PostgreSQL
- swaggo/swag
- gin-gonic/gin
- gorm.io/gorm

## Сборка и пересборка

- Swagger:
  ```bash
  make swagger
  ```
- Docker:
  ```bash
  make docker-build
  ```
