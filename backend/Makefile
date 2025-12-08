.PHONY: run build test clean migrate-up migrate-down docker-up docker-down

# Запуск приложения
run:
	go run cmd/server/main.go

# Сборка приложения
build:
	go build -o bin/server cmd/server/main.go

# Запуск тестов
test:
	go test ./...

# Очистка бинарных файлов
clean:
	rm -rf bin/

# Запуск Docker Compose
docker-up:
	docker-compose up -d

# Остановка Docker Compose
docker-down:
	docker-compose down

# Миграции Prisma (если используете)
migrate-up:
	npx prisma migrate dev

migrate-down:
	npx prisma migrate reset

# Установка зависимостей
deps:
	go mod download
	go mod tidy

# Форматирование кода
fmt:
	go fmt ./...

# Линтинг
lint:
	golangci-lint run

