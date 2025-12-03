# FrogShort 🐸

Минималистичный URL-shortener на Go, упакованный в Docker и docker-compose.

## Функциональность

- Создание коротких ссылок:
  - `POST /shorten` с JSON `{"url": "https://example.com"}`
  - или `GET /shorten?url=https://example.com`
- Редирект по короткому коду: `GET /{code}` → `302` на оригинальный URL
- Хранение в памяти (in-memory) — идеально для демо и локальной разработки.

---

## Запуск локально (без Docker)

```bash
go mod tidy
go run main.go
```

По умолчанию сервис слушает порт `8080`  
(можно переопределить переменной окружения `PORT`).

---

## Примеры

Создание короткой ссылки (POST):

```bash
curl -X POST http://localhost:8080/shorten \
  -H "Content-Type: application/json" \
  -d '{"url":"https://www.amazon.fr/..."}'
```

Создание короткой ссылки (GET):

```bash
curl "http://localhost:8080/shorten?url=https://www.amazon.fr/..."
```

Ответ:

```json
{"short":"http://localhost:8080/aB3kL9Q"}
```

Переход по короткой ссылке:

```text
http://localhost:8080/aB3kL9Q
```

---

## Docker

Собрать образ:

```bash
make docker-build
```

Запустить:

```bash
make docker-run
```

Остановить контейнер:

```bash
make docker-stop
```

---

## Docker Compose

Запуск через docker-compose:

```bash
make compose-up
```

Остановка:

```bash
make compose-down
```

Сервис будет доступен на:  
`http://localhost:8080`
# FrogShort
