# go-chat

Real-time чат на Go с WebSocket, JWT авторизацией и PostgreSQL.

## Стек

| Что | Чем |
|---|---|
| Язык | Go |
| WebSocket | gorilla/websocket |
| Роутер | go-chi/chi |
| БД | PostgreSQL + pgx |
| Кеш | Redis |
| Миграции | golang-migrate |
| JWT | golang-jwt/jwt |
| DI | uber/fx |
| Контейнеры | Docker + Docker Compose |

## Возможности

- Регистрация и авторизация через JWT
- Создание комнат и общение в реальном времени
- История сообщений с пагинацией
- Множество одновременных подключений через горутины и каналы

## Запуск

### Требования
- Docker
- Docker Compose

### 1. Клонировать репозиторий
```bash
git clone https://github.com/SilverName608/go-chat.git
cd go-chat
```

### 2. Создать .env файл
```bash
cp .env.example .env
```

### 3. Запустить
```bash
docker-compose up --build
```

Приложение будет доступно на `http://localhost:8040`

## Локальный запуск без Docker

### Требования
- Go 1.25.6+
- PostgreSQL
- Redis

### 1. Заполнить .env
```env
DB_DSN=postgres://postgres:password@localhost:5432/go_chat?sslmode=disable
REDIS_DSN=redis://localhost:6379
HTTP_PORT=8040
JWT_SECRET=supersecretkey
```

### 2. Применить миграции
```bash
migrate -path migrations -database "$DB_DSN" up
```

### 3. Запустить
```bash
go run cmd/server/main.go
```

## Структура проекта

```
go-chat/
├── cmd/server/         # точка входа
├── internal/
│   ├── api/            # HTTP и WebSocket хендлеры
│   ├── application/    # реализация бизнес логики
│   ├── config/         # конфигурация
│   ├── di/             # dependency injection
│   ├── domain/         # модели и интерфейсы сервисов
│   ├── hub/            # WebSocket ядро
│   └── infrastructure/ # репозитории, БД, Redis
├── migrations/         # SQL миграции
└── web/                # фронтенд
```

## API

### Авторизация
| Метод | Путь | Описание |
|---|---|---|
| POST | /api/v1/auth/register | Регистрация |
| POST | /api/v1/auth/login | Вход |

### Комнаты
| Метод | Путь | Описание |
|---|---|---|
| GET | /api/v1/rooms | Список комнат |
| POST | /api/v1/rooms | Создать комнату |
| GET | /api/v1/rooms/{id} | Комната по ID |
| DELETE | /api/v1/rooms/{id} | Удалить комнату |

### WebSocket
| Метод | Путь | Описание |
|---|---|---|
| GET | /api/v1/ws/{roomID} | Подключиться к комнате |