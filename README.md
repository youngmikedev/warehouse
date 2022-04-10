# Веб сервер для ведения складского учета

## Запуск

- Установить golang.
- Установить docker и docker compose.
- Запустить базу данных в docker: 
```
docker-compose up -d
```
- Запустить приложение:
```
go run cmd/main.go
```

## Список задач
- [X] Пользлователь
    - [X] Аутентификация
    - [X] Хэндлеры
    - [X] Сервис
    - [X] Репозиторий
- [X] Товары
    - [X] Хэндлеры
    - [X] Сервис
    - [X] Репозиторий

## Описание
Приложение использует чистую архитектуру со слоями домен (`domain`), сервис (`service`), доставка (`delivery`) и репозиторий (`repository`).

Конфигурируется приложение с помощью [переменных окружения (env)](#Переменные-окружения-env).

В качестве транспортного слоя используется swagger.

Для хранения данный используется Postgres.
Тестировалось и разрабатывалось на docker образе `postgres:11-alpine`.

Основная логика в сервисах.

Логирование осуществляется с помощью `zerolog`.

## Данное приложение использует два модуля с кодогенерацией [go-swagger](https://github.com/go-swagger/go-swagger) и [ent](https://github.com/ent/ent)

### [Go-swagger](https://github.com/go-swagger/go-swagger)

> This package contains a golang implementation of Swagger 2.0 (aka OpenAPI 2.0): it knows how to serialize and deserialize swagger specifications.

Спецификация: `./internal/delivery/swagger/swagger.yaml`.

Сгенерированные файлы: `./internal/delivery/swagger/gen`.

Хэндлеры: `./internal/delivery/swagger/swagger.go`.

URI документации: `api/v1/docs`

При изменении спецификации нужно заново сгенерировать go файлы, для этого используется команда:
```
swagger generate server -A delivery -t ./internal/delivery/swagger/gen -f ./internal/delivery/swagger/swagger.yaml --exclude-main
```


### [Ent](https://github.com/ent/ent) (ORM)

> Simple, yet powerful entity framework for Go, that makes it easy to build and maintain applications with large data-models.

Схема для кодогенерации находится по пути: `./internal/repository/postgres/ent/schema`.

Сгенерированные файлы находятся по пути: `./internal/repository/postgres/ent`.

При изменении спецификации нужно заново сгенерировать go файлы, для этого используется команда:
```
go generate ./internal/repository/postgres/ent
```

## Переменные окружения (env)

| Ключ  | Описание |
|-|-|
| LOG_LEVEL | уровень логирования |
| HTTP_PORT | на каком порту принимать соединения |
| HTTP_HOST | на каком хосте принимать соединения |
| DB_PORT | на каком порту находится база данных |
| DB_HOST | на каком хосту находится база данных |
| DB_USER | имя пользователя базы данных |
| DB_PAS | пароль базы данных |
| DB_NAME | название базы данных |
| TOK | сигнатура для создания токенов |
| TOK_EXPIRES_MIN | время жизни access токена |
| REF_TOK_EXPIRES_MIN | время жизни refresh токена |
### Значения по умолчанию
```
LOG_LEVEL=0
HTTP_PORT=8000
HTTP_HOST=
DB_PORT=5432
DB_HOST=localhost
DB_USER=root
DB_PAS=example
DB_NAME=warehouse
TOK=60e09d0d8fa190a9c6edb7bd
TOK_EXPIRES_MIN=30
REF_TOK_EXPIRES_MIN=360
```