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
   gjson [global options] command [command options]

VERSION:
   v0.0.3 (2e32198281acb1308e1b50abd7675ce9e116329c)

COMMANDS:
   completion                      command completion
   validation, validate, valid, v  json file validation
   help, h                         Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --log-level value, -l value  log level [$LOG_LEVEL]
   --help, -h                   show help
   --version, -v                print the version
```

## validate

```
NAME:
   gjson validation - json file validation

USAGE:
   gjson validation command [command options] [file or directory...]

OPTIONS:
   --log-level value, -l value                              log level [$LOG_LEVEL]
   --include value, -i value [ --include value, -i value ]  include file extension (default: "**/*.json", "**/*.json.golden") [$INCLUDE]
   --exclude value, -e value [ --exclude value, -e value ]  exclude file extension (default: "**/*.invalid.json", "**/*.invalid.json.golden", "**/node_modules/**") [$EXCLUDE]
   --help, -h                                               show help
```
