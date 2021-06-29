## Log Analyzer

[![Go Reference](https://pkg.go.dev/badge/github.com/itzmanish/go-loganalyzer.svg)](https://pkg.go.dev/github.com/itzmanish/go-loganalyzer) [![Go](https://github.com/itzmanish/go-loganalyzer/actions/workflows/go.yml/badge.svg)](https://github.com/itzmanish/go-loganalyzer/actions/workflows/go.yml) [![codecov](https://codecov.io/gh/itzmanish/go-loganalyzer/branch/master/graph/badge.svg?token=7434KW1MLY)](https://codecov.io/gh/itzmanish/go-loganalyzer) [![Go Report Card](https://goreportcard.com/badge/github.com/itzmanish/go-loganalyzer)](https://goreportcard.com/report/github.com/itzmanish/go-loganalyzer)

## Features

- [x] Logger
- [x] Config
- [x] Client
- [x] Server
- [x] Retry manager
- [x] Store queue

## ToDos

- [ ] Exponential backoff/retry
- [ ] HTTP/GRPC endpoint for getting logs from other service
- [ ] Filter wrapper for filtering logs
- [ ] HTTP API for querying logs from server
- [ ] Filter and Search on server side

## Installation

```
go get -u github.com/itzmanish/go-loganalyzer
```

> Make sure $GOPATH/bin directory is in your path

## Usages

```
$ go-loganalyzer --config .loganalyzer_example.json

Using config file:
loganalyzer is a log analyzer tool which provides a server and agent.
Server command is used to start a server. Whereas agent command is used to start
an agent so that logs can be sent from host machine to server and stored in File system or S3.

Usage:
  loganalyzer [flags]
  loganalyzer [command]

Available Commands:
  agent       Log analyzer agent for collecting logs and sending to server.
  help        Help about any command
  server      Log analyzer server to collect logs from agent and process it.

Flags:
      --config string   config file (default is $HOME/.config/.loganalyzer.yaml)
  -h, --help            help for loganalyzer
  -v, --version         version for loganalyzer

Use "loganalyzer [command] --help" for more information about a command.
```

## Architecture

![LogAnalyzer](https://user-images.githubusercontent.com/12438068/123430818-3e56c380-d5e6-11eb-9020-83b00984deea.png)
![LogAnalyzer-agent-v0](https://user-images.githubusercontent.com/12438068/123430891-4f073980-d5e6-11eb-8b7c-ead15c3adf8f.png)

## Contributions

> Note
> This is just a practice exercise and not production ready.

If anyone wants to make this project real or have some good features/roadmap in mind. Feel free to raise a issue.
Every contributions to this project is ❤️ welcome.

## License

This project is licence under [Apache Licence 2.0](https://github.com/itzmanish/go-loganalyzer/blob/master/LICENSE)
