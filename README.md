# dummy

[![build-img]][build-url]
[![pkg-img]][pkg-url]
[![reportcard-img]][reportcard-url]
[![coverage-img]][coverage-url]
[![version-img]][version-url]

Run mock server based off an API contract with one command

## Features
- support OpenAPI 3.x

## Installation
```shell
go install github.com/neotoolkit/dummy/cmd/dummy@latest
```

## Usage
Dummy can help you run mock server based off an API contract, which helps people see how your API will work before you even have it built. Run it locally with the `dummy server` command to run your API on a HTTP server you can interact with.
```shell
dummy s openapi.yml
```
```shell
dummy s https://raw.githubusercontent.com/neotoolkit/dummy/main/examples/docker/openapi.yml
```
More usage [examples](examples)

## Documentation
See [these docs][pkg-url].

## License
[MIT License](LICENSE).

## Sponsors
<p>
  <a href="https://evrone.com/?utm_source=github&utm_campaign=dotenv-linter">
    <img src="https://raw.githubusercontent.com/neotoolkit/.github/main/assets/sponsored_by_evrone.svg"
      alt="Sponsored by Evrone">
  </a>
</p>

[build-img]: https://github.com/neotoolkit/dummy/workflows/build/badge.svg
[build-url]: https://github.com/neotoolkit/dummy/actions
[pkg-img]: https://pkg.go.dev/badge/neotoolkit/dummy
[pkg-url]: https://pkg.go.dev/github.com/neotoolkit/dummy
[reportcard-img]: https://goreportcard.com/badge/neotoolkit/dummy
[reportcard-url]: https://goreportcard.com/report/neotoolkit/dummy
[coverage-img]: https://codecov.io/gh/neotoolkit/dummy/branch/main/graph/badge.svg
[coverage-url]: https://codecov.io/gh/neotoolkit/dummy
[version-img]: https://img.shields.io/github/v/release/neotoolkit/dummy
[version-url]: https://github.com/neotoolkit/dummy/releases
