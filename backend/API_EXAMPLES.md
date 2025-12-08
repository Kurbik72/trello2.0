# Примеры использования API

## Получение фоновых изображений

```bash
curl -X GET http://localhost:8080/api/default-backgrounds
```

**Ответ:**

```json
[
  {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "src": "https://picsum.photos/id/100/800/400"
  },
  {
    "id": "550e8400-e29b-41d4-a716-446655440001",
    "src": "https://picsum.photos/id/200/800/400"
  }
]
```

## Получение досок пользователя (с JWT токеном)

```bash
curl -X GET \
  "http://localhost:8080/api/get-boards" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

**Ответ:**

```json
[
  {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "title": "My Board",
    "isFavorite": false,
    "linkToBoard": "/board/550e8400-e29b-41d4-a716-446655440000",
    "backgroundSrc": "https://picsum.photos/id/100/800/400"
  }
]
```

## Получение досок пользователя (без JWT, через query параметр)

```bash
curl -X GET \
  "http://localhost:8080/api/get-boards?user_id=550e8400-e29b-41d4-a716-446655440000"
```

## Создание новой доски

```bash
curl -X POST \
  http://localhost:8080/api/save-board \
  -H "Content-Type: application/json" \
  -d '{
    "title": "New Board",
    "backgroundId": "550e8400-e29b-41d4-a716-446655440000",
    "user_id": "6949e844-9b4b-4dec-9d42-2d4216415f53"
  }'
```

**Ответ:**

```json
{
  "id": "550e8400-e29b-41d4-a716-446655440001",
  "title": "New Board",
  "isFavorite": false,
  "linkToBoard": "/board/550e8400-e29b-41d4-a716-446655440001",
  "backgroundSrc": "https://picsum.photos/id/100/800/400"
}
```

## Переключение статуса избранного

```bash
curl -X PATCH \
  http://localhost:8080/api/save-favorite-board \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "boardId": "550e8400-e29b-41d4-a716-446655440000"
  }'
```

**Ответ:**

```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "title": "My Board",
  "isFavorite": true,
  "linkToBoard": "/board/550e8400-e29b-41d4-a716-446655440000",
  "backgroundSrc": "https://picsum.photos/id/100/800/400"
}
```

## Health Check

```bash
curl -X GET http://localhost:8080/health
```

**Ответ:**

```json
{
  "status": "ok"
}
```
