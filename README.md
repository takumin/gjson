# gjson
Golang JSON Tool

[![CI](https://github.com/takumin/gjson/actions/workflows/integration.yml/badge.svg)](https://github.com/takumin/gjson/actions/workflows/integration.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/takumin/gjson)](https://goreportcard.com/report/github.com/takumin/gjson)

## Install

### GitHub Releases

Download a prebuilt binary from [GitHub Releases](https://github.com/takumin/gjson/releases) and install it in $PATH.

### aqua

[aqua](https://aquaproj.github.io/) is a CLI Version Manager.

```bash
aqua g -i takumin/gjson
```

## usage

```
NAME:
   gjson - Golang JSON Tool

USAGE:
   gjson [global options] command [command options] [arguments...]

VERSION:
   dev (a39580b79e5739d9872728be106d5c051f8ed51f)

COMMANDS:
   completion                      command completion
   validation, validate, valid, v  json file validation
   help, h                         Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h                   show help (default: false)
   --log-level value, -l value  log level [$LOG_LEVEL]
   --version, -v                print the version (default: false)
```

## validate

```
NAME:
   gjson validation - json file validation

USAGE:
   gjson validation [command options] [arguments...]

OPTIONS:
   --directory value, -d value  search base directory (default: ".") [$DIRECTORY]
   --excludes value, -e value  exclude files extensions (default: "invalid.json.golden") [$EXCLUDES]
   --includes value, -i value  include files extensions (default: "json,json.golden") [$INCLUDES]
   --log-level value, -l value  log level [$LOG_LEVEL]
```
