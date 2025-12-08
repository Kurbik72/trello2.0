# Trello Backend API

RESTful API для управления досками Trello, написанный на Go с использованием Gin framework.

## Архитектура

Проект следует принципам Clean Architecture с разделением на слои:

```
backend/
├── cmd/
│   └── server/          # Точка входа приложения
├── internal/
│   ├── config/          # Конфигурация приложения
│   ├── database/        # Подключение к БД
│   ├── models/          # Модели данных и DTO
│   ├── repository/      # Слой доступа к данным
│   ├── service/         # Бизнес-логика
│   ├── handler/         # HTTP handlers
│   ├── middleware/      # Middleware (auth, CORS)
│   └── router/          # Настройка роутинга
└── pkg/
    ├── jwt/             # JWT утилиты
    └── password/        # Хеширование паролей
```

## Требования

- Go 1.21+
- PostgreSQL 16+
- Docker и Docker Compose (опционально)

## Установка

1. Клонируйте репозиторий и перейдите в директорию backend:
```bash
cd backend
```

2. Установите зависимости:
```bash
go mod download
```

3. Настройте переменные окружения:
```bash
cp .env.example .env
# Отредактируйте .env файл при необходимости
```

4. Запустите PostgreSQL через Docker Compose:
```bash
docker-compose up -d
```

5. Запустите миграции Prisma (если используете):
```bash
npx prisma migrate dev
```

6. Запустите сервер:
```bash
go run cmd/server/main.go
```

Сервер будет доступен по адресу `http://localhost:8080`

## API Endpoints

### Публичные endpoints

#### GET /api/default-backgrounds
Получить список доступных фоновых изображений.

**Response:**
```json
[
  {
    "id": "uuid",
    "src": "https://example.com/image.jpg"
  }
]
```

### Защищенные endpoints (требуют JWT токен)

#### GET /api/get-boards?user_id={user_id}
Получить все доски пользователя.

**Headers:**
```
Authorization: Bearer <token>
```

**Response:**
```json
[
  {
    "id": "uuid",
    "title": "Board Title",
    "isFavorite": false,
    "linkToBoard": "/board/uuid",
    "backgroundSrc": "https://example.com/image.jpg"
  }
]
```

#### POST /api/save-board
Создать новую доску.

**Headers:**
```
Authorization: Bearer <token>
Content-Type: application/json
```

**Request Body:**
```json
{
  "title": "Board Title",
  "backgroundId": "uuid"
}
```

**Response:**
```json
{
  "id": "uuid",
  "title": "Board Title",
  "isFavorite": false,
  "linkToBoard": "/board/uuid",
  "backgroundSrc": "https://example.com/image.jpg"
}
```

#### PATCH /api/save-favorite-board
Переключить статус избранного для доски.

**Headers:**
```
Authorization: Bearer <token>
Content-Type: application/json
```

**Request Body:**
```json
{
  "boardId": "uuid"
}
```

**Response:**
```json
{
  "id": "uuid",
  "title": "Board Title",
  "isFavorite": true,
  "linkToBoard": "/board/uuid",
  "backgroundSrc": "https://example.com/image.jpg"
}
```

## Безопасность

- JWT токены для аутентификации
- Хеширование паролей с использованием bcrypt
- Валидация входных данных
- Проверка прав доступа к ресурсам
- CORS middleware для защиты от CSRF

## Переменные окружения

```env
# Database
DATABASE_URL=postgres://user:password@localhost:5432/trello_db?sslmode=disable

# Server
PORT=8080
GIN_MODE=debug

# JWT
JWT_SECRET=your-secret-key-change-in-production
JWT_EXPIRATION_HOURS=24
REFRESH_TOKEN_EXPIRATION_DAYS=7
```

## Разработка

### Запуск в режиме разработки

```bash
GIN_MODE=debug go run cmd/server/main.go
```

### Сборка

```bash
go build -o bin/server cmd/server/main.go
```

### Тестирование

```bash
go test ./...
```

## Структура базы данных

- **users** - Пользователи системы
- **boards** - Доски пользователей
- **default_backgrounds** - Фоновые изображения

Подробная схема в `prisma/schema.prisma`

