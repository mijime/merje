# merje

Supports the merge hash from any config files.

[![CircleCI](https://circleci.com/gh/mijime/merje.svg?style=svg)](https://circleci.com/gh/mijime/merje)

## Description

## Usage

```
merje [OPTIONS] [SOURCE FILES ...]

Application Options:
  -i, --input-format:  input format
  -f, --format:        output format
  -o, --out:           output path
  -t, --type:          merge type (default: or)
  -v, --version        print a version

Help Options:
  -?                  Show this help message
  -h, --help          Show this help message
```

## Example

Merge from [examples/input-1.json](examples/input-1.json) [examples/input-2.yml](examples/input-2.yml) to [examples/input-3.toml](examples/input-3.toml)

```bash
$ merje examples/input-1.json examples/input-2.yml --out examples/input-3.toml
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
