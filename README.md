# merje

Supports the merge hash from any config files.

## Description

## Usage

```
Usage of merje:
  -decode string
        json/yaml/toml
  -encode string
        json/yaml/toml/template
  -merge string
        or/and/xor (default "or")
  -out string

  -version
        dev
```

## Example

Merge from [examples/input-1.json](examples/input-1.json) [examples/input-2.yml](examples/input-2.yml) to [examples/input-3.toml](examples/input-3.toml)

```bash
$ merje --out examples/input-3.toml examples/input-1.json examples/input-2.yml
```

## Install

To install, use `go get`:

```bash
$ go get -d github.com/mijime/merje
```

or download: [https://github.com/mijime/merje/releases/latest](https://github.com/mijime/merje/releases/latest)

## Contribution

1. Fork ([https://github.com/mijime/merje/fork](https://github.com/mijime/merje/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[mijime](https://github.com/mijime)
