# Сервис домов на Go + PostgreSQL с использованием Docker

## Структура проекта
```
├── Dockerfile # Dockerfile для сборки и запуска Go сервиса
├── docker-compose.yml # Файл для настройки и запуска сервисов с использованием Docker Compose
├── go.mod # Модульные зависимости Go
├── go.sum # Контрольные суммы для зависимостей Go
├── README.md # Описание проекта
├── .gitignore # Файл для игнорирования ненужных файлов в Git
├── cmd/
│ └── app/
│ └── main.go # entrypoint сервиса
├── internal/
│ ├── api/
│ │ ├── server.go # Конфигурация API сервера
│ │ └── swagger.yaml # Файл описания API (Swagger)
│ ├── auth/
│ │ ├── auth_middleware.go # Middleware для аутентификации
│ │ ├── login_handler.go # Обработчик входа (DummyLogin)
│ │ └── service.go 
│ ├── flat/
│ │ ├── flat_handler.go # Обработчик запросов для сущности Flat
│ │ └── service.go 
│ ├── house/
│ │ ├── house_handler.go # Обработчик запросов для сущности House
│ │ └── service.go 
│ ├── models/
│ │ └── models.go # Определение моделей данных
│ ├── repository/
│ │ ├── flat_repository.go # Репозиторий для сущности Flat
│ │ └── house_repository.go # Репозиторий для сущности House
│ └── tools/
│ └── tokenutils.go # Утилиты для работы с токенами
├── db/
│ ├── populate.sql # Скрипт для заполнения БД тестовыми данными
│ └── migrations/
│ ├── migrate_down.sql # Скрипт для отката миграций
│ └── migrate_up.sql # Скрипт для применения миграций
```

## Требования для запуска проекта в контейнере

- Docker
- Docker Compose

## Как запустить проект

1. Клонируйте репозиторий на ваш локальный компьютер:

    ```bash
    git clone https://github.com/dmitrykhayorov/avito_test.git
    cd avito_test
    ```

2. Убедитесь, что Docker и Docker Compose установлены и работают на вашем компьютере.

3. Запустите сервисы с использованием Docker Compose:

    ```bash
    docker-compose up --build
    ```

   Эта команда выполнит следующие шаги:
    - Сборка Docker образа для сервиса домов.
    - Запуск контейнера с базой данных PostgreSQL.
    - Применение миграций для создания таблиц в базе данных.
    - Запуск Go сервиса, который будет подключаться к базе данных PostgreSQL.

4. Сервис будет доступен по адресу `http://localhost:8080`. Порт можно изменить в файле docker-compose.yml

## Примечания

- По умолчанию, база данных PostgreSQL будет доступна на порту `5432`.
- Файл миграции `migrate_up.sql` создаёт таблицы, если они ещё не существуют.
- Переменные окружения, такие как `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASSWORD`, и `DB_NAME`, настроены в файле `docker-compose.yml` и используются Go сервисом для подключения к базе данных.

