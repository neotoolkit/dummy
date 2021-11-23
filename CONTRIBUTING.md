# Contributing to Dummy
[Russian](CONTRIBUTING_RU.md)
# Project structure
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
# Run tests

```
go test ./test/...
```

# Pull Requests
1. Fork the repo and create your branch from `main`.
2. If you've added code that should be tested, add tests.
3. Ensure the test suite passes.
4. Make sure your code lints.