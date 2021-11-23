# Контрибьютинг в Dummy
[English](CONTRIBUTING.md)
# Структура проекта
```
├── cmd
│   └── dummy
│       └── main.go
├── internal
│   ├── command
│   ├── config
│   ├── exitcode
│   ├── logger
│   ├── openapi3
│   └── server
├── test
│   └── cases
│       └── '{case name}'
│           ├── '{openapi path}'
│           │   └── '{method}'
│           │       └── '{response}'.json
│           ├── header.txt
│           └── openapi3.yml
├── .gitignore
├── .golangci.yaml
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
