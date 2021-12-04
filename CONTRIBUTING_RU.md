# Контрибьютинг в Dummy
[English](CONTRIBUTING.md)
# Структура проекта
```
├── cmd
│   └── dummy
│       └── main.go                         # Точка входа CLI
├── internal
│   ├── command                             # Реализация CLI-команд
│   ├── config                              # Конфигурация
│   ├── exitcode                            # Коды ошибок
│   ├── logger                              # Логирование
│   ├── openapi3                            # Пакет для сериализации OpenAPI спецификации
│   └── server                              # Реализация мок-сервера
├── .gitignore
├── .golangci.yml
├── CONTRIBUTING.md
├── CONTRIBUTING_RU.md
├── go.mod
├── go.sum
├── LICENSE
├── README.md
└── Taskfile.yml
```
# Запуск тестов

```
go test ./test/...
```

# Pull Requests
1. Делаем форк репозитория и создаем ветку от `main`.
2. Если вы добавили код, который следует протестировать, добавьте тесты.
3. Убедитесь что все тесты проходят.
4. Убедитесь что ваш код проходит линтеры.
