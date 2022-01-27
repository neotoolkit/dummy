Run mock server based off an API contract with one command.

[![Go Report Card](https://goreportcard.com/badge/github.com/go-dummy/dummy)](https://goreportcard.com/report/github.com/go-dummy/dummy)
[![codecov](https://codecov.io/gh/go-dummy/dummy/branch/main/graph/badge.svg?token=2J45SL2XJS)](https://codecov.io/gh/go-dummy/dummy)

### Installation
Dummy requires Go > 1.17
```bash
go install github.com/go-dummy/dummy/cmd/dummy@latest
```
### Usage
Dummy can help you run mock server based off an API contract, which helps people see how your API will work before you even have it built. Run it locally with the `dummy server` command to run your API on a HTTP server you can interact with.
```bash
dummy server ./openapi.yml
```
### Local run
#### Requirements
- go version 1.17 for run linter and tests (make all)
- go version 1.18 for run project
- golangci-lint (`$ go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.43.0`)
```bash
make local
```
More usage [examples](examples)

### :heart:Sponsors
<p>
  <a href="https://evrone.com/?utm_source=github&utm_campaign=dotenv-linter">
    <img src="https://www.mgrachev.com/assets/static/sponsored_by_evrone.svg?sanitize=true"
      alt="Sponsored by Evrone">
  </a>
</p>
