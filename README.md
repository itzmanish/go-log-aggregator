## Log aggregator

[![Go Reference](https://pkg.go.dev/badge/github.com/itzmanish/go-log-aggregator.svg)](https://pkg.go.dev/github.com/itzmanish/go-log-aggregator) [![Go](https://github.com/itzmanish/go-log-aggregator/actions/workflows/go.yml/badge.svg)](https://github.com/itzmanish/go-log-aggregator/actions/workflows/go.yml) [![codecov](https://codecov.io/gh/itzmanish/go-log-aggregator/branch/master/graph/badge.svg?token=7434KW1MLY)](https://codecov.io/gh/itzmanish/go-log-aggregator) [![Go Report Card](https://goreportcard.com/badge/github.com/itzmanish/go-log-aggregator)](https://goreportcard.com/report/github.com/itzmanish/go-log-aggregator)

## Features

---

- [x] Logger
- [x] Config
- [x] Client
- [x] Server
- [x] Retry manager
- [x] Store queue
- [x] Buffering before storing

## ToDos

---

- [ ] Exponential backoff (I am not sure if we need it here.)
- [ ] HTTP API for querying logs from server
- [ ] Filter and Search on server side
- [ ] Filter wrapper for filtering logs
- [ ] HTTP/GRPC endpoint for getting logs from other service

> I am not going to create a query algorithm or query api because this is unneccessary, complex and time consuming.
> So my focus is to make it more like Fluentd. There will be plugin in future for data visualizing tools/platforms. eg: Elasticsearch,datadog,logstash etc. (Again If you got a solution of valid argument point, create a issue and I am open to discussion.)

## Roadmap

---

Make its architecture like flutend.

## Installation

---

```
go get -u github.com/itzmanish/go-log-aggregator
```

> Make sure $GOPATH/bin directory is in your path

## Usages

---

```bash
$ go-log-aggregator -h

log-aggregator is a log aggregating tool which provides a server and agent.
Server command is used to start a server. Whereas agent command is used to start
an agent so that logs can be sent from host machine to server and stored in File system or S3.

Usage:
  log-aggregator [flags]
  log-aggregator [command]

Available Commands:
  agent       Log aggregator agent for collecting logs and sending to server.
  help        Help about any command
  server      Log aggregator server to collect logs from agent and process it.

Flags:
      --config string   config file (default is $HOME/.config/.log-aggregator.yaml)
  -h, --help            help for log-aggregator
  -v, --version         version for log-aggregator

Use "log-aggregator [command] --help" for more information about a command.
```

## Architecture

---

![LogAnalyzer](https://user-images.githubusercontent.com/12438068/123430818-3e56c380-d5e6-11eb-9020-83b00984deea.png)
![LogAnalyzer-agent-v0](https://user-images.githubusercontent.com/12438068/123430891-4f073980-d5e6-11eb-8b7c-ead15c3adf8f.png)

## Contributions

---

> Note
> This is just a practice exercise and not production ready.

If anyone wants to make this project real or have some good features/roadmap in mind. Feel free to raise a issue.
Every contributions to this project is ❤️ welcome.

## License

---

This project is licence under [Apache Licence 2.0](https://github.com/itzmanish/go-log-aggregator/blob/master/LICENSE)
